// +build integration

package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/codenotary/immudb/pkg/client"
	"github.com/codenotary/immudb/pkg/server"
	"github.com/stretchr/testify/assert"

	"github.com/codenotary/ctrlt/pkg/container"
	"github.com/codenotary/ctrlt/pkg/notary"

	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/docker"
	"github.com/codenotary/ctrlt/pkg/logger"
	"github.com/codenotary/ctrlt/pkg/util"
)

var _ = (func() interface{} {
	_ = di.Register(
		di.Entry{
			Name: Logger,
			Maker: func() (interface{}, error) {
				return logger.NewSimpleLogger("ctrl-t/test", os.Stderr), nil
			},
		},
		di.Entry{
			Name: ImmuClient,
			Maker: func() (interface{}, error) {
				return client.DefaultClient(), nil
			},
		},
		di.Entry{
			Name: Notary,
			Maker: func() (interface{}, error) {
				return notary.NewImmuNotary()
			}},
		di.Entry{
			Name: DockerClient,
			Maker: func() (interface{}, error) {
				return docker.NewMockClient()
			},
		},
		di.Entry{
			Name: ContainerNotary,
			Maker: func() (interface{}, error) {
				return container.NewDockerNotary()
			},
		},
		di.Entry{
			Name: ApiServer,
			Maker: func() (interface{}, error) {
				return NewApiServer()
			},
		})
	return nil
})()

func TestMain(m *testing.M) {
	var code int
	if err := util.WithImmuServer(func(immuServer *server.ImmuServer) error {
		if err := di.Initialize(); err != nil {
			return err
		} else {
			code = m.Run()
			return di.Terminate()
		}
	}); err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestListContainersIntegration(t *testing.T) {
	server := di.LookupOrPanic(ApiServer).(*apiServer)
	req, err := http.NewRequest("GET", "/containers", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.listContainers)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, "application/json", rr.Header().Get("content-type"))
	assert.JSONEq(t, `[
			{"Image":{"Name":"name","Hash":"hash"},
			"Notarization":{"Hash":"","Status":"Unknown","Index":0}}
		]`, rr.Body.String())
}

func TestNotarizeIntegration(t *testing.T) {
	server := di.LookupOrPanic(ApiServer).(*apiServer)
	req, err := http.NewRequest("POST", "/notarize?hash=hash&status=Notarized", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.notarize)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, "application/json", rr.Header().Get("content-type"))
	assert.JSONEq(t, `"ok"`, rr.Body.String())
}

func TestBulkNotarizeIntegration(t *testing.T) {
	server := di.LookupOrPanic(ApiServer).(*apiServer)
	req, err := http.NewRequest("POST", "/bulk-notarize?query=test&status=Notarized", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.bulkNotarize)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, "application/json", rr.Header().Get("content-type"))
	assert.JSONEq(t, `"ok"`, rr.Body.String())
}

func TestHealthCheckIntegration(t *testing.T) {
	server := di.LookupOrPanic(ApiServer).(*apiServer)
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.healthCheck)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, "application/json", rr.Header().Get("content-type"))
	assert.Equal(t, `"ok"`, strings.TrimSpace(rr.Body.String()))
}

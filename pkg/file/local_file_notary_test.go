package file

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/logger"
	"github.com/codenotary/ctrlt/pkg/mocks"
	"github.com/codenotary/ctrlt/pkg/notary"
)

const (
	expectedHash  = "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3"
	expectedIndex = uint64(0)
)

var file *os.File

func TestMain(m *testing.M) {
	var err error
	file, err = ioutil.TempFile("", "notarization")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(file.Name())
	err = ioutil.WriteFile(file.Name(), []byte("123"), os.ModePerm)
	if err != nil {
		panic(err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestNotarize(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockNotary := mocks.NewMockNotary(ctrl)
	defer ctrl.Finish()
	fileNotary := LocalFileNotary{
		logger: logger.NewSimpleLogger("ctrlt", os.Stdout),
		notary: mockNotary,
	}
	mockNotary.EXPECT().
		Notarize(expectedHash, constants.Notarized).
		Return(&notary.Notarization{
			Hash:      expectedHash,
			Status:    constants.Notarized,
			StoreMeta: notary.NewStoreMeta(expectedIndex),
		}, nil).
		Times(1)
	notarization, err := fileNotary.Notarize(file.Name(), constants.Notarized)
	assert.NoError(t, err)
	assert.Equal(t, notarization.Status, constants.Notarized)
	assert.Equal(t, notarization.Hash, expectedHash)
	assert.Equal(t, notarization.StoreMeta.(map[string]interface{})["index"], expectedIndex)
}

func TestAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockNotary := mocks.NewMockNotary(ctrl)
	defer ctrl.Finish()
	fileNotary := LocalFileNotary{
		logger: logger.NewSimpleLogger("ctrlt", os.Stdout),
		notary: mockNotary,
	}
	mockNotary.EXPECT().
		Authenticate(expectedHash).
		Return(&notary.Notarization{
			Hash:      expectedHash,
			Status:    constants.Notarized,
			StoreMeta: notary.NewStoreMeta(expectedIndex),
		}, nil).
		Times(1)
	notarization, err := fileNotary.Authenticate(file.Name())
	assert.NoError(t, err)
	assert.Equal(t, notarization.Status, constants.Notarized)
	assert.Equal(t, notarization.Hash, expectedHash)
	assert.Equal(t, notarization.StoreMeta.(map[string]interface{})["index"], expectedIndex)
}

func TestHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockNotary := mocks.NewMockNotary(ctrl)
	defer ctrl.Finish()
	fileNotary := LocalFileNotary{
		logger: logger.NewSimpleLogger("ctrlt", os.Stdout),
		notary: mockNotary,
	}
	mockNotary.EXPECT().
		History(expectedHash).
		Return([]*notary.Notarization{{
			Hash:      expectedHash,
			Status:    constants.Notarized,
			StoreMeta: notary.NewStoreMeta(expectedIndex),
		}}, nil).
		Times(1)
	history, err := fileNotary.History(file.Name())
	assert.NoError(t, err)
	assert.Len(t, history, 1)
	assert.Equal(t, history[0].Status, constants.Notarized)
	assert.Equal(t, history[0].Hash, expectedHash)
	assert.Equal(t, history[0].StoreMeta.(map[string]interface{})["index"], expectedIndex)
}

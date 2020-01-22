package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/codenotary/di/pkg/di"

	"github.com/codenotary/ctrlt/pkg/util"
)

const defaultAddr = "127.0.0.1:4040"

func NewCtrlTCmd() *cobra.Command {
	return &cobra.Command{
		Use: "ctrlt",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := di.Initialize(); err != nil {
				util.Die("initializing failed:", err)
			}
			signalChan := make(chan os.Signal)
			go func() {
				<-signalChan
				if err := di.Terminate(); err != nil {
					util.Die("termination failed:", err)
				} else {
					os.Exit(0)
				}
			}()
			signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
		},
		Run: func(cmd *cobra.Command, args []string) {
			addr := os.Getenv("CTRLT_ADDRESS")
			if addr == "" {
				addr = defaultAddr
			}
			_, _ = fmt.Fprintln(os.Stdout, "starting server at", addr)
			if err := (&http.Server{Addr: addr}).ListenAndServe(); err != nil {
				util.Die("unable to start server:", err)
			}
		},
	}
}

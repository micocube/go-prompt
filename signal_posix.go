//go:build !windows
// +build !windows

package prompt

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/micocube/go-prompt/internal/debug"
)

func (p *Prompt) handleSignals(exitCh chan int, winSizeCh chan *WinSize, stop chan struct{}) {
	in := p.in
	sigCh := make(chan os.Signal, 1)
	signal.Notify(
		sigCh,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGWINCH,
	)

	for {
		select {
		case <-stop:
			debug.Log("stop handleSignals")
			return
		case s := <-sigCh:
			switch s {

			case syscall.SIGTERM: // kill -SIGTERM XXXX
				debug.Log("Catch SIGTERM")
				exitCh <- 1

			case syscall.SIGQUIT: // kill -SIGQUIT XXXX
				debug.Log("Catch SIGQUIT")
				exitCh <- 0

			case syscall.SIGWINCH:
				debug.Log("Catch SIGWINCH")
				winSizeCh <- in.GetWinSize()
			}
		}
	}
}

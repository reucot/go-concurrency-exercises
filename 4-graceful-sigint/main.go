//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

//******************** - Добавленный код

package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//********************
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	//********************

	// Create a process
	proc := MockProcess{}

	// Run the process (blocking)
	go proc.Run()

	//********************
	exit_chan := make(chan int)
	go func() {
		for {
			s := <-sigChan
			switch s {

			case syscall.SIGINT:
				proc.Stop()

			case syscall.SIGTERM:
				exit_chan <- 0

			case syscall.SIGQUIT:
				exit_chan <- 0

			default:
				exit_chan <- 1
			}
		}
	}()
	exitCode := <-exit_chan
	os.Exit(exitCode)
	//********************
}

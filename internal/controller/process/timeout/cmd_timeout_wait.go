package timeout

import (
	"errors"
	"os/exec"
	"time"
)

var (
	ErrTimeout = errors.New("Cmd wait timeout reached")
)

type CmdTimeoutWait struct {
}

func NewCmdTimeoutWait() CmdTimeoutWait {
	return newCmdTimeoutWait()
}

func newCmdTimeoutWait() CmdTimeoutWait {
	return CmdTimeoutWait{}
}

func (cmdTimeoutWait *CmdTimeoutWait) Wait(cmd *exec.Cmd, timeoutSeconds time.Duration) error {
	processIsDone := make(chan error, 1)

	go func() {
		processIsDone <- cmd.Wait()
	}()

	select {
	case err := <-processIsDone:
		return err
	case <-time.After(timeoutSeconds):
		return ErrTimeout
	}
}

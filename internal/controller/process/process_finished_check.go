package process

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
	"time"
	"vbox-vm-setup/internal/controller/process/timeout"
)

type CurrentTimeFunc func() time.Time

type ProcessFinishedCheck interface {
	EnsureProcessFinished(cmd *exec.Cmd) error
}

type ProcessFinishedCheckImpl struct {
	cmdTimeoutWait timeout.CmdTimeoutWait
}

func NewProcessFinishedCheck() ProcessFinishedCheck {
	return newProcessFinishedCheck(timeout.NewCmdTimeoutWait())
}

func newProcessFinishedCheck(waitWithTimeout timeout.CmdTimeoutWait) ProcessFinishedCheck {
	return &ProcessFinishedCheckImpl{cmdTimeoutWait: waitWithTimeout}
}

func (processFinishedCheck *ProcessFinishedCheckImpl) EnsureProcessFinished(cmd *exec.Cmd) error {
	err := processFinishedCheck.cmdTimeoutWait.Wait(cmd, timeout.DEFAULT_COMMAND_TIMEOUT)

	if err == timeout.ErrTimeout {
		log.Info().Msgf("Process [%d] did not finish as expected and is still running", cmd.Process.Pid)
		return err
	} else if err != nil {
		return err
	} else {
		if cmd.ProcessState.ExitCode() == 0 {
			log.Info().Msgf("Process finished successfully")
			return nil
		} else {
			return fmt.Errorf("process failed with exit code %d", cmd.ProcessState.ExitCode())
		}
	}
}

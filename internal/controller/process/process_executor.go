package process

import (
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
)

type ProcessExecutor interface {
	ExecuteProcess(ctx ProcessContext) error
}

type ProcessExecutorImpl struct {
	processFinishedCheck ProcessFinishedCheck
}

func NewProcessExecutor() ProcessExecutor {
	return newProcessExecutor(NewProcessFinishedCheck())
}

func newProcessExecutor(processFinishedCheck ProcessFinishedCheck) ProcessExecutor {
	return &ProcessExecutorImpl{processFinishedCheck: processFinishedCheck}
}

func (processExecutor ProcessExecutorImpl) ExecuteProcess(ctx ProcessContext) error {
	cmd := exec.Command(ctx.Executable, ctx.Args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	log.Debug().
		Str("executable", ctx.Executable).
		Strs("arguments", ctx.Args).
		Msgf("Executing command")

	if err := cmd.Start(); err != nil {
		return err
	}

	return processExecutor.processFinishedCheck.EnsureProcessFinished(cmd)
}

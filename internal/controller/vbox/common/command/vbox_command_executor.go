package command

import (
	"create-vbox-vm/internal/controller/configuration"
	"create-vbox-vm/internal/controller/process"
	"create-vbox-vm/internal/controller/vbox/common/command/ds"
	"create-vbox-vm/internal/controller/vbox/common/command/utils"
)

type VBoxCommandExecutor interface {
	Execute(configuration configuration.Configuration, commandDefinition ds.VBoxManageCommand) error
}

type VBoxCommandExecutorImpl struct {
	processExecutor    process.ProcessExecutor
	vBoxManageProvider utils.VBoxManageProvider
}

func NewVBoxCommandExecutor() VBoxCommandExecutor {
	return newVBoxCommandExecutor(process.NewProcessExecutor(), utils.NewVBoxManageProviderImpl())
}

func newVBoxCommandExecutor(processExecutor process.ProcessExecutor, vBoxManageProvider utils.VBoxManageProvider) VBoxCommandExecutor {
	return VBoxCommandExecutorImpl{processExecutor: processExecutor, vBoxManageProvider: vBoxManageProvider}
}

func (executor VBoxCommandExecutorImpl) Execute(configuration configuration.Configuration, commandDefinition ds.VBoxManageCommand) error {
	return executor.processExecutor.ExecuteProcess(process.ProcessContext{
		executor.vBoxManageProvider.GetVBoxManage(configuration),
		append([]string{commandDefinition.Name}, commandDefinition.Args...),
	})
}

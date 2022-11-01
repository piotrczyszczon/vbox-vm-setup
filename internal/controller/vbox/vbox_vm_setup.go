package vbox

import (
	"github.com/rs/zerolog/log"
	"vbox-vm-setup/internal/controller/configuration"
	"vbox-vm-setup/internal/controller/vbox/create"
	"vbox-vm-setup/internal/controller/vbox/install"
)

type VBoxVmSetup interface {
	PrepareNewVm(configuration configuration.Configuration) error
}

type VBoxVmSetupImpl struct {
	vboxVmCreationFacade create.VBoxVmCreationFacade
	vboxInstallOSFacade  install.VBoxInstallOSFacade
}

func NewVBoxVmSetup() VBoxVmSetup {
	return newVBoxVmSetup(create.NewVBoxVmCreationFacade(), install.NewVBoxInstallOSFacade())
}

func newVBoxVmSetup(vboxVmCreationFacade create.VBoxVmCreationFacade, vboxInstallOSFacade install.VBoxInstallOSFacade) VBoxVmSetup {
	return VBoxVmSetupImpl{vboxVmCreationFacade: vboxVmCreationFacade, vboxInstallOSFacade: vboxInstallOSFacade}
}

func (vboxVmSetup VBoxVmSetupImpl) PrepareNewVm(configuration configuration.Configuration) error {
	err := vboxVmSetup.vboxVmCreationFacade.Execute(configuration)
	if err != nil {
		log.Err(err).Msgf("Failed to Create VM")
		return err
	}

	err = vboxVmSetup.vboxInstallOSFacade.Execute(configuration)
	if err != nil {
		log.Err(err).Msgf("Failed to Install OS")
		return err
	}

	return nil
}

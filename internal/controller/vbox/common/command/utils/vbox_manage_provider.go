package utils

import (
	"vbox-vm-setup/internal/controller/configuration"
)

type VBoxManageProvider interface {
	GetVBoxManage(configuration configuration.Configuration) string
}

type VBoxManageProviderImpl struct {
}

func NewVBoxManageProviderImpl() VBoxManageProvider {
	return newVBoxManageProviderImpl()
}

func newVBoxManageProviderImpl() VBoxManageProvider {
	return &VBoxManageProviderImpl{}
}

func (provider VBoxManageProviderImpl) GetVBoxManage(configuration configuration.Configuration) string {
	if len(configuration.VBoxManage) != 0 {
		return configuration.VBoxManage
	} else {
		return "VBoxManage"
	}
}

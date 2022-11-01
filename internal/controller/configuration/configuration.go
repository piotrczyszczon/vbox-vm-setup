package configuration

type Configuration struct {
	VBoxManage string  `yaml:"vbox-manage"`
	General    General `yaml:"general"`
	Create     Create  `yaml:"create"`
	Install    Install `yaml:"install"`
}

type General struct {
	Name        string
	IsoFilePath string `yaml:"iso-file-path"`
}

type Create struct {
	Predefined PredefinedCreate `yaml:"predefined"`
	Raw        Raw              `yaml:"raw"`
}

type PredefinedCreate struct {
	VmsHome            string `yaml:"vms-home"`
	OsType             string `yaml:"os-type" default:"Ubuntu_64"`
	Cpus               string `yaml:"cpus" default:"1"`
	Memory             string `yaml:"memory" default:"2048"`
	VirtualMemory      string `yaml:"virtual-memory" default:"128"`
	GraphicsController string `yaml:"graphics-controller" default:"vmsvga"`
	RtcUseUtc          string `yaml:"rtc-use-utc" default:"on"`
	ClipboardMode      string `yaml:"clipboard-mode" default:"bidirectional"`
	DragAndDrop        string `yaml:"drag-and-drop" default:"bidirectional"`
	Pae                string `yaml:"pae" default:"off"`
	HddSize            string `yaml:"hdd-size" default:"10240"`
	NatRule            string `yaml:"nat-rule" default:""`
}

type Raw struct {
	Commands []RawCommandDefinition `yaml:"commands"`
}

type RawCommandDefinition struct {
	Name string
	Args []string
}

type Install struct {
	Predefined PredefinedInstall `yaml:"predefined"`
	Raw        Raw               `yaml:"raw"`
}

type PredefinedInstall struct {
	FullUserName       string      `yaml:"full-user-name" default:"MyUser"`
	Credentials        Credentials `yaml:"credentials"`
	TimeZone           string      `yaml:"time-zone" default:"UTC"`
	PostInstallCommand string      `yaml:"post-install-command" default:""`
}

type Credentials struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

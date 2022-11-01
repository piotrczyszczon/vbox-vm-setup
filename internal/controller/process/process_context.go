package process

type Args []string

type ProcessContext struct {
	Executable string
	Args       Args
}

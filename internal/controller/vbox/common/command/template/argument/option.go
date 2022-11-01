package argument

import "fmt"

type Option struct {
	Name string
}

func Opt(name string) Option {
	return Option{Name: name}
}

func (option Option) AsString() string {
	return fmt.Sprintf("--%s", option.Name)
}

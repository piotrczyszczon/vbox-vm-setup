package argument

import "fmt"

type Argument struct {
	Name  string
	Value string
}

func Arg(name string, value string) Argument {
	return Argument{Name: name, Value: value}
}

func (argument Argument) AsString() string {
	return fmt.Sprintf("--%s=%s", argument.Name, argument.Value)
}

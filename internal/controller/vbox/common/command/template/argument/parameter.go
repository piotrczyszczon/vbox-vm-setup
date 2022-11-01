package argument

import "fmt"

type Parameter struct {
	Name string
}

func Param(name string) Parameter {
	return Parameter{Name: name}
}

func (parameter Parameter) AsString() string {
	return fmt.Sprintf("%s", parameter.Name)
}

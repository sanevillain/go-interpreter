package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnclosedEnvironemnt(outer *Environment) *Environment {
	return nil
}

func NewEnvironment() *Environment {
	return nil
}

func (e *Environment) Get(name string) (Object, bool) {
	return nil, false
}

func (e *Environment) Set(name string, val Object) Object {
	return nil
}

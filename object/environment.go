package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnclosedEnvironemnt(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	return &Environment{
		store: map[string]Object{},
	}
}

func (e *Environment) Get(name string) (Object, bool) {
	val, ok := e.store[name]
	if !ok && e.outer != nil {
		val, ok = e.outer.Get(name)
	}

	return val, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

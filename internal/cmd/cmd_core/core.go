package cmd_core

type Commandable interface {
	Exec(arg string, uuid uint64) (string, error)
	Use() string
	Desc() string
	Restricted() bool
	ModsOnly() bool
}

// DefaultCommand
// implement `Commandable` with composite:
//
//	type ExampleCommand struct { cmd_core.DefaultCommand }
//
// required methods: `Exec()`, `Use()`, `Desc()`
// default methods: `Restricted()`, `ModsOnly()`, returns bool by default
// override the required methods, and default methods, if necessary
type DefaultCommand struct{}

func (d DefaultCommand) Exec(_ string, _ uint64) (string, error) {
	panic("unimplemented")
}
func (d DefaultCommand) Use() string      { panic("unimplemented") }
func (d DefaultCommand) Desc() string     { panic("unimplemented") }
func (d DefaultCommand) Restricted() bool { return false }
func (d DefaultCommand) ModsOnly() bool   { return false }

const MaxLayouts = 20

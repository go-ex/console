package contract

type Console interface {
	Init()
	HasCommand() bool
	RunCommand() error
	RunCommandHelp(name string, configure Configure) error
	RunHelp()
}

type MapCommand struct {
	Command       Command
	CommandConfig Configure
}

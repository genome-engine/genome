package command

import "os/exec"

type DefaultCommandExecutor struct {
	command *exec.Cmd
}

func New(executor string, args []string) *DefaultCommandExecutor {
	return &DefaultCommandExecutor{exec.Command(executor, args...)}
}

func (command *DefaultCommandExecutor) Execute() error {
	return command.command.Run()
}

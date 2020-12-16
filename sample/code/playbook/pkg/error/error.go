package error

import "playbook/pkg/output"

func FailHandleCommand(err error)  {
	if err != nil {
		output.PrintCliError(err)
	}
}
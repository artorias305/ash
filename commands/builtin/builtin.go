package builtin

import "slices"

var BuiltinCommands = []string{"echo", "type", "exit", "pwd", "cd"}

func IsBuiltin(command string) bool {
	return slices.Contains(BuiltinCommands, command)
}

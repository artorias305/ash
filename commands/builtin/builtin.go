package builtin

import "slices"

var BuiltinCommands = []string{"echo", "type", "exit", "pwd"}

func IsBuiltin(command string) bool {
	return slices.Contains(BuiltinCommands, command)
}

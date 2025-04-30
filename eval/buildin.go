package eval

import (
	"fmt"

	"github.com/Serein-sz/knife/environment"
)

var builtins = map[string]*environment.Builtin{
	"print": {Name: "print", Function: Print},
}

func Print(args ...environment.Object) environment.Object {
	for i, a := range args {
		fmt.Print(a.Inspect())
		if i != len(args)-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
	return nil
}

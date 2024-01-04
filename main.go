/*
Copyright © 2022 dd84ai <dd84ai@gmail.com>

A set of tools for gaming community Freelancer Discovery
in order to be more productive during its development configuration
*/
package main

import (
	"darklint/cmd"
	_ "darklint/fldarklint/settings" // bootstrapp settings for initialization
)

func main() {
	cmd.Execute()
}

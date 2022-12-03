/*
Copyright © 2022 dd84ai <dd84ai@gmail.com>

A set of tools for gaming community Freelancer Discovery
in order to be more productive during its development configuration
*/
package main

import (
	"darktool/cmd"
	_ "darktool/tools/logger" // bootstrapped logger
)

func main() {
	cmd.Execute()
}

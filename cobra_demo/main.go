/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

/*
Initalizing a module
	1. Create a new directory
	2. cd into that directory
	3. run go mod init <MODNAME>

Initalizing a Cobra CLI application
	1. cd $HOME/code/myapp
	2. cobra-cli init
	3. go run main.go

Add commands to a project
	cobra-cli add serve
	cobra-cli add config
	cobra-cli add create -p 'configCmd'
*/

import "cobra_demo/cmd"

func main() {
	cmd.Execute()
}

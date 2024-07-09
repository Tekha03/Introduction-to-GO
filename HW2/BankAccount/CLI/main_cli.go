package main

import "BankAccount/cli"

func main() {
	command := &cli.Command{}
	command.Init()
	command.Execute()
}

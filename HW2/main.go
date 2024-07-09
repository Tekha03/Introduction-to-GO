package main

import "BankAccount/bank_server"

func main() {
	r: = bank_server.SetupRouter()
	r.Run(":8000")
}

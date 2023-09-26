package main

import "fmt"

type Account struct {
	accountId int
	balance   int
	name      string
}

type ManagerAccount struct {
	Account
}

func (m Account) Person() {
	fmt.Printf("Accounts id : %d\nBalance : %d\nName : %s\n\n", m.accountId, m.balance, m.name)
}

func main() {
	p1 := Account{1, 23000, "Pratik"}
	p2 := Account{2, 25000, "Aditya"}

	p1.Person()
	p2.Person()
}

package main

import "fmt"

type bank struct {
	name    string
	menu    []string
	manager Manager
	users   []*User
}

type logger interface {
	Login(id, pass string)
}

//NewBank func
func NewBank(name string) *bank {

	return &bank{
		name: name,
		menu: []string{"Create Account", "Login"},
		manager: Manager{
			id:   "admin",
			pass: "0000",
		},
	}
}

func (b *bank) PrintBankName() {
	fmt.Println("\t" + b.name + "\n\t````````````")
}

func (b *bank) DisplayMenu() {
	fmt.Println("\t  Options")
	fmt.Println("\t  ```````")
	for idx, item := range b.menu {
		fmt.Println(idx+1, ".", item)
	}
	fmt.Println("0 . Exit")

}

func (b *bank) SelectOption(s string) string {
	var choice string
	fmt.Print(s)
	_, err := fmt.Scanln(&choice)

	if err != nil {
		return b.SelectOption(s)
	}

	return choice

}

func (b *bank) OpenAccount() {

	var id, pass, name string
	fmt.Print("Enter your id : ")
	fmt.Scanln(&id)
	fmt.Print("Enter your pass : ")
	fmt.Scanln(&pass)
	fmt.Print("Enter your name : ")
	fmt.Scanln(&name)

	user := b.manager.CreateAccount(id, pass, name)
	b.users = append(b.users, user)
}

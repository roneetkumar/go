package main

//Manager Struct
type Manager struct {
	id, pass string
}

//CreateAccount fnc
func (m Manager) CreateAccount(id, pass, name string) *User {

	acc := []Account{
		{
			bal: 0,
			typ: "Savings",
		},
	}

	return &User{
		id:       id,
		pass:     pass,
		name:     name,
		accounts: acc,
	}

}

//AddAccount func
func (m Manager) AddAccount(id, typ string, b *bank) {

	acc := Account{
		bal: 0,
		typ: typ,
	}

	for _, user := range b.users {
		if user.id == id {
			user.accounts = append(user.accounts, acc)
		}
	}
}

//RemoveAccount func
func (m Manager) RemoveAccount(id, typ string, b *bank) {

	for _, user := range b.users {
		if user.id == id {
			for i, acc := range user.accounts {
				if acc.typ == typ {
					user.accounts = append(user.accounts[:i], user.accounts[i+1:]...)
				}
			}
		}
	}
}

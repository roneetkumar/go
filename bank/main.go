package main

func main() {

	bank := NewBank("The New Bank")

	bank.PrintBankName()
	bank.DisplayMenu()

	for {
		option := bank.SelectOption("Select the option from menu: ")

		switch option {
		case "1":
			bank.OpenAccount()
		case "2":

		default:

		}

	}
}

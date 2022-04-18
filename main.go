package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
* createBill - function that create bill with user input name and returns the empty bill with name
 */
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	//create bill
	b := newBill(name)
	fmt.Println("Created the bill - ", name)

	return b
}

/*
* promptOptions - function that takes input to add item, save bill and add tip to a bill
* @bill: bill struct to be updated with prompts
 */
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("\nOptions:\n\ta - add item \n\tu - update bill item \n\td - delete bill item \n\ts - save bill \n\tt - add tip \n\tp - print bill\n\tq - end program \n\nChoose option: ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number...")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item added -", name, price)
		promptOptions(b)
	case "u":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number...")
			promptOptions(b)
		}
		b.updateItem(name, p)

		fmt.Println("item updated -", name, price)
		promptOptions(b)
	case "d":
		name, _ := getInput("Name of item to be deleted: ", reader)
		b.deleteItem(name)

		fmt.Println("item deleted -", name)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number...")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip has been updated to", tip)
		promptOptions(b)

	case "s":
		b.save()

		fmt.Println("bill has been saved as", b.name)
		promptOptions(b)
	case "p":
		fs, _ := b.format()

		fmt.Println(fs)
		promptOptions(b)
	case "q":
		fmt.Println("Ending program now")
		return
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}

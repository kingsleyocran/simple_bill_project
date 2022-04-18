package main

import "fmt"

var men int = 345

/*
* bill - defines a struct that has:
* @name: name of the bill recipient
* @items:  map of items on the bill
* @tip: tip on the bill
 */
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

/*
* newBill - function to create new bill
* @name: name of the bill
 */
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

/*
* format - receiver function that formats output of bill
 */
func (b *bill) format() (string, float64) {
	var total float64 = 0

	fs := "Bill Breakdown \n"

	//add total output formating
	fs += fmt.Sprintf("%-25v ...%v\n", "name:", (*b).name)

	//list all items in output format
	for k, v := range (*b).items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	//add total output formating
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", (*b).tip)

	//add total output formating
	fs += fmt.Sprintf("\n%-25v ...$%0.2f\n", "total:", (total + (*b).tip))

	return fs, total
}

/*
* updateTip - receiver function that adds a tip to the bill
*@tip: tip to be added
 */
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}

/*
* addItem - receiver function that adds a tip to the bill
*@name: name of the item
*@price: price of the item
 */
func (b *bill) addItem(name string, price float64) {
	(*b).items[name] = price
}

func (b *bill) updateItem(name string, price float64) {
	(*b).items[name] = price
}

func (b *bill) deleteItem(name string) {
	delete((*b).items, name)
}

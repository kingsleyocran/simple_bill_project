package bill

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

package main

import (
	"fmt"
	"os"
)

/*
* save - receiver function that writes the bill to a .txt file
 */
func (b *bill) save() {
	fs, _ := b.format()
	data := []byte(fs)
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")
}

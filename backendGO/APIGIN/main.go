package main

import (
	"fmt"
	"strconv"
	"strings"
	"umachay/database"
)

func main() {
	fmt.Println("Hello, World!")
	database.Conection()
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	fmt.Println(StockList(b, c))
}

func StockList(listArt []string, listCat []string) string {
	var result string
	for i := 0; i < len(listArt); i++ {
		for j := 0; j < len(listCat); j++ {

			var art = strings.Split(listArt[i], " ")[1]
			var cat = listCat[j]
			result += art + " - " + cat + "\n"
			var a:= 
			fmt.Println(strconv.Atoi(art))
		}
	}
	return result
}

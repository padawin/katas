package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*GildedRoseItem{
		{Item{"+5 Dexterity Vest", 10, 20}, nil},
		{Item{"Aged Brie", 2, 0}, nil},
		{Item{"Elixir of the Mongoose", 5, 7}, nil},
		{Item{"Sulfuras, Hand of Ragnaros", 0, 80}, nil},
		{Item{"Sulfuras, Hand of Ragnaros", -1, 80}, nil},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20}, nil},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 10, 49}, nil},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 49}, nil},
		{Item{"Conjured Mana Cake", 3, 6}, nil}, // <-- :O
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("name, sellIn, quality")
		for i := 0; i < len(items); i++ {
			fmt.Println(items[i])
		}
		fmt.Println("")
		UpdateQuality(items)
	}
}

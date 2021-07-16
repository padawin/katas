package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*GildedRoseItem{
		{Item{"+5 Dexterity Vest", 10, 20}, false, ""},
		{Item{"Aged Brie", 2, 0}, false, "item_update_brie"},
		{Item{"Elixir of the Mongoose", 5, 7}, false, ""},
		{Item{"Sulfuras, Hand of Ragnaros", 0, 80}, false, "item_update_sulfura"},
		{Item{"Sulfuras, Hand of Ragnaros", -1, 80}, false, "item_update_sulfura"},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20}, false, "item_update_backstage_pass"},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 10, 49}, false, "item_update_backstage_pass"},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 49}, false, "item_update_backstage_pass"},
		{Item{"Conjured Mana Cake", 3, 6}, true, ""}, // <-- :O
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

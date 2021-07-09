package main

type Item struct {
	name            string
	sellIn, quality int
}

type GildedRoseItem struct {
	Item
}

func UpdateQuality(items []*GildedRoseItem) {
	for i := 0; i < len(items); i++ {
		item := items[i]
		var qualityEvolution int
		var sellInEvolution int

		// Items specific cases
		// TODO move out
		if item.name == "Aged Brie" {
			qualityEvolution = 1
			sellInEvolution = -1
		} else if item.name == "Sulfuras, Hand of Ragnaros" {
			qualityEvolution = 0
			sellInEvolution = 0
		} else if item.name == "Backstage passes to a TAFKAL80ETC concert" {
			if item.sellIn <= 0 {
				qualityEvolution = -item.quality
			} else if item.sellIn <= 5 {
				qualityEvolution = 3
			} else if item.sellIn <= 10 {
				qualityEvolution = 2
			} else {
				qualityEvolution = 1
			}
			sellInEvolution = -1
		} else {
			// Generic case
			qualityEvolution = -1
			sellInEvolution = -1
		}

		if item.sellIn <= 0 {
			qualityEvolution *= 2
		}
		item.quality += qualityEvolution
		if qualityEvolution > 0 && item.quality > 50 {
			item.quality = 50
		} else if item.quality < 0 {
			item.quality = 0
		}
		item.sellIn += sellInEvolution
	}
}

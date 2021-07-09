package main

type Item struct {
	name            string
	sellIn, quality int
}

type ItemQualityEvolution int
type ItemSellInEvolution int

type GildedRoseItem struct {
	Item
	update *func(item *GildedRoseItem) (ItemQualityEvolution, ItemSellInEvolution)
}

func UpdateQuality(items []*GildedRoseItem) {
	for i := 0; i < len(items); i++ {
		item := items[i]
		var qualityEvolution ItemQualityEvolution
		var sellInEvolution ItemSellInEvolution

		if item.update != nil {
			qualityEvolution, sellInEvolution = (*item.update)(item)
		} else {
			// Generic case
			qualityEvolution = -1
			sellInEvolution = -1
		}

		if item.sellIn <= 0 {
			qualityEvolution *= 2
		}
		item.quality += int(qualityEvolution)
		if qualityEvolution > 0 && item.quality > 50 {
			item.quality = 50
		} else if item.quality < 0 {
			item.quality = 0
		}
		item.sellIn += int(sellInEvolution)
	}
}

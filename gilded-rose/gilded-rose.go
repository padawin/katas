package main

import (
	"fmt"
)

type Item struct {
	name            string
	sellIn, quality int
}

type ItemQualityEvolution int
type ItemSellInEvolution int

type GildedRoseItem struct {
	Item
	isConjured   bool
	updateScript string
}

const (
	defaultQualityEvolution = -1
	defaultSellInEvolution  = -1
)

func UpdateQuality(items []*GildedRoseItem) {
	for i := 0; i < len(items); i++ {
		item := items[i]
		qualityEvolution, sellInEvolution := getQualityAndSellInEvolution(item)
		updateItemQualityAndSellIn(item, qualityEvolution, sellInEvolution)
	}
}

// getQualityAndSellInEvolution returns the delta of which the item's quality
// and sellIn must change
func getQualityAndSellInEvolution(item *GildedRoseItem) (ItemQualityEvolution, ItemSellInEvolution) {
	var qualityEvolution ItemQualityEvolution
	var sellInEvolution ItemSellInEvolution
	var err error

	qualityEvolution, sellInEvolution, err = executeItemUpdateScript(item, defaultQualityEvolution, defaultSellInEvolution)
	if item.updateScript != "" && err != nil {
		fmt.Println(err)
	}

	if item.sellIn <= 0 {
		qualityEvolution *= 2
	}

	if item.isConjured {
		qualityEvolution *= 2
	}

	return qualityEvolution, sellInEvolution
}

// updateItemQualityAndSellIn sets the new values of the item's quality and
// sellIn, based on the provided evolutions, and ensure the resulting values are
// respecting the product's specs:
// - The Quality of an item is never more than 50 (which is read as "never
// increases above 50. An item of quality 100 with a quality decrease of 1 would
// have a new quality of 99)
// - The Quality of an item is never negative
func updateItemQualityAndSellIn(item *GildedRoseItem, qualityEvolution ItemQualityEvolution, sellInEvolution ItemSellInEvolution) {
	item.quality += int(qualityEvolution)
	if qualityEvolution > 0 && item.quality > 50 {
		item.quality = 50
	} else if item.quality < 0 {
		item.quality = 0
	}
	item.sellIn += int(sellInEvolution)
}

package main

func updateBrie(item *GildedRoseItem) (ItemQualityEvolution, ItemSellInEvolution) {
	return 1, -1
}

func updateSulfuras(item *GildedRoseItem) (ItemQualityEvolution, ItemSellInEvolution) {
	return 0, 0
}

func updateBackstagePass(item *GildedRoseItem) (ItemQualityEvolution, ItemSellInEvolution) {
	var qualityEvolution ItemQualityEvolution
	if item.sellIn <= 0 {
		qualityEvolution = ItemQualityEvolution(-item.quality)
	} else if item.sellIn <= 5 {
		qualityEvolution = 3
	} else if item.sellIn <= 10 {
		qualityEvolution = 2
	} else {
		qualityEvolution = 1
	}
	return qualityEvolution, -1
}

package main

import (
	"fmt"

	python "github.com/sbinet/go-python"
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

func executeItemUpdateScript(item *GildedRoseItem, defaultQualityEvolution ItemQualityEvolution, defaultSellInEvolution ItemSellInEvolution) (ItemQualityEvolution, ItemSellInEvolution, error) {
	python.Initialize()
	defer python.Finalize()

	itemUpdateModule := python.PyImport_ImportModule(item.updateScript)
	if itemUpdateModule == nil {
		return defaultQualityEvolution, defaultSellInEvolution, fmt.Errorf("Error importing module: %s", item.updateScript)
	}

	updateFunc := itemUpdateModule.GetAttrString("update")
	if updateFunc == nil {
		return defaultQualityEvolution, defaultSellInEvolution, fmt.Errorf("Error importing function `update`")
	}

	args := python.PyTuple_New(1)
	brieDict := python.PyDict_New()
	python.PyDict_SetItem(brieDict, python.PyString_FromString("name"), python.PyString_FromString(item.name))
	python.PyDict_SetItem(brieDict, python.PyString_FromString("quality"), python.PyInt_FromLong(item.quality))
	python.PyDict_SetItem(brieDict, python.PyString_FromString("sellIn"), python.PyInt_FromLong(item.sellIn))
	python.PyTuple_SetItem(args, 0, brieDict)
	res := updateFunc.CallObject(args)
	if !python.PyTuple_Check(res) || python.PyTuple_Size(res) != 2 {
		return defaultQualityEvolution, defaultSellInEvolution, fmt.Errorf("update must return a tuple of 2 elements")
	}
	qe := python.PyTuple_GetItem(res, 0)
	se := python.PyTuple_GetItem(res, 1)
	if !python.PyInt_Check(qe) {
		return defaultQualityEvolution, defaultSellInEvolution, fmt.Errorf("the first returned element of update must be an int")
	} else if !python.PyInt_Check(se) {
		return defaultQualityEvolution, defaultSellInEvolution, fmt.Errorf("the second returned element of update must be an int")
	}
	return ItemQualityEvolution(python.PyInt_AsLong(qe)), ItemSellInEvolution(python.PyInt_AsLong(se)), nil
}

package main

import (
	"fmt"

	python "github.com/sbinet/go-python"
)

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
	itemDict := python.PyDict_New()
	python.PyDict_SetItem(itemDict, python.PyString_FromString("quality"), python.PyInt_FromLong(item.quality))
	python.PyDict_SetItem(itemDict, python.PyString_FromString("sellIn"), python.PyInt_FromLong(item.sellIn))
	python.PyTuple_SetItem(args, 0, itemDict)
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

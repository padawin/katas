package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_SizeIsPreserved(t *testing.T) {
	g := NewGomegaWithT(t)
	var items = []*GildedRoseItem{
		{Item{"foo", 0, 0}, false, nil},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(1))
}

func Test_SellInAndQualityAreUpdated(t *testing.T) {
	g := NewGomegaWithT(t)
	var items = []*GildedRoseItem{
		{Item{"foo", 17, 42}, false, nil},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(1))
	g.Expect(items[0].sellIn).To(Equal(16))
	g.Expect(items[0].quality).To(Equal(41))
}

func Test_QualityDecreasesTwiceAsFastWhenExpired(t *testing.T) {
	g := NewGomegaWithT(t)
	var items = []*GildedRoseItem{
		{Item{"foo", 0, 42}, false, nil},
		{Item{"foo", -1, 42}, false, nil},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(2))
	g.Expect(items[0].sellIn).To(Equal(-1))
	g.Expect(items[0].quality).To(Equal(40))
	g.Expect(items[1].sellIn).To(Equal(-2))
	g.Expect(items[1].quality).To(Equal(40))
}

func Test_QualityIsNeverNegative(t *testing.T) {
	g := NewGomegaWithT(t)
	var items = []*GildedRoseItem{
		{Item{"foo", 0, 1}, false, nil},
		{Item{"foo", -1, 0}, false, nil},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(2))
	g.Expect(items[0].sellIn).To(Equal(-1))
	g.Expect(items[0].quality).To(Equal(0))
	g.Expect(items[1].sellIn).To(Equal(-2))
	g.Expect(items[1].quality).To(Equal(0))
}

func Test_Conjured(t *testing.T) {
	g := NewGomegaWithT(t)
	var items = []*GildedRoseItem{
		{Item{"item with positive sellIn", 10, 17}, true, nil},
		{Item{"item with d-day sellIn", 0, 17}, true, nil},
		{Item{"item with negative sellIn", -1, 17}, true, nil},
		{Item{"item with low quality", -1, 1}, true, nil},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(4))
	g.Expect(items[0].sellIn).To(Equal(9))
	g.Expect(items[0].quality).To(Equal(15))
	g.Expect(items[1].sellIn).To(Equal(-1))
	g.Expect(items[1].quality).To(Equal(13))
	g.Expect(items[2].sellIn).To(Equal(-2))
	g.Expect(items[2].quality).To(Equal(13))
	g.Expect(items[3].sellIn).To(Equal(-2))
	g.Expect(items[3].quality).To(Equal(0))
}

// Specific cases

func Test_SpecificAgedBrieIncreasesQuality(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateBrie
	var items = []*GildedRoseItem{
		{Item{"Aged Brie", 17, 42}, false, &update},
		{Item{"Aged Brie", -1, 49}, false, &update},
		{Item{"Aged Brie", -1, 40}, false, &update},
		// Conjured
		{Item{"Aged Brie", 17, 42}, true, &update},
		{Item{"Aged Brie", -1, 40}, true, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(5))
	g.Expect(items[0].sellIn).To(Equal(16))
	g.Expect(items[0].quality).To(Equal(43))
	g.Expect(items[1].sellIn).To(Equal(-2))
	g.Expect(items[1].quality).To(Equal(50))
	g.Expect(items[2].sellIn).To(Equal(-2))
	g.Expect(items[2].quality).To(Equal(42))
	g.Expect(items[3].sellIn).To(Equal(16))
	g.Expect(items[3].quality).To(Equal(44))
	g.Expect(items[4].sellIn).To(Equal(-2))
	g.Expect(items[4].quality).To(Equal(44))
}

func Test_SpecificAgedBrieIncreasesQualityNeverGreaterThan50(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateBrie
	var items = []*GildedRoseItem{
		{Item{"Aged Brie", -1, 50}, false, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(1))
	g.Expect(items[0].sellIn).To(Equal(-2))
	g.Expect(items[0].quality).To(Equal(50))
}

func Test_SpecificSulfuraSellInAndQualityDoNotChange(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateSulfuras
	var items = []*GildedRoseItem{
		{Item{"Sulfuras, Hand of Ragnaros", 17, 42}, false, &update},
		// Conjured
		{Item{"Sulfuras, Hand of Ragnaros", 17, 42}, true, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(2))
	g.Expect(items[0].sellIn).To(Equal(17))
	g.Expect(items[0].quality).To(Equal(42))
	g.Expect(items[1].sellIn).To(Equal(17))
	g.Expect(items[1].quality).To(Equal(42))
}

func Test_SpecificBackstagePassesIncreaseInQuality(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateBackstagePass
	var items = []*GildedRoseItem{
		{Item{"Backstage passes to a TAFKAL80ETC concert", 17, 42}, false, &update},
		// Conjured
		{Item{"Backstage passes to a TAFKAL80ETC concert", 17, 42}, true, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(2))
	g.Expect(items[0].sellIn).To(Equal(16))
	g.Expect(items[0].quality).To(Equal(43))
	g.Expect(items[1].sellIn).To(Equal(16))
	g.Expect(items[1].quality).To(Equal(44))
}

func Test_SpecificBackstagePassesIncreaseInQualityButNotAbove50(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateBackstagePass
	var items = []*GildedRoseItem{
		{Item{"Backstage passes to a TAFKAL80ETC concert", 17, 50}, false, &update},
		// Conjured
		{Item{"Backstage passes to a TAFKAL80ETC concert", 17, 50}, true, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(2))
	g.Expect(items[0].sellIn).To(Equal(16))
	g.Expect(items[0].quality).To(Equal(50))
	g.Expect(items[1].sellIn).To(Equal(16))
	g.Expect(items[1].quality).To(Equal(50))
}

func Test_SpecificBackstagePassesIncreaseInQualityBy2When10DaysOrLessFromExpiry(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateBackstagePass
	var items = []*GildedRoseItem{
		{Item{"Backstage passes to a TAFKAL80ETC concert", 10, 42}, false, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 6, 42}, false, &update},
		// Conjured
		{Item{"Backstage passes to a TAFKAL80ETC concert", 10, 42}, true, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 6, 42}, true, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(4))
	g.Expect(items[0].sellIn).To(Equal(9))
	g.Expect(items[0].quality).To(Equal(44))
	g.Expect(items[1].sellIn).To(Equal(5))
	g.Expect(items[1].quality).To(Equal(44))
	g.Expect(items[2].sellIn).To(Equal(9))
	g.Expect(items[2].quality).To(Equal(46))
	g.Expect(items[3].sellIn).To(Equal(5))
	g.Expect(items[3].quality).To(Equal(46))
}

func Test_SpecificBackstagePassesIncreaseInQualityBy2When10DaysOrLessFromExpiryButNotAbove50(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateBackstagePass
	var items = []*GildedRoseItem{
		{Item{"Backstage passes to a TAFKAL80ETC concert", 10, 50}, false, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 10, 49}, false, &update},
		// Conjured
		{Item{"Backstage passes to a TAFKAL80ETC concert", 10, 50}, true, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 10, 49}, true, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(4))
	g.Expect(items[0].sellIn).To(Equal(9))
	g.Expect(items[0].quality).To(Equal(50))
	g.Expect(items[1].sellIn).To(Equal(9))
	g.Expect(items[1].quality).To(Equal(50))
	g.Expect(items[2].sellIn).To(Equal(9))
	g.Expect(items[2].quality).To(Equal(50))
	g.Expect(items[3].sellIn).To(Equal(9))
	g.Expect(items[3].quality).To(Equal(50))
}

func Test_SpecificBackstagePassesIncreaseInQualityBy3When5DaysOrLessFromExpiry(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateBackstagePass
	var items = []*GildedRoseItem{
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 42}, false, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 1, 42}, false, &update},
		// Conjured
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 42}, true, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 1, 42}, true, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(4))
	g.Expect(items[0].sellIn).To(Equal(4))
	g.Expect(items[0].quality).To(Equal(45))
	g.Expect(items[1].sellIn).To(Equal(0))
	g.Expect(items[1].quality).To(Equal(45))
	g.Expect(items[2].sellIn).To(Equal(4))
	g.Expect(items[2].quality).To(Equal(48))
	g.Expect(items[3].sellIn).To(Equal(0))
	g.Expect(items[3].quality).To(Equal(48))
}

func Test_SpecificBackstagePassesIncreaseInQualityBy3When5DaysOrLessFromExpiryButNotAbove50(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateBackstagePass
	var items = []*GildedRoseItem{
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 50}, false, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 49}, false, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 48}, false, &update},
		// Conjured
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 50}, true, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 49}, true, &update},
		{Item{"Backstage passes to a TAFKAL80ETC concert", 5, 48}, true, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(6))
	g.Expect(items[0].sellIn).To(Equal(4))
	g.Expect(items[0].quality).To(Equal(50))
	g.Expect(items[1].sellIn).To(Equal(4))
	g.Expect(items[1].quality).To(Equal(50))
	g.Expect(items[2].sellIn).To(Equal(4))
	g.Expect(items[2].quality).To(Equal(50))
	g.Expect(items[3].sellIn).To(Equal(4))
	g.Expect(items[3].quality).To(Equal(50))
	g.Expect(items[4].sellIn).To(Equal(4))
	g.Expect(items[4].quality).To(Equal(50))
	g.Expect(items[5].sellIn).To(Equal(4))
	g.Expect(items[5].quality).To(Equal(50))
}

func Test_SpecificBackstagePassesQualityBecomes0WhenExpired(t *testing.T) {
	g := NewGomegaWithT(t)
	update := updateBackstagePass
	var items = []*GildedRoseItem{
		{Item{"Backstage passes to a TAFKAL80ETC concert", 0, 42}, false, &update},
		// Conjured
		{Item{"Backstage passes to a TAFKAL80ETC concert", 0, 42}, true, &update},
	}

	UpdateQuality(items)

	g.Expect(len(items)).To(Equal(2))
	g.Expect(items[0].sellIn).To(Equal(-1))
	g.Expect(items[0].quality).To(Equal(0))
	g.Expect(items[1].sellIn).To(Equal(-1))
	g.Expect(items[1].quality).To(Equal(0))
}

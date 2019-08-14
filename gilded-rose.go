package main

type Item struct {
	name            string
	sellIn, quality int
}

const (
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	AgedBrie        = "Aged Brie"
)

func UpdateItems(items []*Item) {
	for _, item := range items {
		updateQuality(item)
		updateSellIn(item)
	}
}

func updateQuality(item *Item) {
	if item.name == BackstagePasses {
		updateBackstagePassesQuality(item)
		return
	}

	if item.name != AgedBrie {
		if item.name != Sulfuras && item.quality > 0 {
			item.quality = item.quality - 1
		}
	} else if item.quality < 50 {
		item.quality = item.quality + 1
	}

	if item.sellIn < 0 {
		if item.name != AgedBrie {
			if item.quality > 0 && item.name != Sulfuras {
				item.quality = item.quality - 1
			}
		} else if item.quality < 50 {
			item.quality = item.quality + 1
		}
	}

}

func updateBackstagePassesQuality(item *Item) {
	if item.sellIn < 0 {
		item.quality = 0
		return
	}

	if item.quality > 50 {
		return
	}

	if item.sellIn <= 5 {
		item.quality += 3
		return
	}

	if item.sellIn <= 10 {
		item.quality += 2
		return
	}

	item.quality += 1
}

func updateSellIn(item *Item) {
	if item.name != Sulfuras {
		item.sellIn = item.sellIn - 1
	}
}

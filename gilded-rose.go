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

func UpdateQuality(items []*Item) {
	for _, item := range items {
		if item.name != AgedBrie && item.name != BackstagePasses {
			if item.quality > 0 {
				if item.name != Sulfuras {
					item.quality = item.quality - 1
				}
			}
		} else {
			if item.quality < 50 {
				item.quality = item.quality + 1
				if item.name == BackstagePasses {
					if item.sellIn < 11 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
					if item.sellIn < 6 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
				}
			}
		}

		if item.name != Sulfuras {
			item.sellIn = item.sellIn - 1
		}

		if item.sellIn < 0 {
			if item.name != AgedBrie {
				if item.name != BackstagePasses {
					if item.quality > 0 {
						if item.name != Sulfuras {
							item.quality = item.quality - 1
						}
					}
				} else {
					item.quality = item.quality - item.quality
				}
			} else {
				if item.quality < 50 {
					item.quality = item.quality + 1
				}
			}
		}
	}
}

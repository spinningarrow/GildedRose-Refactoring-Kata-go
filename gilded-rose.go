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
	for i := 0; i < len(items); i++ {

		if items[i].name != AgedBrie && items[i].name != BackstagePasses {
			if items[i].quality > 0 {
				if items[i].name != Sulfuras {
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if items[i].name == BackstagePasses {
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		if items[i].name != Sulfuras {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != AgedBrie {
				if items[i].name != BackstagePasses {
					if items[i].quality > 0 {
						if items[i].name != Sulfuras {
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}
		}
	}

}

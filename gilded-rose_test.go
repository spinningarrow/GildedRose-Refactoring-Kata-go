package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gilded Rose", func() {
	var (
		name    string
		quality int
		sellIn  = 10
		item    = &Item{}
		items   = []*Item{
			item,
		}
	)

	JustBeforeEach(func() {
		item = &Item{name, sellIn, quality}
		items = []*Item{
			item,
		}
		UpdateQuality(items)
	})

	Context("when the item is not special", func() {
		BeforeEach(func() {
			name = "foo"
		})

		Context("when quality is above 0", func() {
			BeforeEach(func() {
				quality = 1
			})

			It("lowers quality", func() {
				Expect(item.quality).To(Equal(quality - 1))
			})
		})

		Context("when the sell by date has passed", func() {
			BeforeEach(func() {
				sellIn = -1
				quality = 10
			})

			It("degrades in quality twice as fast", func() {
				Expect(item.quality).To(Equal(quality - 2))
			})
		})

		Context("when the quality is 0", func() {
			BeforeEach(func() {
				quality = 0
			})

			It("does not decrease the quality", func() {
				Expect(item.quality).To(Equal(quality))
			})
		})

		It("lowers sellIn", func() {
			Expect(item.sellIn).To(Equal(sellIn - 1))
		})
	})

	Context("when the item is Aged Brie", func() {
		BeforeEach(func() {
			name = AgedBrie
			quality = 10
			sellIn = 10
		})

		Context("when quality is below 50", func() {
			BeforeEach(func() {
				quality = 10
			})

			It("increases in quality", func() {
				Expect(item.quality).To(Equal(quality + 1))
			})
		})

		Context("when quality is 50 or higher", func() {
			BeforeEach(func() {
				quality = 50
			})

			It("does not increase in quality", func() {
				Expect(item.quality).To(Equal(50))
			})
		})
	})

	Context("when the item is legendary", func() {
		BeforeEach(func() {
			name = Sulfuras
		})

		It("does not lower SellIn", func() {
			Expect(item.sellIn).To(Equal(sellIn))
		})

		It("does not lower quality", func() {
			Expect(item.quality).To(Equal(quality))
		})
	})

	Context("when the item is a backstage pass", func() {
		BeforeEach(func() {
			name = BackstagePasses
		})

		Context("when there are more than 10 days remaining", func() {
			BeforeEach(func() {
				sellIn = 20
				quality = 20
			})

			It("increases quality by 1", func() {
				Expect(item.quality).To(Equal(quality + 1))
			})

			It("decreases sellIn by 1", func() {
				Expect(item.sellIn).To(Equal(sellIn - 1))
			})
		})

		Context("when there are 10 days or less remaining", func() {
			BeforeEach(func() {
				sellIn = 10
				quality = 10
			})

			It("increases quality by 2", func() {
				Expect(item.quality).To(Equal(quality + 2))
			})
		})

		Context("when there are 5 days or less remaining", func() {
			BeforeEach(func() {
				sellIn = 5
				quality = 10
			})

			It("increases quality by 3", func() {
				Expect(item.quality).To(Equal(quality + 3))
			})
		})

		Context("when the concert is over", func() {
			BeforeEach(func() {
				sellIn = -1
			})

			It("drops quality to 0", func() {
				Expect(item.quality).To(Equal(0))
			})
		})
	})
})

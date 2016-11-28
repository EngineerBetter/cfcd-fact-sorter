package main

import (
	"sort"
)

func main() {
}

type Fact struct {
	Id          string
	Description string
}

type ItemFacts struct {
	ItemId string `yaml:"item_id"`
	Facts  []Fact
}

func (itemFacts ItemFacts) Len() int {
	return len(itemFacts.Facts)
}

func (itemFacts ItemFacts) Less(i, j int) bool {
	return itemFacts.Facts[i].Id < itemFacts.Facts[j].Id
}

func (itemFacts ItemFacts) Swap(i, j int) {
	itemFacts.Facts[i], itemFacts.Facts[j] = itemFacts.Facts[j], itemFacts.Facts[i]
}

type Items struct {
	Items []ItemFacts
}

func (items Items) Sort() {
	sort.Sort(items)
	for _, item := range items.Items {
		sort.Sort(item)
	}
}

func (items Items) Len() int {
	return len(items.Items)
}

func (items Items) Less(i, j int) bool {
	return items.Items[i].ItemId < items.Items[j].ItemId
}

func (items Items) Swap(i, j int) {
	items.Items[i], items.Items[j] = items.Items[j], items.Items[i]
}

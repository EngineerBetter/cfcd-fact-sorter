package main

import ()

type Fact struct {
	Id          string
	Description string
}

type ItemFacts struct {
	ItemId string `yaml:"item_id"`
	Facts  []Fact
}

type Items struct {
	Items []ItemFacts
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

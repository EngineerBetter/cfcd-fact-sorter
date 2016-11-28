package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	args := os.Args
	path := args[1]

	if !exists(path) {
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(1)
	}

	items := Items{}
	err = yaml.Unmarshal(bytes, &items)
	if err != nil {
		os.Exit(1)
	}

	items.Sort()
	outBytes, err := yaml.Marshal(&items)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("---")
	fmt.Println(string(outBytes))
}

func exists(path string) bool {
	_, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
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

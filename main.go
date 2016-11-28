package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
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

	prependedBytes := append([]byte("---\n"), outBytes...)
	ioutil.WriteFile(path, prependedBytes, 0644)
}

func exists(path string) bool {
	_, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

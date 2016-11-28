package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args

	var bytes []byte
	var err error

	if len(args) == 2 {
		path := args[1]

		if !exists(path) {
			os.Exit(1)
		}

		bytes, err = ioutil.ReadFile(path)
		if err != nil {
			os.Exit(1)
		}
	} else {
		if !terminal.IsTerminal(0) {
			bytes, _ = ioutil.ReadAll(os.Stdin)
		} else {
			os.Exit(1)
		}
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
	if len(args) == 2 {
		path := args[1]
		ioutil.WriteFile(path, prependedBytes, 0644)
	} else {
		fmt.Println(string(prependedBytes))
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

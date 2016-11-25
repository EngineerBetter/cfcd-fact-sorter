package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/EngineerBetter/cfcd-fact-sorter"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
)

var _ = Describe("sorting shizzle", func() {
	It("sorts YAML", func() {
		bytes, err := ioutil.ReadFile("fixtures/facts.yml")
		Ω(err).ShouldNot(HaveOccurred())

		items := Items{}
		err = yaml.Unmarshal(bytes, &items)
		sort.Sort(items)
		Ω(items.Items[0].ItemId).Should(Equal("FDAZ00000"))
	})
})

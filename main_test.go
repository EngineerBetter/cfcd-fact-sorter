package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"io/ioutil"
	"os/exec"
)

var cliPath string

var _ = Describe("cfcf-fact-sorter", func() {
	BeforeSuite(func() {
		var err error
		cliPath, err = Build("github.com/EngineerBetter/cfcd-fact-sorter")
		Ω(err).ShouldNot(HaveOccurred(), "Error building source")
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})

	Context("when the input file exists", func() {
		var sorted string

		BeforeEach(func() {
			bytes, err := ioutil.ReadFile("fixtures/sorted.yml")
			Ω(err).ShouldNot(HaveOccurred())
			sorted = string(bytes)
		})

		It("outputs sorted YAML to stdout", func() {
			command := exec.Command(cliPath, "fixtures/facts.yml")
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred(), "Error running CLI: "+cliPath)
			Eventually(session).Should((Exit(0)))
			Ω(session).Should(Say(sorted))
		})
	})
})

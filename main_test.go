package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cf-guardian/guardian/kernel/fileutils"

	"io/ioutil"
	"os/exec"
	"path/filepath"
)

var cliPath string
var fixturePath string

var _ = Describe("cfcd-fact-sorter", func() {
	var sorted string

	BeforeSuite(func() {
		var err error
		cliPath, err = Build("github.com/EngineerBetter/cfcd-fact-sorter")
		Ω(err).ShouldNot(HaveOccurred(), "Error building source")

		tmpDir, err := ioutil.TempDir("", "cfcd-fact-sorter")
		Ω(err).ShouldNot(HaveOccurred())
		fixturePath = filepath.Join(tmpDir, "facts.yml")

		fu := fileutils.New()
		err = fu.Copy(fixturePath, "fixtures/facts.yml")
		Ω(err).ShouldNot(HaveOccurred())

		bytes, err := ioutil.ReadFile("fixtures/sorted.yml")
		Ω(err).ShouldNot(HaveOccurred())
		sorted = string(bytes)
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})

	Context("when the input file exists", func() {
		It("outputs sorted YAML over the input file specified", func() {
			command := exec.Command(cliPath, fixturePath)
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred(), "Error running CLI: "+cliPath)
			Eventually(session).Should((Exit(0)))

			bytes, err := ioutil.ReadFile(fixturePath)
			actual := string(bytes)
			Ω(actual).Should(Equal(sorted))
		})
	})

	Context("when no arguments passed", func() {
		It("outputs sorted YAML to STDOUT", func() {
			command := exec.Command(cliPath)
			in, err := command.StdinPipe()
			Ω(err).ShouldNot(HaveOccurred())

			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred(), "Error running CLI: "+cliPath)

			bytes, err := ioutil.ReadFile(fixturePath)
			Ω(err).ShouldNot(HaveOccurred())
			in.Write(bytes)
			in.Close()

			Eventually(session).Should((Exit(0)))
			Ω(session).Should(Say(sorted))
		})
	})
})

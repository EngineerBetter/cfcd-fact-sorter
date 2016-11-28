package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/cf-guardian/guardian/kernel/fileutils"

	"io/ioutil"
	"os/exec"
	"path/filepath"
)

var cliPath string
var fixturePath string

var _ = Describe("cfcd-fact-sorter", func() {
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
})

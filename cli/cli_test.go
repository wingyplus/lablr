package cli_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/lablr/cli"
	"io/ioutil"
	"os"
	"testing"
)

func TestXMLSpecs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CLI Suite")
}

func File(path string) ([]byte, error) {
	return ioutil.ReadFile(os.Getenv("GOPATH") + "/src/github.com/wingyplus/lablr/" + path)
}

var _ = Describe("CLI", func() {
	It("should be return Share search config when input --share-search-config option and xml", func() {
		var modelContent, _ = File("tests/iss_2_Model.xml")
		var c cli.LablrCLI = cli.LablrCLI{
			Options: []string{"--share-search-config"},
			Content: modelContent,
		}

		var result, err = c.Process()
		var expectContent, _ = File("tests/iss_2.xml")

		Expect(err).To(BeNil())
		Expect(string(result)).To(Equal(string(expectContent)))
	})
})

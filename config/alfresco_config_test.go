package config_test

import (
	"encoding/xml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/lablr/config"
)

var _ = Describe("Alfresco Config", func() {
	It("have configs", func() {
		var alfrescoConfig *config.AlfrescoConfig = new(config.AlfrescoConfig)
		var cnf config.Config = config.Config{Evaluator: "string-compare", Condition: "DocumentLibrary"}

		alfrescoConfig.AddConfig(cnf)

		Expect(len(alfrescoConfig.Configs)).To(Equal(1))

		alfrescoConfig.AddConfig(cnf)

		Expect(len(alfrescoConfig.Configs)).To(Equal(2))
	})

	Context("marshal XML", func() {
		var alfrescoConfig *config.AlfrescoConfig
		BeforeEach(func() {
			alfrescoConfig = new(config.AlfrescoConfig)
		})
		It("should parse to xml", func() {
			expectXML := `<alfresco-config></alfresco-config>`
			actualXML, _ := xml.Marshal(alfrescoConfig)

			Expect(string(actualXML)).To(Equal(expectXML))
		})
		It("should has config in alfresco-config", func() {
			alfrescoConfig.AddConfig(config.Config{Evaluator: "string-compare", Condition: "DocumentLibrary"})

			expectXML := `<alfresco-config><config evaluator="string-compare" condition="DocumentLibrary"></config></alfresco-config>`
			actualXML, _ := xml.Marshal(alfrescoConfig)

			Expect(string(actualXML)).To(Equal(expectXML))
		})
	})
})

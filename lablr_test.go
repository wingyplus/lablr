package lablr_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/lablr"
	"testing"
)

func TestLablrSpecs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lablr Suite")
}

var _ = Describe("Lablr", func() {
	Describe("Model unmarshaller", func() {
		var model lablr.Model
		var err error
		var testModel string = `
		<?xml version="1.0" encoding="UTF-8"?>
		<model name="test:testModel" xmlns="http://www.alfresco.org/model/dictionary/1.0">
			<description>Test Model</description>
			<author>test@mail.com</author>
			<version>1.0</version>

			<imports>
				<import uri="http://www.alfresco.org/model/dictionary/1.0" prefix="d" />
				<import uri="http://www.alfresco.org/model/content/1.0" prefix="cm" />
			</imports>

			<namespaces>
				<namespace uri="http://www.alfresco.org/model/test/1.0" prefix="test" />
			</namespaces>

			<types>
				<type name="test:typeName">
					<title>Test Name</title>
					<parent>cm:content</parent>
					<properties>
						<property name="test:propertyName">
							<type>d:text</type>
						</property>
					</properties>
				</type>
			</types>
		</model>
		`
		BeforeEach(func() {
			model, err = lablr.NewModel(testModel)
			Expect(err).To(BeNil())
		})
		Context("unmarshall from xml model string", func() {
			It("should be return Model object", func() {
				Expect(err).To(BeNil())
				Expect(model.Name).To(Equal("test:testModel"))
				Expect(model.Description).To(Equal("Test Model"))
				Expect(model.Version).To(Equal("1.0"))
				Expect(model.Author).To(Equal("test@mail.com"))
			})
		})

		Describe("Type", func() {
			Context("unmarshall from xml model string", func() {
				It("should be return Type from Model object after parsing", func() {
					Expect(len(model.Types)).To(Equal(1))

					var modelType lablr.Type = model.Types[0]
					Expect(modelType.Name).To(Equal("test:typeName"))
					Expect(modelType.Title).To(Equal("Test Name"))
				})
			})
		})

		Describe("Property", func() {
			Context("unmarshall from xml model string", func() {
				It("should be return Property from Type object after parsing", func() {
					var properties []lablr.Property = model.Types[0].Properties
					Expect(len(properties)).To(Equal(1))

					Expect(properties[0].Name).To(Equal("test:propertyName"))
				})
			})
		})
	})
})

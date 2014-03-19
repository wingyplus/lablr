package lablr_test

import (
	"encoding/xml"
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
		var (
			model     lablr.Model
			err       error
			testModel []byte = []byte(`
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
							<property name="test:propertyName_2">
								<type>d:text</type>
							</property>
						</properties>
					</type>
				</types>
			</model>
			`)
		)
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
			Context("convert name", func() {
				It("should be convert from test:testModel to test_testModel", func() {
					Expect(model.ModelName()).To(Equal("test_testModel"))
				})
			})
		})

		Describe("Property", func() {
			Context("unmarshall from xml model string", func() {
				It("should be return Property from Type object after parsing", func() {
					var properties []lablr.Property = model.Types[0].Properties
					Expect(len(properties)).To(Equal(2))

					Expect(properties[0].Name).To(Equal("test:propertyName"))
				})
			})
			Context("convert name", func() {
				It("should be convert from test:propertyName to test_propertyName", func() {
					var properties []lablr.Property = model.Types[0].Properties

					Expect(properties[0].PropertyName()).To(Equal("test_propertyName"))
					Expect(properties[1].PropertyName()).To(Equal("test_propertyName_2"))
				})
			})
		})
	})
	Describe("Share config custom generator", func() {
		Describe("Model config", func() {
			Describe("FieldVisibility", func() {
				Describe("Field", func() {
					It("should be generate from Property object", func() {
						var property lablr.Property = lablr.Property{Name: "test:testProperty"}
						var field lablr.Field = lablr.NewField(property)
						Expect(field.Id).To(Equal("test:testProperty"))

						property = lablr.Property{Name: "test:fieldVisibility"}
						field = lablr.NewField(property)
						Expect(field.Id).To(Equal("test:fieldVisibility"))
					})
					Context("marshall XML", func() {
						It("should be marshall to expect XML", func() {
							expectXML := `<show id="test:testProperty"></show>`
							actualXML, _ := xml.Marshal(&lablr.Field{Id: "test:testProperty"})

							Expect(string(actualXML)).To(Equal(expectXML))

							expectXML = `<show id="test:testProperty2"></show>`
							actualXML, _ = xml.Marshal(lablr.Field{Id: "test:testProperty2"})

							Expect(string(actualXML)).To(Equal(expectXML))
						})
					})
				})
				It("should be return List of Field", func() {
					var fieldVisibility *lablr.FieldVisibility = new(lablr.FieldVisibility)

					fieldVisibility.Add(lablr.Field{Id: "test:testProperty"})
					Expect(fieldVisibility.Fields[0].Id).To(Equal("test:testProperty"))

					fieldVisibility.Add(lablr.Field{Id: "test:testProperty"})
					Expect(fieldVisibility.Fields[1].Id).To(Equal("test:testProperty"))
				})
				Context("marshall to xml", func() {
					It("should be return expect XML", func() {
						var fieldVisibility *lablr.FieldVisibility = new(lablr.FieldVisibility)
						fieldVisibility.Add(lablr.Field{Id: "test:testProperty"})

						expectXML := `<field-visibility><show id="test:testProperty"></show></field-visibility>`
						actualXML, _ := xml.Marshal(fieldVisibility)

						Expect(string(actualXML)).To(Equal(expectXML))

						fieldVisibility.Add(lablr.Field{Id: "test:testProperty"})
						expectXML = `<field-visibility><show id="test:testProperty"></show><show id="test:testProperty"></show></field-visibility>`
						actualXML, _ = xml.Marshal(fieldVisibility)

						Expect(string(actualXML)).To(Equal(expectXML))
					})
				})
			})
			Describe("Form", func() {
				It("should has id", func() {
					var form lablr.Form = lablr.Form{Id: "search"}
					Expect(form.Id).To(Equal("search"))

					form = lablr.Form{Id: "doclib-inline-edit"}
					Expect(form.Id).To(Equal("doclib-inline-edit"))
				})
				It("should has field-visibility", func() {
					var fieldVisibility *lablr.FieldVisibility = new(lablr.FieldVisibility)
					fieldVisibility.Add(lablr.Field{Id: "test:testProperty"})

					var form lablr.Form = lablr.Form{Id: "search", FieldVisibility: fieldVisibility}
					Expect(form.FieldVisibility.Fields[0].Id).To(Equal("test:testProperty"))
				})
				Context("marshall to XML", func() {
					It("should marshall to expect XML", func() {
						var fieldVisibility *lablr.FieldVisibility = new(lablr.FieldVisibility)
						fieldVisibility.Add(lablr.Field{Id: "test:testProperty"})

						var form lablr.Form = lablr.Form{Id: "search", FieldVisibility: fieldVisibility}

						expectXML := `<form id="search"><field-visibility><show id="test:testProperty"></show></field-visibility></form>`
						actualXML, _ := xml.Marshal(&form)

						Expect(string(actualXML)).To(Equal(expectXML))
					})
				})
			})
			Describe("Config", func() {
				It("should has forms", func() {
					var config *lablr.Config = &lablr.Config{}
					config.AddForm(lablr.Form{Id: "search"})

					Expect(len(config.Forms)).To(Equal(1))
				})
				It("should be evaluator", func() {
					var config *lablr.Config = &lablr.Config{Evaluator: "string-compare"}
					Expect(config.Evaluator).To(Equal("string-compare"))
				})
				It("should has condition", func() {
					var config *lablr.Config = &lablr.Config{Condition: "test:testType"}
					Expect(config.Condition).To(Equal("test:testType"))
				})
				Context("marshal to XML", func() {
					It("should be return expect XML", func() {
						var fieldVisibility *lablr.FieldVisibility = new(lablr.FieldVisibility)
						fieldVisibility.Add(lablr.Field{Id: "test:testProperty"})
						var form lablr.Form = lablr.Form{Id: "search", FieldVisibility: fieldVisibility}
						var config *lablr.Config = &lablr.Config{Evaluator: "string-compare", Condition: "test:testType"}
						config.AddForm(form)

						expectXML := `<config evaluator="string-compare" condition="test:testType"><forms><form id="search"><field-visibility><show id="test:testProperty"></show></field-visibility></form></forms></config>`
						actualXML, _ := xml.Marshal(config)

						Expect(string(actualXML)).To(Equal(expectXML))
					})
				})
			})
		})
	})
	Describe("Search generator", func() {
		It("should return search config", func() {
			var config *lablr.Config = lablr.NewSearchConfig(lablr.Type{Name: "test:testType"})
			Expect(config.Evaluator).To(Equal("model-type"))
			Expect(config.Condition).To(Equal("test:testType"))

			config = lablr.NewSearchConfig(lablr.Type{Name: "test:testType2"})
			Expect(config.Condition).To(Equal("test:testType2"))

			properties := []lablr.Property{
				lablr.Property{Name: "test:testProperty"},
			}
			t := lablr.Type{Name: "test:testType", Properties: properties}
			config = lablr.NewSearchConfig(t)
			Expect(len(config.Forms)).To(Equal(1))
			Expect(len(config.Forms[0].FieldVisibility.Fields)).To(Equal(1))
			Expect(config.Forms[0].FieldVisibility.Fields[0].Id).To(Equal("test:testProperty"))

			properties = []lablr.Property{
				lablr.Property{Name: "test:testProperty"},
				lablr.Property{Name: "test:testProperty2"},
			}
			t = lablr.Type{Name: "test:testType", Properties: properties}
			config = lablr.NewSearchConfig(t)
			Expect(len(config.Forms[0].FieldVisibility.Fields)).To(Equal(2))
			Expect(config.Forms[0].FieldVisibility.Fields[1].Id).To(Equal("test:testProperty2"))
		})
	})
})

package config_test

import (
	"encoding/xml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/lablr"
	"github.com/wingyplus/lablr/config"
	"testing"
)

func TestShareConfigSpecs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Share config Suite")
}

var _ = Describe("Share config", func() {
	Describe("FieldVisibility", func() {
		Describe("Field", func() {
			It("should be generate from Property object", func() {
				var property lablr.Property = lablr.Property{Name: "test:testProperty"}
				var field config.Field = config.NewField(property)
				Expect(field.Id).To(Equal("test:testProperty"))

				property = lablr.Property{Name: "test:fieldVisibility"}
				field = config.NewField(property)
				Expect(field.Id).To(Equal("test:fieldVisibility"))
			})
			Context("marshall XML", func() {
				It("should be marshall to expect XML", func() {
					expectXML := `<show id="test:testProperty"></show>`
					actualXML, _ := xml.Marshal(&config.Field{Id: "test:testProperty"})

					Expect(string(actualXML)).To(Equal(expectXML))

					expectXML = `<show id="test:testProperty2"></show>`
					actualXML, _ = xml.Marshal(config.Field{Id: "test:testProperty2"})

					Expect(string(actualXML)).To(Equal(expectXML))
				})
			})
		})
		It("should be return List of Field", func() {
			var fieldVisibility *config.FieldVisibility = new(config.FieldVisibility)

			fieldVisibility.Add(config.Field{Id: "test:testProperty"})
			Expect(fieldVisibility.Fields[0].Id).To(Equal("test:testProperty"))

			fieldVisibility.Add(config.Field{Id: "test:testProperty"})
			Expect(fieldVisibility.Fields[1].Id).To(Equal("test:testProperty"))
		})
		Context("marshall to xml", func() {
			It("should be return expect XML", func() {
				var fieldVisibility *config.FieldVisibility = new(config.FieldVisibility)
				fieldVisibility.Add(config.Field{Id: "test:testProperty"})

				expectXML := `<field-visibility><show id="test:testProperty"></show></field-visibility>`
				actualXML, _ := xml.Marshal(fieldVisibility)

				Expect(string(actualXML)).To(Equal(expectXML))

				fieldVisibility.Add(config.Field{Id: "test:testProperty"})
				expectXML = `<field-visibility><show id="test:testProperty"></show><show id="test:testProperty"></show></field-visibility>`
				actualXML, _ = xml.Marshal(fieldVisibility)

				Expect(string(actualXML)).To(Equal(expectXML))
			})
		})
	})
	Describe("Form", func() {
		It("should has id", func() {
			var form config.Form = config.Form{Id: "search"}
			Expect(form.Id).To(Equal("search"))

			form = config.Form{Id: "doclib-inline-edit"}
			Expect(form.Id).To(Equal("doclib-inline-edit"))
		})
		It("should has field-visibility", func() {
			var fieldVisibility *config.FieldVisibility = new(config.FieldVisibility)
			fieldVisibility.Add(config.Field{Id: "test:testProperty"})

			var form config.Form = config.Form{Id: "search", FieldVisibility: fieldVisibility}
			Expect(form.FieldVisibility.Fields[0].Id).To(Equal("test:testProperty"))
		})
		Context("marshall to XML", func() {
			It("should marshall to expect XML", func() {
				var fieldVisibility *config.FieldVisibility = new(config.FieldVisibility)
				fieldVisibility.Add(config.Field{Id: "test:testProperty"})

				var form config.Form = config.Form{Id: "search", FieldVisibility: fieldVisibility}

				expectXML := `<form id="search"><field-visibility><show id="test:testProperty"></show></field-visibility></form>`
				actualXML, _ := xml.Marshal(&form)

				Expect(string(actualXML)).To(Equal(expectXML))
			})
		})
	})
	Describe("Config", func() {
		It("should has forms", func() {
			var cnf *config.Config = config.NewConfig("", "")
			cnf.AddForm(config.Form{Id: "search"})

			Expect(len(cnf.Forms.Forms)).To(Equal(1))
		})
		It("should be evaluator", func() {
			var cnf *config.Config = &config.Config{Evaluator: "string-compare"}
			Expect(cnf.Evaluator).To(Equal("string-compare"))
		})
		It("should has condition", func() {
			var cnf *config.Config = &config.Config{Condition: "test:testType"}
			Expect(cnf.Condition).To(Equal("test:testType"))
		})
		Context("marshal to XML", func() {
			It("should be return expect XML", func() {
				var fieldVisibility *config.FieldVisibility = new(config.FieldVisibility)
				fieldVisibility.Add(config.Field{Id: "test:testProperty"})
				var form config.Form = config.Form{Id: "search", FieldVisibility: fieldVisibility}
				var cnf *config.Config = config.NewConfig("string-compare", "test:testType")
				cnf.AddForm(form)

				expectXML := `<config evaluator="string-compare" condition="test:testType"><forms><form id="search"><field-visibility><show id="test:testProperty"></show></field-visibility></form></forms></config>`
				actualXML, _ := xml.Marshal(cnf)

				Expect(string(actualXML)).To(Equal(expectXML))
			})
			It("should be omited when has no value in Form", func() {
				var cnf *config.Config = &config.Config{Evaluator: "string-compare", Condition: "DocumentLibrary"}

				expectXML := `<config evaluator="string-compare" condition="DocumentLibrary"></config>`
				actualXML, _ := xml.Marshal(cnf)

				Expect(string(actualXML)).To(Equal(expectXML))
			})
		})
	})
	Describe("Forms", func() {
		It("have forms", func() {
			var forms *config.Forms = new(config.Forms)
			forms.AddForm(config.Form{})

			Expect(len(forms.Forms)).To(Equal(1))
		})
		It("should marshal to xml", func() {
			var forms *config.Forms = new(config.Forms)

			expectXML := `<forms></forms>`
			actualXML, _ := xml.Marshal(forms)

			Expect(string(actualXML)).To(Equal(expectXML))
		})
	})
})

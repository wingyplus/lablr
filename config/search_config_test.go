package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/lablr"
	"github.com/wingyplus/lablr/config"
)

var _ = Describe("Share search config", func() {
	Describe("Search config", func() {
		It("should return evaluator and condition for search config", func() {
			var cnf *config.Config = config.NewSearchConfig(lablr.Type{Name: "test:testType"})
			Expect(cnf.Evaluator).To(Equal("model-type"))
			Expect(cnf.Condition).To(Equal("test:testType"))

			cnf = config.NewSearchConfig(lablr.Type{Name: "test:testType2"})
			Expect(cnf.Condition).To(Equal("test:testType2"))
		})
		It("should return expect field-visibility by given 1 properties", func() {
			properties := []lablr.Property{
				lablr.Property{Name: "test:testProperty"},
			}
			t := lablr.Type{Name: "test:testType", Properties: properties}
			cnf := config.NewSearchConfig(t)
			Expect(len(cnf.Forms.Forms)).To(Equal(1))
			Expect(len(cnf.Forms.Forms[0].FieldVisibility.Fields)).To(Equal(1))
			Expect(cnf.Forms.Forms[0].FieldVisibility.Fields[0].Id).To(Equal("test:testProperty"))
		})
		It("should return expect field-visibility by given 2 properties", func() {
			properties := []lablr.Property{
				lablr.Property{Name: "test:testProperty"},
				lablr.Property{Name: "test:testProperty2"},
			}
			t := lablr.Type{Name: "test:testType", Properties: properties}
			cnf := config.NewSearchConfig(t)
			Expect(len(cnf.Forms.Forms[0].FieldVisibility.Fields)).To(Equal(2))
			Expect(cnf.Forms.Forms[0].FieldVisibility.Fields[1].Id).To(Equal("test:testProperty2"))
		})
	})
})

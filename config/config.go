package config

import (
	"encoding/xml"
	"github.com/wingyplus/lablr"
)

type Field struct {
	XMLName xml.Name `xml:"show,omitempty"`
	Id      string   `xml:"id,attr"`
}

func NewField(property lablr.Property) Field {
	return Field{Id: property.Name}
}

type FieldVisibility struct {
	XMLName xml.Name `xml:"field-visibility"`
	Fields  []Field
}

func (fieldVisibility *FieldVisibility) Add(f Field) {
	fieldVisibility.Fields = append(fieldVisibility.Fields, f)
}

type Form struct {
	XMLName         xml.Name `xml:"form"`
	Id              string   `xml:"id,attr"`
	FieldVisibility *FieldVisibility
}

type Forms struct {
	XMLName xml.Name `xml:"forms,omitempty"`
	Forms   []Form
}

func (fs *Forms) AddForm(f Form) {
	fs.Forms = append(fs.Forms, f)
}

type Config struct {
	XMLName   xml.Name `xml:"config"`
	Evaluator string   `xml:"evaluator,attr"`
	Condition string   `xml:"condition,attr"`
	Forms     *Forms
}

func NewConfig(evaluator, condition string) (cnf *Config) {
	cnf = &Config{Evaluator: evaluator, Condition: condition}
	cnf.Forms = new(Forms)

	return
}

func (c *Config) AddForm(f Form) {
	c.Forms.AddForm(f)
}

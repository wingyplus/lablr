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

type Config struct {
	XMLName   xml.Name `xml:"config"`
	Evaluator string   `xml:"evaluator,attr"`
	Condition string   `xml:"condition,attr"`
	Forms     []Form   `xml:"forms>form"`
}

func (c *Config) AddForm(f Form) {
	c.Forms = append(c.Forms, f)
}

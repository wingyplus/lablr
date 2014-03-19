package lablr

import (
	"encoding/xml"
	"strings"
)

type Model struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description"`
	Version     string `xml:"version"`
	Author      string `xml:"author"`
	Types       []Type `xml:"types>type"`
}

type Type struct {
	Name       string     `xml:"name,attr"`
	Title      string     `xml:"title"`
	Properties []Property `xml:"properties>property"`
}

type Property struct {
	Name string `xml:"name,attr"`
}

func (p Property) PropertyName() string {
	return strings.Replace(p.Name, ":", "_", 1)
}

func NewModel(modelContent []byte) (model Model, err error) {
	model = Model{}
	err = xml.Unmarshal(modelContent, &model)

	return
}

func (m Model) ModelName() string {
	return strings.Replace(m.Name, ":", "_", 1)
}

type Field struct {
	XMLName xml.Name `xml:"show,omitempty"`
	Id      string   `xml:"id,attr"`
}

func NewField(property Property) Field {
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

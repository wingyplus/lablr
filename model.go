package lablr

import "encoding/xml"

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

func NewModel(modelContent []byte) (model Model, err error) {
	model = Model{}
	err = xml.Unmarshal(modelContent, &model)

	return
}

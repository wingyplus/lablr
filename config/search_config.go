package config

import "github.com/wingyplus/lablr"

func NewSearchConfig(t lablr.Type) (config *Config) {
	fieldVisibility := new(FieldVisibility)
	if len(t.Properties) != 0 {
		for _, property := range t.Properties {
			fieldVisibility.Add(NewField(property))
		}
	}

	form := Form{Id: "search", FieldVisibility: fieldVisibility}

	config = NewConfig("model-type", t.Name)
	config.AddForm(form)

	return config
}

package config

import "encoding/xml"

type AlfrescoConfig struct {
	XMLName xml.Name `xml:"alfresco-config"`
	Configs []*Config
}

func (alfConf *AlfrescoConfig) AddConfig(c *Config) {
	alfConf.Configs = append(alfConf.Configs, c)
}

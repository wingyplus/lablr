package cli

import (
	"encoding/xml"
	"github.com/wingyplus/lablr"
	"github.com/wingyplus/lablr/config"
)

type LablrCLI struct {
	Options []string
	Content []byte
}

func (lc LablrCLI) Process() ([]byte, error) {
	m, err := lablr.NewModel(lc.Content)
	if err != nil {
		return nil, err
	}

	alfConfig := new(config.AlfrescoConfig)
	for _, t := range m.Types {
		alfConfig.AddConfig(config.NewSearchConfig(t))
	}

	return xml.MarshalIndent(alfConfig, "", "    ")
}

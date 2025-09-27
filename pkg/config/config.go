package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/UserL4mbda/netbox-importer/pkg/netbox"
)

// LoadFromFile charge la configuration depuis un fichier YAML
func LoadFromFile(filename string) (*netbox.Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config netbox.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

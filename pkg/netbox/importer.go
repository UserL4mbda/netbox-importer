package netbox

import (
	"fmt"
	"log"
)

// Importer gère l'import de configuration vers NetBox
type Importer struct {
	client *Client
}

// NewImporter crée un nouvel importeur
func NewImporter(client *Client) *Importer {
	return &Importer{
		client: client,
	}
}

// ImportConfig importe une configuration YAML dans NetBox
func (i *Importer) ImportConfig(config *Config) error {
	log.Println("Début de l'import NetBox...")

	// Implémentation de la logique d'import...
	// (utiliser les méthodes du client)

	return i.importSites(config)
}

func (i *Importer) importSites(config *Config) error {
	// Créer le site
	var site Site
	err := i.client.Create("dcim/sites", config.Site, &site)
	if err != nil {
		return fmt.Errorf("erreur création site: %v", err)
	}

	log.Printf("Site créé: %s (ID: %d)", site.Name, site.ID)
	return nil
}

// ImportFromFile importe depuis un fichier YAML
func (i *Importer) ImportFromFile(filename string) error {
	config, err := LoadConfig(filename)
	if err != nil {
		return err
	}

	return i.ImportConfig(config)
}

// LoadConfig charge la configuration depuis un fichier YAML
func LoadConfig(filename string) (*Config, error) {
	// Implémentation de la lecture YAML...
	return &Config{}, nil
}

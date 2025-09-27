package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/UserL4mbda/netbox-importer/pkg/netbox"
)

func main() {
	var (
		netboxURL   = flag.String("url", "", "URL de NetBox")
		netboxToken = flag.String("token", "", "Token API NetBox")
		configFile  = flag.String("config", "", "Fichier de configuration YAML")
	)
	flag.Parse()

	if *netboxURL == "" || *netboxToken == "" || *configFile == "" {
		fmt.Println("Usage: netbox-importer -url <url> -token <token> -config <file>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Créer le client NetBox
	client := netbox.NewClient(*netboxURL, *netboxToken)
	importer := netbox.NewImporter(client)

	// Importer la configuration
	err := importer.ImportFromFile(*configFile)
	if err != nil {
		log.Fatalf("Erreur lors de l'import: %v", err)
	}

	log.Println("Import terminé avec succès!")
}

# NetBox Importer

Bibliothèque Go pour importer des configurations YAML dans NetBox via l'API.

## Installation

```bash
go get github.com/UserL4mbda/netbox-importer
```

## Utilisation

### Comme bibliothèque

```go
import "github.com/UserL4mbda/netbox-importer/pkg/netbox"

func main() {
    client := netbox.NewClient("https://netbox.example.com", "votre-token")
    importer := netbox.NewImporter(client)
    
    err := importer.ImportFromFile("config.yaml")
    if err != nil {
        log.Fatal(err)
    }
}
```

### CLI
`cmd/netbox-importer/main.go`

```bash
netbox-importer \
    -url https://netbox.example.com \
    -token votre-token-api \
    -config config.yaml
```

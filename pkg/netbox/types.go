package netbox

// Structures pour l'API NetBox
type Site struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
}

type Location struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Display     string `json:"display,omitempty"`
	SiteID      int    `json:"site"`
	TypeID      int    `json:"location_type,omitempty"`
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
}

type Device struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Display     string `json:"display,omitempty"`
	DeviceType  int    `json:"device_type"`
	Role        int    `json:"role"`
	SiteID      int    `json:"site"`
	LocationID  int    `json:"location,omitempty"`
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
}

type Cable struct {
	Type          string             `json:"type"`
	Status        string             `json:"status"`
	Description   string             `json:"description,omitempty"`
	ATerminations []CableTermination `json:"a_terminations"`
	BTerminations []CableTermination `json:"b_terminations"`
}

type CableTermination struct {
	ObjectType string `json:"object_type"`
	ObjectID   int    `json:"object_id"`
}

// Structures pour la configuration YAML
type Config struct {
	Site          SiteConfig           `yaml:"site"`
	LocationTypes []LocationTypeConfig `yaml:"location_types"`
	Locations     []LocationConfig     `yaml:"locations"`
	DeviceTypes   []DeviceTypeConfig   `yaml:"device_types"`
	Devices       []DeviceConfig       `yaml:"devices"`
	Cables        []CableConfig        `yaml:"cables"`
}

type SiteConfig struct {
	Name        string `yaml:"name"`
	Slug        string `yaml:"slug"`
	Status      string `yaml:"status"`
	Description string `yaml:"description"`
}

// LocationTypeConfig correspond aux types de lieux (ex: "Salle serveur", "Bureau")
type LocationTypeConfig struct {
	Name string `yaml:"name"`
	Slug string `yaml:"slug"`
}

// LocationConfig correspond à chaque élément de la liste "locations"
type LocationConfig struct {
	Name        string `yaml:"name"`
	Slug        string `yaml:"slug"`
	Display     string `yaml:"display"`
	Site        string `yaml:"site"` // Fait référence au nom du site
	Status      string `yaml:"status"`
	Description string `yaml:"description"`
}

// DeviceTypeConfig correspond à chaque élément de la liste "device_types"
type DeviceTypeConfig struct {
	Model        string `yaml:"model"`
	Slug         string `yaml:"slug"`
	Manufacturer string `yaml:"manufacturer"` // Fait référence au nom du fabricant
}

// DeviceConfig correspond à chaque élément de la liste "devices"
type DeviceConfig struct {
	Name        string `yaml:"name"`
	Display     string `yaml:"display"`
	DeviceType  string `yaml:"device_type"` // Fait référence au modèle du type d'appareil
	Site        string `yaml:"site"`
	Location    string `yaml:"location"` // Fait référence au nom du lieu
	Status      string `yaml:"status"`
	Description string `yaml:"description"`
}

type CableConfig struct {
	FromDevice  string `yaml:"from_device"`
	FromPort    string `yaml:"from_port"`
	ToDevice    string `yaml:"to_device"`
	ToPort      string `yaml:"to_port"`
	Type        string `yaml:"type"`
	Status      string `yaml:"status"`
	Description string `yaml:"description"`
}

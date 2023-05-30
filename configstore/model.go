package configstore

// swagger:model Config
type Config struct {
	// Id of the config
	// in: string
	Id string `json:"Id"`

	// Map of entries of the config
	// in: map[string]string
	Entries map[string]string `json:"entries"`

	//Labels of the config
	//in: string
	Labels string `json:"labels"`

	//Version of the config
	//in: string
	Version string `json:"version"`
}

type Service struct {
	Data map[string]*[]Config
}

// swagger:model CfGroup
type CfGroup struct {
	// Id of the cfgroup
	// in: string
	Id string `json:"Id"`

	// Configurations of the cfgroup
	// in: []*Config
	Configurations []*Config `json:"Configurations"`

	// Version of the cfgroup
	// in: string
	Version string `json:"version"`
}

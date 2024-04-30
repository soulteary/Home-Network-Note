package devices

type Device struct {
	Name    string `yaml:"Name"`
	TurnOn  string `yaml:"TurnOn"`
	TurnOff string `yaml:"TurnOff"`
	TLDR    string `yaml:"TLDR,omitempty"`
	Info    struct {
		Syetem      string `yaml:"Syetem,omitempty"`
		CPU         string `yaml:"CPU,omitempty"`
		GPU         string `yaml:"GPU,omitempty"`
		RAM         string `yaml:"RAM,omitempty"`
		Disk        string `yaml:"Disk,omitempty"`
		LAN         string `yaml:"LAN,omitempty"`
		WiFi        string `yaml:"WiFi,omitempty"`
		Thunderbolt string `yaml:"Thunderbolt,omitempty"`
	} `yaml:"Info"`
	Description string
}

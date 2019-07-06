package conf

type Config struct {
	Db struct {
		User     string `yaml:"user"`
		Passwd   string `yaml:"passwd"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	}

	Http struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
	}
}

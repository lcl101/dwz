package conf

var Conf Config

type Config struct {
	Db struct {
		User     string `yaml:"user"`
		Passwd   string `yaml:"passwd"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
	}

	Http struct {
		IP   string `yaml:"ip"`
		Port int    `yaml:"port"`
	}

	Url struct {
		Home string `yaml:"home"`
	}

	Gen struct {
		Base      string `yaml:"base"`
		Length    int    `yaml:"length"`
		Unique    bool   `yaml:"unique"`
		Humanity  bool   `yaml:"humanity"`
		Algorithm int    `yaml:"algorithm"`
	}
}

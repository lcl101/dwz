package conf

var Conf Config

type Config struct {
	LogLevel string `yaml:"loglevel"`
	Db       DB     `yaml:"Db"`
	Http     HTTP   `yaml:"Http"`
	Url      URL    `yaml:"Url"`
	Gen      GEN    `yaml:"Gen"`
}

type DB struct {
	User     string `yaml:"user"`
	Passwd   string `yaml:"passwd"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type HTTP struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type URL struct {
	Home string `yaml:"home"`
}

type GEN struct {
	Base      string `yaml:"base"`
	Length    int    `yaml:"length"`
	Unique    bool   `yaml:"unique"`
	Humanity  bool   `yaml:"humanity"`
	Algorithm int    `yaml:"algorithm"`
}

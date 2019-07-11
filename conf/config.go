package conf

var Conf Config

type Config struct {
	LogLevel string `yaml:"loglevel"`
	Db       DB     `yaml:"db"`
	Http     HTTP   `yaml:"http"`
	Url      URL    `yaml:"url"`
	Gen      GEN    `yaml:"gen"`
	App      APP    `yaml:"app"`
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

type APP struct {
	AppKey    string `yaml:"appKey"`
	AppSecret string `yaml:"appSecret"`
}

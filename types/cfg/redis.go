package cfg

type Redis struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Pass   string `yaml:"pass"`
	Prefix string `yaml:"prefix"`
}

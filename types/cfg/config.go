package cfg

type Config struct {
	Db         Database `yaml:"db"`
	Http       Http     `yaml:"http"`
	Redis      Redis    `yaml:"redis"`
	SessionKey string   `yaml:"session_key"`
}

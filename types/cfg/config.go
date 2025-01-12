package cfg

type Config struct {
	Db         Database    `yaml:"db"`
	Http       Http        `yaml:"http"`
	Redis      Redis       `yaml:"redis"`
	SessionKey string      `yaml:"session_key"`
	AdminIds   []int64     `yaml:"admin_ids"`
	Migrations []Migration `yaml:"migrations"`
}

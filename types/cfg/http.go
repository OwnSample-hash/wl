package cfg

type Http struct {
	Bind         string `yaml:"bind"`
	Port         int    `yaml:"port"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	IdealTimeout int    `yaml:"ideal_timeout"`
}

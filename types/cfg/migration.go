package cfg

type Migration struct {
	Path       string `yaml:"path"`
	Hash       string `yaml:"hash"`
	DBName     string `yaml:"db_name"`
	ObjectType string `yaml:"object_type"`
}

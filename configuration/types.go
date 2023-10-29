package configuration

type App struct {
	Port   string `yaml:"port"`
	Logger bool   `yaml:"logger"`
}

type Database struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}

type Configuration struct {
	DB  Database `yaml:"Database"`
	App App      `yaml:"App"`
	Jwt Jwt      `yaml:"Jwt"`
}

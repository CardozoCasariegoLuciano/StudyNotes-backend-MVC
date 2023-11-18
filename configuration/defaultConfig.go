package configuration

func setDefaultConfig(config *Configuration) *Configuration {
	app := App{
		Port:   ":5999",
		Logger: false,
	}

	jwt := Jwt{
		Secret: "publico",
	}

	database := Database{
		Port:     3607,
		Host:     "localhost",
		Password: "123123123",
		User:     "root",
		Name:     "Table",
	}

	config.App = app
	config.DB = database
	config.Jwt = jwt

	return config
}

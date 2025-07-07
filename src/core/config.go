package config

type Config struct {
	CDN struct {
		Url string
	}
}

type Map struct {
	Config
}

func DefaultConfig() Map {
	c := Map{
		Config: Config{
			CDN: struct {
				Url string
			}{
				Url: "https://cdn.rafailmdzdv.ru/",
			},
		},
	}
	return c
}

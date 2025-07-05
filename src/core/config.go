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
				Url: "http://91.108.245.148:8000/",
			},
		},
	}
	return c
}

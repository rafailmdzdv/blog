package config

type Config struct {
	CDN struct {
		Url string
	}

	API struct {
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
				Url: "https://cdn.rafailmdzdv.ru/ico/",
			},
			API: struct {
				Url string
			}{
				Url: "https://app.blog.rafailmdzdv.ru",
			},
		},
	}
	return c
}

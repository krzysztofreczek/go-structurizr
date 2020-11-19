package scraper

type Configuration struct {
	packages []string
}

func NewConfiguration(packages ...string) Configuration {
	return Configuration{
		packages: packages,
	}
}

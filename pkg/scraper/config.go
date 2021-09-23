package scraper

// Configuration is an open structure that contains scraper configuration.
//
// Packages contain prefixes of packages for scraper to scrape.
// Each object of package that does not match any of predefined
// package prefixes is omitted and its internal structure is not scraped.
// When no package prefix is provided, the scraper will stop
// scraping given structure on a root level.
type Configuration struct {
	Packages []string
}

// NewConfiguration instantiates Configuration with a set of package
// prefixes provided with variadic argument.
func NewConfiguration(packages ...string) Configuration {
	return Configuration{
		Packages: packages,
	}
}

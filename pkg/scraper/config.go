package scraper

// Configuration is an open structure that holds the scraper configuration.
//
// Packages contain prefixes of packages for the scraper to process.
// Any package object that does not match the provided prefixes will be omitted,
// and its internal structure will not be scraped.
// If no package prefixes are provided, the scraper will only process the root level of the structure.
type Configuration struct {
	Packages []string
}

// NewConfiguration creates a Configuration with the specified package prefixes.
//
// It takes a variadic argument to accept multiple package prefixes.
func NewConfiguration(packages ...string) Configuration {
	return Configuration{
		Packages: packages,
	}
}

package main

import (
	"os"

	"github.com/jaegertracing/jaeger/cmd/collector/app"
	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"
)

func main() {
	s := buildScraper()

	collector := app.New(&app.CollectorParams{})
	svcStructure := s.Scrape(collector)

	f, err := os.Create(".out/output.plantuml")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	v := buildView()
	err = v.RenderStructureTo(svcStructure, f)
	if err != nil {
		panic(err)
	}
}

func buildScraper() scraper.Scraper {
	s, err := scraper.NewScraperFromConfigFile("scraper.yaml")
	if err != nil {
		panic(err)
	}
	return s
}

func buildView() view.View {
	v, err := view.NewViewFromConfigFile("scraper.yaml")
	if err != nil {
		panic(err)
	}
	return v
}

package main

import (
	"github.com/jaegertracing/jaeger/cmd/collector/app"
	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"
	"os"
)

func main() {
	c := scraper.NewConfiguration("")
	s := scraper.NewScraper(c)

	r, err := scraper.NewRule().
		WithApplyFunc(func(name string, groups ...string) model.Info {
			return model.ComponentInfo(name)
		}).
		Build()
	if err != nil {
		panic(err)
	}

	err = s.RegisterRule(r)
	if err != nil {
		panic(err)
	}

	collector := app.New(&app.CollectorParams{})
	svcStructure := s.Scrape(collector)

	v := view.NewView().
		WithTitle("jaeger collector").
		Build()

	f, err := os.Create("output.plantuml")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	err = v.RenderStructureTo(svcStructure, f)
	if err != nil {
		panic(err)
	}
}

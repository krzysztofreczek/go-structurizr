package main

import (
	"image/color"
	"os"

	"github.com/jaegertracing/jaeger/cmd/collector/app"
	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"
)

var (
	appBGColor     = color.RGBA{R: 0x1a, G: 0x45, B: 0x77, A: 0xff}
	appFTColor     = color.White
	handlerBGColor = color.RGBA{R: 0x2d, G: 0x69, B: 0xb7, A: 0xff}
	handlerFTColor = color.White
	auxBGColor     = color.RGBA{R: 0xc8, G: 0xc8, B: 0xc8, A: 0xff}
	auxFTColor     = color.Black
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
	c := scraper.NewConfiguration(
		"github.com/jaegertracing",
		"net/http",
		"go.uber.org/zap",
		"google.golang.org/grpc",
	)
	c.LogDebug = true

	s := scraper.NewScraper(c)

	r := buildScraperRuleForJaegerHandlers()
	if err := s.RegisterRule(r); err != nil {
		panic(err)
	}

	r = buildScraperRuleForJaegerAppComponents()
	if err := s.RegisterRule(r); err != nil {
		panic(err)
	}

	r = buildScraperRuleForAuxiliaryComponents()
	if err := s.RegisterRule(r); err != nil {
		panic(err)
	}

	return s
}

func buildScraperRuleForJaegerHandlers() scraper.Rule {
	r, err := scraper.NewRule().
		WithPkgRegexps(".*/jaeger.*").
		WithNameRegexp(".*Handler.*").
		WithApplyFunc(func(name string, groups ...string) model.Info {
			return model.ComponentInfo(name, "handler component", "", "HANDLER")
		}).
		Build()
	if err != nil {
		panic(err)
	}
	return r
}

func buildScraperRuleForJaegerAppComponents() scraper.Rule {
	r, err := scraper.NewRule().
		WithPkgRegexps(".*/jaeger.*").
		WithApplyFunc(func(name string, groups ...string) model.Info {
			return model.ComponentInfo(name, "app component", "", "APP")
		}).
		Build()
	if err != nil {
		panic(err)
	}
	return r
}

func buildScraperRuleForAuxiliaryComponents() scraper.Rule {
	r, err := scraper.NewRule().
		WithNameRegexp(".*(Server|Logger)$").
		WithApplyFunc(func(name string, groups ...string) model.Info {
			return model.ComponentInfo(name, "auxiliary component", "", "AUX")
		}).
		Build()
	if err != nil {
		panic(err)
	}
	return r
}

func buildView() view.View {
	return view.NewView().
		WithTitle("Components: Jaeger Collector").
		WithRootComponentTag("APP").
		WithComponentStyle(
			view.NewComponentStyle("APP").
				WithBackgroundColor(appBGColor).
				WithFontColor(appFTColor).
				Build()).
		WithComponentStyle(
			view.NewComponentStyle("AUX").
				WithBackgroundColor(auxBGColor).
				WithFontColor(auxFTColor).
				Build()).
		WithComponentStyle(
			view.NewComponentStyle("HANDLER").
				WithBackgroundColor(handlerBGColor).
				WithFontColor(handlerFTColor).
				Build()).
		Build()
}

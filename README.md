# go-structurizr
This library allows you to auto-generate C4 component diagrams out from the Golang code.

## How it works?
The library provides a set of tools that allow you to scrape and render given object into a [C4 component](https://c4model.com/) diagram in [*.plantuml](https://plantuml.com/) format.
Scraper component reflects given structure in accordance with structure interfaces, predefined rules and configuration. You may pass the scraped structure into a view definition with you can then rendered into a plantUML diagram code. 

Scraper identifies components to scrape in one od the following scenarios:
* type that is scraped implements an interface `model.HasInfo`.
* type that is scraped applies to one of the rules registered in the scraper.

## Components

### Component Info

Structure `mode.Info` is a basic structure that defines a component included in the scraped structure of your code.
```go
type Info struct {
	Kind        string      // kind of scraped component
	Description string      // component description
	Technology  string      // technology used within the component
	Tags        []string    // tags are used to match view styles to component
}
```

### Scraper

In order to instantiate the scraper you need to provide scraper configuration which contains a slice of prefixes of packages that you want to reflect. Types that do not match any of the given prefixes will not be traversed. 
```go
config := scraper.NewConfiguration(
    "github.com/karhoo/svc-billing",
)
s := scraper.NewScraper(config)
```

Having a scraper instantiated, you can register a set of rules that will allow the scraper to identify the components you want to include in the output structure.

Each rule consists of:
* Set of package regexp's - only types of package matching at least one package regexp is processed
* Name regexp - only type of name matching regexp is processed
* Apply function - function that produces `model.Info` describing the component included in the scraped structure

```go
r, _ := scraper.NewRule().
    WithPkgRegexps("github.com/krzysztofreczek/pkg/foo/.*").
    WithNameRegexp("^(.*)Client$").
    WithApplyFunc(
        func() model.Info {
            return model.ComponentInfo("foo client", "client of a foo service", "TAG")
        }).
    Build()
_ = s.RegisterRule(r)
```

Eventually, having the scraper instantiated and configured you can use it to scrape any structure you want. Scraper returns a struct `model.Structure`.
```go
structure := s.Scrap(app)
```

### View

In order to render scraped structure, you will need to instantiate and configure a view.
View consists of:
* title
* component styles - styles are applied to the components by matching first of component tags with style ids
* additional styling (i.e. line color)

In order to instantiate default view, use the view builder:
```go
v := view.NewView().Build()
```

In case you need to customize it, use available builder methods:
```go
v := view.NewView().
    WithComponentStyle(
        view.NewComponentStyle("TAG").
            WithBackgroundColor(color.White).
            WithFontColor(color.Black).
            Build(),
    ).
    Build()
```

As the view is initialized, you can now render the structure into planUML diagram.
```go
outFile, _ := os.Create("c4.plantuml")
defer func() {
    _ = outFile.Close()
}()

err = v.RenderTo(structure, outFile)
```

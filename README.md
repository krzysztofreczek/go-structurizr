# go-structurizr
This library allows you to auto-generate C4 component diagrams out from the Golang code.

## How it works?
The library provides a set of tools (Scraper and View) that allow you to scrape and render given Golang structures into a [C4 component](https://c4model.com/) diagram in [*.plantuml](https://plantuml.com/) format.

**Scraper** component reflects given structure in accordance with structure interfaces, predefined rules and configuration. 

You may pass the scraped structure into a **View** definition which you can then render into a plantUML diagram code. 

Scraper identifies components to scrape in one of the following cases:
* type that is being examined implements an interface `model.HasInfo`.
* type that is being examined applies to one of the rules registered in the scraper.

## Components

### Component Info

Structure `model.Info` is a basic structure that defines a component included in the scraped structure of your code.
```go
type Info struct {
	Kind        string      // kind of scraped component
	Name        string      // component name
	Description string      // component description
	Technology  string      // technology used within the component
	Tags        []string    // tags are used to match view styles to component
}
```

### Scraper

Scraper may be instantiated in one of two ways:
* from the code
* from the YAML file

In order to instantiate the scraper you need to provide scraper configuration which contains a slice of prefixes of packages that you want to reflect. Types that do not match any of the given prefixes will not be traversed. 
```go
config := scraper.NewConfiguration(
    "github.com/org/pkg",
)
s := scraper.NewScraper(config)
```

Having a scraper instantiated, you can register a set of rules that will allow the scraper to identify the components to include in the output structure.

Each rule consists of:
* Set of package regexp's - only types in a package matching at least one of the package regexp's will be processed
* Name regexp - only type of name matching regexp will be processed
* Apply function - function that produces `model.Info` describing the component included in the scraped structure.

```go
r, err := scraper.NewRule().
    WithPkgRegexps("github.com/org/pkg/foo/.*").
    WithNameRegexp("^.*Client$").
    WithApplyFunc(
        func(name string, _ ...string) model.Info {
            return model.ComponentInfo(name, "foo client", "gRPC", "TAG")
        }).
    Build()
err = s.RegisterRule(r)
```

The apply function has two arguments: name and groups matched from the name regular expression. 

See the example:
```go
r, err := scraper.NewRule().
    WithPkgRegexps("github.com/org/pkg/foo/.*").
    WithNameRegexp(`^(\w*)\.(\w*)Client$`).
    WithApplyFunc(
        func(_ string, groups ...string) model.Info {
            // Do some groups sanity checks first, then:
            n := fmt.Sprintf("Client of external %s service", groups[1])
            return model.ComponentInfo(n, "foo client", "gRPC", "TAG")
        }).
    Build()
err = s.RegisterRule(r)
```

Alternatively, you can instantiate the scraper form YAML configuration file:
```yaml
// go-structurizr.yml
configuration:
  pkgs:
    - "github.com/org/pkg"

rules:
  - name_regexp: "^.*Client$"
    pkg_regexps:
      - "github.com/org/pkg/foo/.*"
    component:
      description: "foo client"
      technology: "gRPC"
      tags:
        - TAG
```

Regex groups may also be used within yaml rule definition. Here you can find an example:
```yaml
rules:
  - name_regexp: "(\\w*)\\.(\\w*)Client$"
    pkg_regexps:
      - "github.com/org/pkg/foo/.*"
    component:
      name: "Client of external {1} service"
      description: "foo client"
      technology: "gRPC"
      tags:
        - TAG
```

```go
s, err := scraper.NewScraperFromConfigFile("./go-structurizr.yml")
```

Eventually, having the scraper instantiated and configured you can use it to scrape any structure you want. Scraper returns a struct `model.Structure`.
```go
structure := s.Scrape(app)
```

### View

Similarly, to the scraper, view may be instantiated in one of two ways:
* from the code
* from the YAML file

In order to render scraped structure, you will need to instantiate and configure a view.
View consists of:
* title
* component styles - styles are applied to the components by matching first of component tags with style ids
* additional styling (i.e. line color)
* component tags - if specified, view will contain only components tagged with one of the view tags. When no tag is defined, all components will be included in the rendered view.
* root component tags - if specified, view will contain only those components which have connection (direct or in-direct) to at least one of components with root tag.

In order to instantiate default view, use the view builder:
```go
v := view.NewView().Build()
```

In case you need to customize it, use available builder methods:
```go
v := view.NewView().
    WithTitle("Title")
    WithComponentStyle(
        view.NewComponentStyle("TAG").
            WithBackgroundColor(color.White).
            WithFontColor(color.Black).
            WithBorderColor(color.Black).
            WithShape("database").
            Build(),
    ).
    WithComponentTag("TAG").
    WithRootComponentTag("ROOT").
    Build()
```

Alternatively, you can instantiate the view form YAML configuration file:
```yaml
// go-structurizr.yml
view:
  title: "Title"
  line_color: 000000ff
  styles:
    - id: TAG
      background_color: ffffffff
      font_color: 000000ff
      border_color: 000000ff
      shape: database
  root_component_tags:
    - ROOT
  component_tags:
    - TAG
```

```go
v, err := view.NewViewFromConfigFile("./go-structurizr.yml")
```

As the view is initialized, you can now render the structure into planUML diagram.
```go
outFile, _ := os.Create("c4.plantuml")
defer func() {
    _ = outFile.Close()
}()

err = v.RenderStructureTo(structure, outFile)
```

## Examples
You may find a couple of examples implemented in the `cmd` directory. In order to run any of those examples, please run the shell script attached.

## Good practices
The best results and experience in using the library will be ensured by enforcing the following practices:
- Having a solid and well-organized application context following clean-architecture principles will make your diagrams simple and easy to read. Also, this will allow you to create a short list of component types and styles.
- Following consistent naming conventions will help you in creating simple and straight-forward scraper rules.

## Full code documentation
https://pkg.go.dev/github.com/krzysztofreczek/go-structurizr

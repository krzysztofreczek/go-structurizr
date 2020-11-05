package scraper_test

import (
	"fmt"
	"testing"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
)

const pkg = "github.com/karhoo/go-structurizr"

func TestScraper_Scrap(t *testing.T) {
	sConfig := scraper.NewConfiguration(pkg)
	s := scraper.NewScraper(sConfig)

	app := NewApp()
	out := s.Scrap(app)

	fmt.Print(out)
	fmt.Print(s.RenderGraphViz())
}

type App struct {
	Command    PublicCommandWrapper
	CommandBis PublicCommandWrapperWithPrivateCommand
	Query      PublicQueryWrapper
	Repository Repository
	Client     Client
}

func NewApp() *App {
	repo := Repository{}
	client := Client{
		component: &UndocumentedComponent{},
	}
	return &App{
		Command: PublicCommandWrapper{
			Command: &command{
				Repository: &repo,
				Client:     &client,
			},
		},
		CommandBis: PublicCommandWrapperWithPrivateCommand{
			command: &commandBis{
				Repository: &repo,
				Client:     &client,
			},
		},
		Query: PublicQueryWrapper{
			Query: &query{
				Repository: &repo,
				Client:     &client,
			},
		},
		Repository: repo,
		Client:     client,
	}
}

type PublicCommandWrapper struct {
	Command *command
}

type command struct {
	Repository *Repository
	Client     *Client
}

func (c command) Info() model.Info {
	return model.ComponentInfo("command")
}

type PublicCommandWrapperWithPrivateCommand struct {
	command *commandBis
}

type commandBis struct {
	Repository *Repository
	Client     *Client
}

func (c commandBis) Info() model.Info {
	return model.ComponentInfo("command-bis")
}

type PublicQueryWrapper struct {
	Query *query
}

type query struct {
	Repository *Repository
	Client     *Client
}

func (q *query) Info() model.Info {
	return model.ComponentInfo("query")
}

type Repository struct{}

func (r Repository) Info() model.Info {
	return model.ComponentInfo("repository")
}

type Client struct {
	component *UndocumentedComponent
}

func (c Client) Info() model.Info {
	return model.ComponentInfo("client")
}

type UndocumentedComponent struct{}

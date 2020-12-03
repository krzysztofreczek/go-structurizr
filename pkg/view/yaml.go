package view

import (
	"encoding/hex"
	"image/color"
	"log"

	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
)

func toView(c yaml.Config) (View, error) {
	v := NewView().WithTitle(c.View.Title)

	if c.View.LineColor != "" {
		col, err := decodeHexColor(c.View.LineColor)
		if err != nil {
			return view{}, err
		}
		v.WithLineColor(col)
	}

	for _, s := range c.View.Styles {
		style := NewComponentStyle(s.ID)

		if s.BackgroundColor != "" {
			col, err := decodeHexColor(s.BackgroundColor)
			if err != nil {
				return view{}, err
			}
			style.WithBackgroundColor(col)
		}

		if s.FontColor != "" {
			col, err := decodeHexColor(s.FontColor)
			if err != nil {
				return view{}, err
			}
			style.WithFontColor(col)
		}

		if s.BorderColor != "" {
			col, err := decodeHexColor(s.BorderColor)
			if err != nil {
				return view{}, err
			}
			style.WithBorderColor(col)
		}

		if s.Shape != "" {
			style.WithShape(s.Shape)
		}

		v.WithComponentStyle(style.Build())
	}

	for _, t := range c.View.ComponentTags {
		v.WithComponentTag(t)
	}

	for _, t := range c.View.RootComponentTags {
		v.WithRootComponentTag(t)
	}

	return v.Build(), nil
}

func decodeHexColor(s string) (color.Color, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	return color.RGBA{R: b[0], G: b[1], B: b[2], A: 255}, nil
}

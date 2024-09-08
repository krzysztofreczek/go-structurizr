package model

// HasInfo represents a simple getter method that returns component information.
//
// The HasInfo interface indicates that a type can provide its own component
// information. All types that implement this interface are automatically
// detected by the default implementation of the scraper.
type HasInfo interface {
	Info() Info
}

const (
	infoKindComponent = "component"
)

// Info struct contains details about a component.
//
// Name is the name of the component.
// Kind represents the component's level or type in C4 diagrams.
// Description provides an explanation of the component's responsibility.
// Technology describes the technology the component is based on.
// Tags is a set of generic strings used to group and reference components.
type Info struct {
	Name        string
	Kind        string
	Description string
	Technology  string
	Tags        []string
}

// ComponentInfo creates a new component with a predefined kind "component".
// Variadic arguments are assigned sequentially to the remaining Info properties.
func ComponentInfo(s ...string) Info {
	return info(infoKindComponent, s...)
}

func info(kind string, s ...string) Info {
	info := Info{
		Kind: kind,
		Tags: make([]string, 0),
	}

	if len(s) > 0 {
		info.Name = s[0]
	}

	if len(s) > 1 {
		info.Description = s[1]
	}

	if len(s) > 2 {
		info.Technology = s[2]
	}

	for i, tag := range s {
		if i > 2 {
			info.Tags = append(info.Tags, tag)
		}
	}

	return info
}

// IsZero checks whether the component is empty.
// A component is considered empty if it does not have a specified kind.
func (i Info) IsZero() bool {
	return i.Kind == ""
}

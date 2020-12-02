package model

// HasInfo wraps simple getter method returning component information.
//
// HasInfo interface informs that the type is able to provide component
// information on its own.
// All the types that implement the interface are automatically detected
// by default implementation of the scraper.
type HasInfo interface {
	Info() Info
}

const (
	infoKindComponent = "component"
)

// Info struct contains all component information details.
//
// Name is a component name.
// Kind is a type that reflects component level in terms of C4 diagrams.
// Description explains the responsibility of the component.
// Technology describes technology that the component is based on.
// Tags is a set of generic string tags that may be used as reference
// to a group of components.
type Info struct {
	Name        string
	Kind        string
	Description string
	Technology  string
	Tags        []string
}

// ComponentInfo instantiates a new component of predefined kind "component".
// Variadic arguments are assigned to the rest of Info properties one-by-one.
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

// IsZero informs if the component is empty.
// If component has no kind specified it is considered as empty.
func (i Info) IsZero() bool {
	return i.Kind == ""
}

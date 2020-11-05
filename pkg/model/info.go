package model

type HasInfo interface {
	Info() Info
}

const (
	infoKindComponent = "component"
)

type Info struct {
	Kind        string
	Description string
	Technology  string
	Tags        []string
}

func ComponentInfo(s ...string) Info {
	return info(infoKindComponent, s...)
}

func info(kind string, s ...string) Info {
	info := Info{
		Kind: kind,
		Tags: make([]string, 0),
	}

	if len(s) > 0 {
		info.Description = s[0]
	}

	if len(s) > 1 {
		info.Technology = s[1]
	}

	for i, tag := range s {
		if i > 1 {
			info.Tags = append(info.Tags, tag)
		}
	}

	return info
}

func (i Info) IsZero() bool {
	return i.Kind == ""
}

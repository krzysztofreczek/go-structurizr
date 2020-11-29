package model

type HasInfo interface {
	Info() Info
}

const (
	infoKindComponent = "component"
)

type Info struct {
	Name        string
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

func (i Info) IsZero() bool {
	return i.Kind == ""
}

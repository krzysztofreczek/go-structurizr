package scraper

import (
	"regexp"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/pkg/errors"
)

var (
	matchAllRegexp = regexp.MustCompile("^.*$")
)

type RuleApplyFunc func() model.Info

type Rule interface {
	Applies(pkg string, name string) bool
	Apply() model.Info
}

type rule struct {
	pkgRegexes []*regexp.Regexp
	nameRegex  *regexp.Regexp
	applyFunc  RuleApplyFunc
}

func newRule(
	pkgRegexes []*regexp.Regexp,
	nameRegex *regexp.Regexp,
	applyFunc RuleApplyFunc,
) (rule, error) {
	if len(pkgRegexes) == 0 {
		return rule{}, errors.New(
			"at least one package expression must be provided",
		)
	}

	for _, rgx := range pkgRegexes {
		if rgx == nil {
			return rule{}, errors.New(
				"any of the package expression must not be nil",
			)
		}
	}

	if nameRegex == nil {
		return rule{}, errors.New(
			"name expression must be provided",
		)
	}

	if applyFunc == nil {
		return rule{}, errors.New(
			"applyFunc function must be defined",
		)
	}

	return rule{
		pkgRegexes: pkgRegexes,
		nameRegex:  nameRegex,
		applyFunc:  applyFunc,
	}, nil
}

func (r rule) Applies(pkg string, name string) bool {
	return r.nameApplies(name) && r.pkgApplies(pkg)
}

func (r rule) Apply() model.Info {
	return r.applyFunc()
}

func (r rule) pkgApplies(pkg string) bool {
	pkgApplies := false
	for _, rgx := range r.pkgRegexes {
		applies := rgx.MatchString(pkg)
		if applies {
			pkgApplies = true
			break
		}
	}
	return pkgApplies
}

func (r rule) nameApplies(name string) bool {
	return r.nameRegex.MatchString(name)
}

type RuleBuilder struct {
	pkgRegexes []string
	nameRegex  string
	applyFunc  RuleApplyFunc
}

func NewRule() *RuleBuilder {
	return &RuleBuilder{}
}

func (b *RuleBuilder) WithPkgRegexps(rgx ...string) *RuleBuilder {
	for _, r := range rgx {
		b.pkgRegexes = append(b.pkgRegexes, r)
	}
	return b
}

func (b *RuleBuilder) WithNameRegexp(rgx string) *RuleBuilder {
	b.nameRegex = rgx
	return b
}

func (b *RuleBuilder) WithApplyFunc(f RuleApplyFunc) *RuleBuilder {
	b.applyFunc = f
	return b
}

func (b RuleBuilder) Build() (Rule, error) {
	pkgRegexes := make([]*regexp.Regexp, 0)
	for _, rgx := range b.pkgRegexes {
		r, err := regexp.Compile(rgx)
		if err != nil {
			return nil, errors.Wrapf(err,
				"could not compile package expression `%s` as correct regular expression", rgx)
		}
		pkgRegexes = append(pkgRegexes, r)
	}

	if len(pkgRegexes) == 0 {
		pkgRegexes = append(pkgRegexes, matchAllRegexp)
	}

	nameRegex := matchAllRegexp
	if b.nameRegex != "" {
		r, err := regexp.Compile(b.nameRegex)
		if err != nil {
			return nil, errors.Wrapf(err,
				"could not compile name expression `%s` as correct regular expression", b.nameRegex)
		}
		nameRegex = r
	}

	return newRule(
		pkgRegexes,
		nameRegex,
		b.applyFunc,
	)
}

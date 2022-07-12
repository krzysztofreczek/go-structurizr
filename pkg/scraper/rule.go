package scraper

import (
	"regexp"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/pkg/errors"
)

var (
	matchAllRegexp = regexp.MustCompile("^.*$")
)

// RuleApplyFunc defines a signature of method returning
// a component information of type model.Info.
//
// Arguments:
// - name is a scraped name of the type in format `package.TypeName`
// - groups is a slice of sub-groups resolved from the rule name
// regular expression
type RuleApplyFunc func(
	name string,
	groups ...string,
) model.Info

// Rule defines an interface of any rule that maybe registered within scraper.
//
// Applies informs if rule should be applied to the given component
// considering its full package name and type name in format `package.TypeName`.
// Apply returns a component information of type model.Info based
// on the type name in format `package.TypeName`.
type Rule interface {
	Applies(
		pkg string,
		name string,
	) bool
	Apply(
		name string,
	) model.Info
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

// Applies informs if rule should be applied to the given component
// considering its full package name and type name in format `package.TypeName`.
//
// Component will be recognised as applicable when all of the following
// conditions are met:
// - package matches at least one of the rule package regular expressions
// - name matches the rule name regular expression
func (r rule) Applies(
	pkg string,
	name string,
) bool {
	return r.nameApplies(name) && r.pkgApplies(pkg)
}

// Apply returns a component information of type model.Info based
// on the type name in format `package.TypeName`.
// Apply will return result of registered RuleApplyFunc application function
// passing the following arguments:
// - name is a scraped name of the type in format `package.TypeName`
// - groups is a slice of sub-groups resolved from the rule name
// regular expression
func (r rule) Apply(
	name string,
) model.Info {
	regex := r.nameRegex

	groups := regex.FindAllStringSubmatch(name, -1)
	if len(groups) != 0 && len(groups[0]) > 1 {
		return r.applyFunc(groups[0][0], groups[0][1:]...)
	}

	return r.applyFunc(name)
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

// RuleBuilder simplifies instantiation of default Rule implementation.
//
// WithPkgRegexps sets a list of package regular expressions.
// WithNameRegexp sets name regular expression.
// WithApplyFunc sets rule application function RuleApplyFunc.
//
// Build returns Rule implementation constructed from the provided expressions
// and application function.
// Build will return an error if at least one of the provided expressions
// is invalid and cannot be compiled.
// Build will return an error if application function RuleApplyFunc is missing.
type RuleBuilder interface {
	WithPkgRegexps(rgx ...string) RuleBuilder
	WithNameRegexp(rgx string) RuleBuilder
	WithApplyFunc(f RuleApplyFunc) RuleBuilder

	Build() (Rule, error)
}

type ruleBuilder struct {
	pkgRegexes []string
	nameRegex  string
	applyFunc  RuleApplyFunc
}

// NewRule returns an empty RuleBuilder.
func NewRule() RuleBuilder {
	return &ruleBuilder{}
}

// WithPkgRegexps sets a list of package regular expressions.
func (b *ruleBuilder) WithPkgRegexps(rgx ...string) RuleBuilder {
	b.pkgRegexes = append(b.pkgRegexes, rgx...)
	return b
}

// WithNameRegexp sets name regular expression.
func (b *ruleBuilder) WithNameRegexp(rgx string) RuleBuilder {
	b.nameRegex = rgx
	return b
}

// WithApplyFunc sets rule application function RuleApplyFunc.
func (b *ruleBuilder) WithApplyFunc(f RuleApplyFunc) RuleBuilder {
	b.applyFunc = f
	return b
}

// Build returns Rule implementation constructed from the provided expressions
// and application function.
//
// In case no regular expression is provided either for name or package,
// those will be filled with regular expression matching all string "^.*$".
//
// Build will return an error if at least one of the provided expressions
// is invalid and cannot be compiled.
// Build will return an error if application function RuleApplyFunc is missing.
func (b ruleBuilder) Build() (Rule, error) {
	pkgRegexes := make([]*regexp.Regexp, 0)
	for _, rgx := range b.pkgRegexes {
		r, err := regexp.Compile(rgx)
		if err != nil {
			return nil, errors.Wrapf(err,
				"could not compile package expression `%s` "+
					"as correct regular expression", rgx)
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
				"could not compile name expression `%s` "+
					"as correct regular expression", b.nameRegex)
		}
		nameRegex = r
	}

	if b.applyFunc == nil {
		return nil, errors.New("apply function must be provided")
	}

	return newRule(
		pkgRegexes,
		nameRegex,
		b.applyFunc,
	)
}

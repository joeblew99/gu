package css

import (
	"bytes"
	"strings"
	"text/template"

	bcss "github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
)

// Rule defines the a single css rule which will be transformed and
// converted into a usable stylesheet during rendering.
type Rule struct {
	template *template.Template
	depends  []*Rule
}

// NewRule returns a new instance of a Rule which provides capability to parse
// and extrapolate the giving content using the provided binding.
func NewRule(rules string, rs ...*Rule) *Rule {
	tmp, err := template.New("css").Parse(rules)
	if err != nil {
		panic(err)
	}

	return &Rule{
		template: tmp,
		depends:  rs,
	}
}

// Stylesheet returns the provided styles using the binding as the argument for the
// provided css template.
func (r *Rule) Stylesheet(bind interface{}, parentNode string) (*bcss.Stylesheet, error) {
	var stylesheet bcss.Stylesheet

	for _, rule := range r.depends {
		sheet, err := rule.Stylesheet(bind, parentNode)
		if err != nil {
			return nil, err
		}

		stylesheet.Rules = append(stylesheet.Rules, sheet.Rules...)
	}

	var content bytes.Buffer
	if err := r.template.Execute(&content, bind); err != nil {
		return nil, err
	}

	sheet, err := parser.Parse(content.String())
	if err != nil {
		return nil, err
	}

	for _, rule := range sheet.Rules {
		r.morphRule(rule, parentNode)
	}

	stylesheet.Rules = append(stylesheet.Rules, sheet.Rules...)

	return &stylesheet, nil
}

// morphRules adjusts the provided rules with the parent selector.
func (r *Rule) morphRule(base *bcss.Rule, parentNode string) {
	for index, sel := range base.Selectors {
		var newSel string

		switch {
		case strings.HasPrefix(sel, ":"):
			newSel = parentNode + "" + sel
		default:
			newSel = parentNode + " " + sel
		}

		base.Selectors[index] = newSel
	}

	for _, rule := range base.Rules {
		if rule.Kind == bcss.AtRule {
			r.morphRule(rule, parentNode)
			continue
		}

		for index, sel := range rule.Selectors {
			var newSel string

			switch {
			case strings.HasPrefix(sel, ":"):
				newSel = parentNode + "" + sel
			default:
				newSel = parentNode + " " + sel
			}

			rule.Selectors[index] = newSel
		}
	}
}

// CSS takes a giving css ruleset with the provided context and returns a
// provided stylesheet, for usage.
// func CSS()

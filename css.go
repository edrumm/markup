package markup

import "strings"

/*
	Represents a CSS declaration
*/
type Declaration struct {
	property string
	value    string
}

/*
	Represents a CSS ruleset
*/
type Rule struct {
	selector     string
	declarations []*Declaration
}

/*
	Represents a full Cascading Stylesheet file
*/
type Stylesheet struct {
	name  string
	rules []*Rule
	/*
	   TODO
	*/
}

/*
	Creates new stylesheet
*/
func NewStylesheet(name string) *Stylesheet {
	return &Stylesheet{name, nil}
	/*
	   TODO
	*/
}

/*
	Adds rule to current stylesheet
*/
func (css *Stylesheet) AddRule(selector string) *Rule {
	r := &Rule{selector, nil}
	css.rules = append(css.rules, r)

	return r
}

/*
	Adds a declaration to current rule
*/
func (r *Rule) AddDeclaration(prop, value string) {
	r.declarations = append(r.declarations, &Declaration{prop, value})
}

/*
	Returns string of CSS syntax
*/
func (r *Rule) CSS() string {
	sb := strings.Builder{}

	/*
		TODO
	*/

	return sb.String()
}

package markup

type Dec struct {
	property string
	value    string
}

type Rule struct {
	selector     string
	declarations []*Dec
}

type Stylesheet struct {
	name  string
	rules []*Rule
	/*
	   TODO
	*/
}

func NewStylesheet(name string) *Stylesheet {
	return &Stylesheet{name, nil}
	/*
	   TODO
	*/
}

func (css *Stylesheet) AddRule(selector string) *Rule {
	/*
		How to link this to Stylesheet...??
	*/
	return &Rule{selector, nil}
}

func (r *Rule) AddDeclaration(prop, value string) {
	r.declarations = append(r.declarations, &Dec{prop, value})
}

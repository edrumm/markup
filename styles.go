package markup

type Declaration struct {
	property string
	value    string
}

type Rule struct {
	selector     string
	declarations []*Declaration
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

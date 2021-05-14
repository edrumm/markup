package markup

type Node interface {
	Html() string
}

type Div struct {
	/*
	   TODO
	*/
}

type HTML struct {
	title      string
	components []Node

	/*
	   TODO
	*/
}

func NewDiv( /* ... */ ) Node {
	return &Div{ /* ... */ }
}

func (*Div) Html() string {
	return "<div>...</div>"
}

/*

 */

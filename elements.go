package markup

import "fmt"

type Node interface {
	Html() string
}

/*
	Represents a generic HTML tag with no nested tags eg. title, etc...
	TODO - attributes
*/
type Tag struct {
	name    string
	content string
}

/*
	Represents html div element
*/
type Div struct {
	/*
	   TODO
	*/
}

/*
	Represents a full html page: file name, head, body
	TODO - add stylesheet []
	TODO - add script []
*/
type HTML struct {
	name string
	head []Node
	body []Node

	/*
	   TODO
	*/
}

/*
	Creates a new html struct, returns *HTML
*/
func NewHtmlDoc(name, title string) *HTML {
	h := []Node{NewTag("title", title)}
	return &HTML{name, h, nil}
}

/*
	Appends an HTML element to the head section
*/
func (h *HTML) AppendHead(elem Node) {
	h.head = append(h.head, elem)
}

/*
	Appends an HTML element to the body section
*/
func (h *HTML) AppendBody(elem Node) {
	h.body = append(h.body, elem)
}

func NewTag(name, content string) *Tag {
	return &Tag{name, content}
}

func (t *Tag) Html() string {
	return fmt.Sprintf("<%s>%s</%s>", t.name, t.content, t.name)
}

func NewDiv( /* ... */ ) Node {
	return &Div{ /* ... */ }
}

func (*Div) Html() string {
	return "<div>...</div>"
}

package markup

import "fmt"

type Node interface {
	Html() string
}

/*
	Represents a basic HTML tag eg. title, etc...
	TODO - attributes
*/
type Tag struct {
	name    string
	content string
}

/*
	Represents a single tag e.g. br, hr, etc...
*/
type SingleTag struct {
	name string
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
	lang string
	head []Node
	body []Node

	/*
	   TODO
	*/
}

/*
	Creates a new html struct, returns *HTML
*/
func NewHtmlDoc(name, lang, title string) *HTML {
	h := []Node{NewTag("title", title)}
	return &HTML{name, lang, h, nil}
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

func (h *HTML) Html() string {
	return fmt.Sprintf("<html lang=\"%s\">", h.lang)
}

func (t *SingleTag) Html() string {
	return fmt.Sprintf("<%s>", t.name)
}

func NewSingleTag(name string) *SingleTag {
	return &SingleTag{name}
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

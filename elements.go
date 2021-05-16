package markup

import (
	"fmt"
	"strings"
)

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
	content []Node
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
	HTML
*/
func NewHtmlDoc(name, lang, title string) *HTML {
	h := []Node{NewTag("title", title)}
	return &HTML{name, lang, h, nil}
}

func (h *HTML) AppendHead(elem Node) {
	h.head = append(h.head, elem)
}

func (h *HTML) AppendBody(elem Node) {
	h.body = append(h.body, elem)
}

func (h *HTML) Html() string {
	return fmt.Sprintf("<html lang=\"%s\">", h.lang)
}

/*
	Regular tag
*/
func NewTag(name, content string) *Tag {
	return &Tag{name, content}
}

func (t *Tag) Html() string {
	return fmt.Sprintf("<%s>%s</%s>", t.name, t.content, t.name)
}

/*
	Single tag
*/
func NewSingleTag(name string) *SingleTag {
	return &SingleTag{name}
}

func (t *SingleTag) Html() string {
	return fmt.Sprintf("<%s>", t.name)
}

/*
	Div
*/
func NewDiv( /* ... */ ) *Div {
	return &Div{nil /* ... */}
}

func (d *Div) Html() string {
	sb := strings.Builder{}
	sb.WriteString("<div>\n")

	for _, elem := range d.content {
		sb.WriteString(fmt.Sprintf("%s\n", elem.Html()))
	}

	sb.WriteString("</div>")

	return sb.String()
}

func (d *Div) AppendToDiv(n Node) {
	d.content = append(d.content, n)
}

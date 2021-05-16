package markup

import (
	"fmt"
	"strings"
)

type Node interface {
	Html() string
}

/*
	Represents a basic HTML tag
	TODO: attributes
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
	Represents a full html page
	TODO: add stylesheet []
	TODO: add script []
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
	Create new HTML struct
*/
func NewHtmlDoc(name, lang, title string) *HTML {
	h := []Node{NewTag("title", title)}
	return &HTML{name, lang, h, nil}
}

/*
	Append tag to HTML head
*/
func (h *HTML) AppendHead(elem Node) {
	h.head = append(h.head, elem)
}

/*
	Append tag to HTML body
*/
func (h *HTML) AppendBody(elem Node) {
	h.body = append(h.body, elem)
}

/*
	Return HTML syntax of tag
*/
func (h *HTML) Html() string {
	return fmt.Sprintf("<html lang=\"%s\">", h.lang)
}

/*
	Create new tag
*/
func NewTag(name, content string) *Tag {
	return &Tag{name, content}
}

/*
	Return HTML syntax of tag
*/
func (t *Tag) Html() string {
	return fmt.Sprintf("<%s>%s</%s>", t.name, t.content, t.name)
}

/*
	Create new single tag
	eg. hr, br
*/
func NewSingleTag(name string) *SingleTag {
	return &SingleTag{name}
}

/*
	Return HTML syntax of tag
*/
func (t *SingleTag) Html() string {
	return fmt.Sprintf("<%s>", t.name)
}

/*
	Create new div element
*/
func NewDiv( /* ... */ ) *Div {
	return &Div{nil /* ... */}
}

/*
	Return HTML syntax of tag, including all nested div tags
*/
func (d *Div) Html() string {
	sb := strings.Builder{}
	sb.WriteString("<div>\n")

	for _, elem := range d.content {
		sb.WriteString(fmt.Sprintf("%s\n", elem.Html()))
	}

	sb.WriteString("</div>")

	return sb.String()
}

/*
	Append tag to div
*/
func (d *Div) AppendToDiv(n Node) {
	d.content = append(d.content, n)
}

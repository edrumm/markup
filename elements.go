package markup

import (
	"fmt"
	"strings"
)

type Element interface {
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
	Represents HTML image (img) tag
*/
type Img struct {
	src    string
	width  int
	height int
	alt    string
}

/*
	Represents html div element
*/
type Div struct {
	content []Element
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
	head []Element
	body []Element

	/*
	   TODO
	*/
}

/*
	Create new HTML struct
*/
func NewHtmlDoc(name, lang, title string) *HTML {
	h := []Element{NewTag("title", title)}
	return &HTML{name, lang, h, nil}
}

/*
	Append tag to HTML head
*/
func (h *HTML) AppendHead(elem Element) {
	h.head = append(h.head, elem)
}

/*
	Append tag to HTML body
*/
func (h *HTML) AppendBody(elem Element) {
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
func (d *Div) AppendToDiv(n Element) {
	d.content = append(d.content, n)
}

/*
	Create new img element (no size specified)
*/
func NewImg(src, alt string) *Img {
	return &Img{src, 0, 0, alt}
}

/*
	Create new img element (with a specified size)
*/
func NewSizedImg(src, alt string, width, height int) *Img {
	return &Img{src, width, height, alt}
}

/*
	Return HTML syntax of img
*/
func (i *Img) Html() string {
	if i.height == 0 && i.width == 0 {
		return fmt.Sprintf(`<img src="%s" alt="%s">`, i.src, i.alt)
	}

	return fmt.Sprintf(`<img src="%s" alt="%s" width="%d" height="%d">`, i.src, i.alt, i.width, i.height)
}

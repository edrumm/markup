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
type HtmlTag struct {
	name    string
	content []string
}

/*
	Represents a single tag e.g. br, hr, etc...
*/
type HtmlSingleTag struct {
	name string
}

/*
	Represents an HTML hyperlink (a) tag
*/
type HtmlHyperlink struct {
	href string
	text string
	/*
		TODO - target attribute
	*/
}

/*
	Represents HTML image (img) tag
*/
type HtmlImg struct {
	src    string
	width  int
	height int
	alt    string
}

/*
	Represents html div element
*/
type HtmlDiv struct {
	children []Element
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
}

/*
	Create new HTML struct
*/
func HtmlPage(name, lang, title string) *HTML {
	h := []Element{Tag("title", title)}
	return &HTML{name, lang, h, nil}
}

/*
	Content tag to HTML head
*/
func (h *HTML) Head(elem ...Element) {
	h.head = append(h.head, elem...)
}

/*
	Content tag to HTML body
*/
func (h *HTML) Body(elem ...Element) {
	h.body = append(h.body, elem...)
}

/*
	Return HTML syntax of HTML tag
*/
func (h *HTML) Html() string {
	return fmt.Sprintf("<html lang=\"%s\">", h.lang)
}

/*
	Create new tag
*/
func Tag(name string, content ...string) *HtmlTag {
	return &HtmlTag{name, content}
}

/*
	Return HTML syntax of tag
*/
func (t *HtmlTag) Html() string {
	if len(t.content) == 1 {
		return fmt.Sprintf("<%s>%s</%s>", t.name, t.content[0], t.name)
	}

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("<%s>\n", t.name))

	for _, elem := range t.content {
		sb.WriteString(fmt.Sprintf("%s\n", elem))
	}

	sb.WriteString(fmt.Sprintf("</%s>", t.name))

	return sb.String()
}

/*
	Create new single tag
	eg. hr, br
*/
func SingleTag(name string) *HtmlSingleTag {
	return &HtmlSingleTag{name}
}

/*
	Return HTML syntax of tag
*/
func (t *HtmlSingleTag) Html() string {
	return fmt.Sprintf("<%s>", t.name)
}

/*
	Create new a tag (hyperlink)
*/
func Hyperlink(href, text string) *HtmlHyperlink {
	return &HtmlHyperlink{href, text}
}

/*
	Return HTML syntax of hyperlink
*/
func (a *HtmlHyperlink) Html() string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, a.href, a.text)
}

/*
	Create new div element
*/
func Div() *HtmlDiv {
	return &HtmlDiv{nil}
}

/*
	Return HTML syntax of div, including all nested tags
*/
func (d *HtmlDiv) Html() string {
	sb := strings.Builder{}
	sb.WriteString("<div>\n")

	for _, elem := range d.children {
		sb.WriteString(fmt.Sprintf("%s\n", elem.Html()))
	}

	sb.WriteString("</div>")

	return sb.String()
}

/*
	Content tag to div
*/
func (d *HtmlDiv) Content(n ...Element) {
	d.children = append(d.children, n...)
}

/*
	Create new img element (no size specified)
*/
func Img(src, alt string) *HtmlImg {
	return &HtmlImg{src, 0, 0, alt}
}

/*
	Create new img element (with a specified size)
*/
func ImgWithSize(src, alt string, width, height int) *HtmlImg {
	return &HtmlImg{src, width, height, alt}
}

/*
	Return HTML syntax of img
*/
func (i *HtmlImg) Html() string {
	if i.height == 0 && i.width == 0 {
		return fmt.Sprintf(`<img src="%s" alt="%s">`, i.src, i.alt)
	}

	return fmt.Sprintf(`<img src="%s" alt="%s" width="%d" height="%d">`, i.src, i.alt, i.width, i.height)
}

/*
	Add ID to tag
*/
func Id(e Element, id string) {
	/*
		TODO
	*/
}

/*
	Add class to tag
*/
func Class(e Element, class string) {
	/*
		TODO
	*/
}

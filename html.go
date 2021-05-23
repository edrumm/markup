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
	content string
}

/*
	Represents a single tag e.g. br, hr, etc...
*/
type HtmlSingleTag struct {
	name string
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
func HtmlPage(name, lang, title string) *HTML {
	h := []Element{Tag("title", title)}
	return &HTML{name, lang, h, nil}
}

/*
	Append tag to HTML head
*/
func (h *HTML) Head(elem Element) {
	h.head = append(h.head, elem)
}

/*
	Append tag to HTML body
*/
func (h *HTML) Body(elem Element) {
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
func Tag(name, content string) *HtmlTag {
	return &HtmlTag{name, content}
}

/*
	Return HTML syntax of tag
*/
func (t *HtmlTag) Html() string {
	return fmt.Sprintf("<%s>%s</%s>", t.name, t.content, t.name)
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
	Create new div element
*/
func Div( /* ... */ ) *HtmlDiv {
	return &HtmlDiv{nil /* ... */}
}

/*
	Return HTML syntax of tag, including all nested div tags
*/
func (d *HtmlDiv) Html() string {
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
func (d *HtmlDiv) AppendToDiv(n Element) {
	d.content = append(d.content, n)
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

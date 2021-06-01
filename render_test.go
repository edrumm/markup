package markup

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHTML_Render(t *testing.T) {
	index := HtmlPage("index.html", "en", "A Markup Page")

	index.Body(
		Tag("h1", "Auto Generated Page"),
		SingleTag("hr"),
		Tag("p", "An HTML page generated by Markup Go package"),
		ImgWithSize("1024px-Go_Logo_Blue.svg.png", "Go logo", 341, 128),
	)

	div := Div()
	div.Content(
		Tag("h2", "Div Header"),
		Tag("p", "Some div text"),
	)

	index.Body(div)
	index.Render()

	f, err := os.Open("index.html")

	defer func() {
		err := f.Close()
		if err != nil {
			t.Fail()
		}
	}()

	if err != nil {
		t.Fail()
	}
}

func TestTag_Html(t *testing.T) {
	h1 := Tag("h1", "Hello")
	assert.Equal(t, h1.Html(), "<h1>Hello</h1>")
}

func TestDiv_AppendToDiv(t *testing.T) {
	d := Div()
	d.Content(Tag("p", "Example text"))
	d.Content(Tag("p", "More example text"))

	assert.NotEmpty(t, d.children)
}

func TestSingleTag_Html(t *testing.T) {
	hr := SingleTag("hr")
	br := SingleTag("br")

	assert.Equal(t, hr.Html(), "<hr>")
	assert.Equal(t, br.Html(), "<br>")
}

func TestImg_Html(t *testing.T) {
	img := Img("myphoto.png", "An Image")
	assert.Equal(t, img.Html(), `<img src="myphoto.png" alt="An Image">`)
}

func TestHtmlHyperlink_Html(t *testing.T) {
	a := Hyperlink("https://github.com/", "GitHub")
	assert.Equal(t, a.Html(), `<a href="https://github.com/">GitHub</a>`)
}

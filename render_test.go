package markup

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHTML_Render(t *testing.T) {
	index := NewHtmlDoc("index.html", "en", "A Markup Page")

	index.AppendBody(NewTag("h1", "Auto Generated Page"))
	index.AppendBody(NewSingleTag("hr"))
	index.AppendBody(NewTag("p", "An HTML page generated by Markup Go package"))
	index.AppendBody(NewSizedImg("1024px-Go_Logo_Blue.svg.png", "Go logo", 341, 128))

	div := NewDiv()
	div.AppendToDiv(NewTag("h2", "Div Header"))
	div.AppendToDiv(NewTag("p", "Some div text"))

	index.AppendBody(div)
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
	h1 := NewTag("h1", "Hello")
	assert.Equal(t, h1.Html(), "<h1>Hello</h1>")
}

func TestDiv_AppendToDiv(t *testing.T) {
	d := NewDiv()
	d.AppendToDiv(NewTag("p", "Example text"))
	d.AppendToDiv(NewTag("p", "More example text"))

	assert.NotEmpty(t, d.content)
}

func TestSingleTag_Html(t *testing.T) {
	hr := NewSingleTag("hr")
	br := NewSingleTag("br")

	assert.Equal(t, hr.Html(), "<hr>")
	assert.Equal(t, br.Html(), "<br>")
}

func TestImg_Html(t *testing.T) {
	img := NewImg("myphoto.png", "An Image")
	assert.Equal(t, img.Html(), `<img src="myphoto.png" alt="An Image">`)
}

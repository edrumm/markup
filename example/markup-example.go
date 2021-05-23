package main

import "github.com/edrumm/markup"

func main() {
	/*
		Create new html page
		(*HTML)
	*/
	index := markup.HtmlPage("index.html", "en", "A Markup Page")

	// Append elements
	index.Body(markup.Tag("h1", "Auto Generated Page"))
	index.Body(markup.SingleTag("hr"))
	index.Body(markup.Tag("p", "An HTML page generated by Markup Go package"))
	index.Body(markup.ImgWithSize("1024px-Go_Logo_Blue.svg.png", "Go logo", 341, 128))

	div := markup.Div()
	div.AppendToDiv(markup.Tag("h2", "Div Header"))
	div.AppendToDiv(markup.Tag("p", "Some div text"))

	index.Body(div)

	/*
		Create new stylesheet
		(*Stylesheet)
	*/
	css := markup.NewStylesheet("style.css")

	// Render page to html file
	index.Render()

	// Render stylesheet to css file
	css.Render()
}

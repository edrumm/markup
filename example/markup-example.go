package main

import "github.com/edrumm/markup"

func main() {
	/*
		Create new html page
		(*HTML)
	*/
	idx := markup.NewHtmlDoc("index.html", "A Markup Page")

	/*
		Create new stylesheet
		(*Stylesheet)
	*/
	css := markup.NewStylesheet("style.css")

	// Render page to html file
	idx.Render()

	// Render stylesheet to css file
	css.Render()
}

package main

import "github.com/edrumm/markup"

func main() {
	// Create new html page
	// (* HTML)
	idx := markup.NewHtmlDoc("index.html", "A Markup Page")

	// Render page to html file
	idx.Render()
}

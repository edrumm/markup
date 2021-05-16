# markup
generate basic html and stylesheets with go

[https://pkg.go.dev/github.com/edrumm/markup]()

### Installation

Install markup
- `go get github.com/edrumm/markup`

Import
- `import "github.com/edrumm/markup"`

### Usage
More detailed description at: [https://pkg.go.dev/github.com/edrumm/markup]()

- Create a basic HTML file
```
package main

import "github.com/edrumm/markup"

func main() {
	// Create an HTML file
	index := markup.NewHtmlDoc("index.html", "en", "Markup Page")
	
	// Add HTML elements
	index.AppendBody(markup.NewTag("h1", "Auto Generated Page"))
	index.AppendBody(markup.NewSingleTag("hr"))
	index.AppendBody(markup.NewTag("p", "A page generated in Go using github.com/edrumm/markup"))
    
	// Render the HTML page to a file
	index.Render()
}
```

- HTML page with div
```
package main

import "github.com/edrumm/markup"

func main() {
	// Create an HTML file
	index := markup.NewHtmlDoc("index.html", "en", "Markup Page")
	
	// Create a new div
	div := markup.NewDiv()
	
	// Add elements to div
	div.AppendToDiv(markup.NewTag("h2", "Div Header"))
	div.AppendToDiv(markup.NewTag("p", "Some div text"))
	
	// Add div to HTML
	index.AppendBody(div)
    
	// Render the HTML page to a file
	index.Render()
}
```

### Coming Soon
- CSS stylesheets
- Script, img, & other HTML tags
- Nested tag support, eg. `<p><a>...</a></p>`

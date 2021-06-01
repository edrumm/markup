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
	index := markup.HtmlPage("index.html", "en", "Markup Page")
	
	// Add HTML elements
	index.Body(
	    markup.Tag("h1", "Auto Generated Page"),
	    markup.SingleTag("hr"),
	    markup.Tag("p", "A page generated in Go using github.com/edrumm/markup"),
    )
    
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
	index := markup.HtmlPage("index.html", "en", "Markup Page")
	
	// Create a new div
	div := markup.Div()
	
	// Add elements to div
	div.Content(
	    markup.Tag("h2", "Div Header"),
	    markup.Tag("p", "Some div text"),
	)
	
	// Add div to HTML
	index.Body(div)
    
	// Render the HTML page to a file
	index.Render()
}
```

### Coming Soon
- CSS stylesheets
- script, a, & other HTML tags
- [x] Nested tag support, eg. `<p><a>...</a></p>`
- [x] Variadic versions of `Head()` and `Body()`

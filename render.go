package markup

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Renders an html struct to an html file
*/
func (h *HTML) Render() {
	f, err := os.Create(h.name)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	w.WriteString("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n")
	for _, elem := range h.head {
		w.WriteString(fmt.Sprintf("%s\n", elem.Html()))
	}

	w.WriteString("</head>\n<body>\n")
	for _, elem := range h.body {
		w.WriteString(fmt.Sprintf("%s\n", elem.Html()))
	}

	w.WriteString("</body>\n</html>")
	w.Flush()
}

/*
	Renders a stylesheet struct to a css file
*/
func (*Stylesheet) Render() {
	fmt.Println("Under Construction...")

	/*
		TODO
	*/
}

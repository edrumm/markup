package markup

import (
	"fmt"
	"os"
)

func errf(err error) {
	if err != nil {
		panic(err)
	}
}

/*
	Renders an html struct to an html file
*/
func (h *HTML) Render() {
	f, err := os.Create(h.name)
	errf(err)

	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err = f.WriteString("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n")
	errf(err)

	for _, elem := range h.head {
		_, err = f.WriteString(fmt.Sprintf("%s\n", elem.Html()))
		errf(err)
	}

	f.WriteString("</head>\n<body>\n")
	for _, elem := range h.body {
		_, err = f.WriteString(fmt.Sprintf("%s\n", elem.Html()))
		errf(err)
	}

	_, err = f.WriteString("</body>\n</html>")
	errf(err)
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

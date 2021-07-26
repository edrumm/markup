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

	_, err = f.WriteString(fmt.Sprintf("<!DOCTYPE html>\n%s\n<head>\n", h.Html()))
	errf(err)

	for _, elem := range h.head {
		_, err = f.WriteString(fmt.Sprintf("%s\n", elem.Html()))
		errf(err)
	}

	_, err = f.WriteString("</head>\n<body>\n")
	errf(err)

	for _, elem := range h.body {
		_, err = f.WriteString(fmt.Sprintf("%s\n", elem.Html()))
		errf(err)
	}

	_, err = f.WriteString("</body>\n</html>")
	errf(err)
}

/*
	Renders a stylesheet struct to a css file

	Only partially implemented
*/
func (s *Stylesheet) Render() {
	f, err := os.Create(s.name)
	errf(err)

	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Under Construction...")

	/*
		TODO
	*/
}

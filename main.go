//go:generate pigeon -o xml.go xml.peg
//go:generate goimports -w xml.go

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type prettyPrinter interface {
	prettyPrint(w io.Writer, depth int)
}

type Element struct {
	name     string
	children []prettyPrinter
}

func (el Element) String() string {
	var buf bytes.Buffer
	el.prettyPrint(&buf, 0)
	return buf.String()
}

func (el Element) prettyPrint(w io.Writer, depth int) {
	indent(w, depth)
	fmt.Fprintf(w, "<%s>\n", el.name)

	for _, c := range el.children {
		c.prettyPrint(w, depth+1)
	}

	indent(w, depth)
	fmt.Fprintf(w, "</%s>\n", el.name)
}

type TextNode string

func (tn TextNode) prettyPrint(w io.Writer, depth int) {
	indent(w, depth)
	fmt.Fprintf(w, "%s\n", tn)
}

func indent(w io.Writer, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Fprint(w, "\t")
	}
}

func main() {
	doc, err := ParseReader("", os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc)
}

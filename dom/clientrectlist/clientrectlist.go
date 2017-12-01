package clientrectlist

import (
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/js"
)

// ClientRectList struct
// js:"ClientRectList,omit"
type ClientRectList struct {
}

// Item fn
// js:"item"
func (*ClientRectList) Item(index uint) (c *clientrect.ClientRect) {
	js.Rewrite("$_.item($1)", index)
	return c
}

// Length prop
// js:"length"
func (*ClientRectList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
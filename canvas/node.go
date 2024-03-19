package canvas

import (
	"fmt"
)

type Node struct {
	// generic fields
	ID     string  `json:"id"`
	Type   string  `json:"type"` // one of "text", "file", "link", "group"
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Color  *string `json:"color,omitempty"`
	// text specific fields
	Text *string `json:"text,omitempty"` // if type == "text"
	// file specific fields
	File    *string `json:"file,omitempty"`    // if type == "file"
	Subpath *string `json:"subpath,omitempty"` // if type == "file"
	// link specific fields
	URL *string `json:"url,omitempty"` // if type == "link"
	// group specific fields
	Label           *string `json:"label,omitempty"`           // if type == "group"
	Background      *string `json:"background,omitempty"`      // if type == "group"
	BackgroundStyle *string `json:"backgroundStyle,omitempty"` // if type == "group"
}

func (n *Node) Validate() error {
	if n == nil {
		return nil
	}

	if n.Type == "" {
		return fmt.Errorf("node type is required")
	} else if n.Type != "text" && n.Type != "file" && n.Type != "link" && n.Type != "group" {
		return fmt.Errorf("invalid node type: %s", n.Type)
	}

	if n.Type == "text" && (n.Text == nil || *n.Text == "") {
		return fmt.Errorf("text type node requires text attribute")
	} else if n.Type == "file" && (n.File == nil || *n.File == "") {
		return fmt.Errorf("file type node requires file attribute")
	} else if n.Type == "link" && (n.URL == nil || *n.URL == "") {
		return fmt.Errorf("link type node requires url attribute")
	} else if n.Type == "group" && n.Label == nil {
		return fmt.Errorf("group type node requires label attribute")
	}

	return nil
}

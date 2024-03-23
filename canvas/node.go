package canvas

import (
	"fmt"

	"github.com/supersonicpineapple/go-jsoncanvas/util"
)

const (
	DefaultWidth  = 250
	DefaultHeight = 60
	DefaultGap    = 30
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
	BackgroundStyle *string `json:"backgroundStyle,omitempty"` // if type == "group"; one of "cover", "ratio", "repeat"
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

	switch n.Type {
	case "text":
		return n.validateText()
	case "file":
		return n.validateFile()
	case "link":
		return n.validateLink()
	case "group":
		return n.validateGroup()
	default:
		return fmt.Errorf("invalid type: %s", n.Type)
	}
}

func (n *Node) validateText() error {
	if n.Text == nil {
		return fmt.Errorf("text type node requires text attribute")
	}
	return nil
}

func (n *Node) validateFile() error {
	if n.File == nil || *n.File == "" {
		return fmt.Errorf("file type node requires file attribute")
	}
	return nil
}

func (n *Node) validateLink() error {
	if n.URL == nil || *n.URL == "" {
		return fmt.Errorf("link type node requires url attribute")
	}
	return nil
}

func (n *Node) validateGroup() error {
	if n.Label == nil {
		return fmt.Errorf("group type node requires label attribute")
	}
	if n.BackgroundStyle != nil && *n.BackgroundStyle != "cover" && *n.BackgroundStyle != "ratio" && *n.BackgroundStyle != "repeat" {
		return fmt.Errorf("invalid background style: %s", *n.BackgroundStyle)
	}
	return nil
}

func NewNode() *Node {
	n := Node{
		ID:     util.NewID(),
		X:      0,
		Y:      0,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	}
	return &n
}

func (n *Node) SetText(s string) *Node {
	n.Type = "text"
	n.Text = &s
	return n
}

func (n *Node) SetFile(path string, subPath ...string) *Node {
	n.Type = "file"
	n.File = &path
	if len(subPath) > 0 {
		n.Subpath = &subPath[0]
	}
	return n
}

func (n *Node) SetLink(url string) *Node {
	n.Type = "link"
	n.URL = &url
	return n
}

func (n *Node) SetGroup(label string) *Node {
	n.Type = "group"
	n.Label = &label
	return n
}

func (n *Node) SetPosition(x, y int) *Node {
	n.X = x
	n.Y = y
	return n
}

func (n *Node) TranslateX(x int) *Node {
	n.X += x
	return n
}

func (n *Node) TranslateY(y int) *Node {
	n.Y += y
	return n
}

func (n *Node) SetWidth(width int) *Node {
	n.Width = width
	return n
}
func (n *Node) SetHeight(height int) *Node {
	n.Height = height
	return n
}

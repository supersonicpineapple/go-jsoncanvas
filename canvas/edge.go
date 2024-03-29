package canvas

import (
	"fmt"

	"github.com/supersonicpineapple/go-jsoncanvas/util"
)

type Edge struct {
	ID    string  `json:"id"`
	Color *string `json:"color,omitempty"`
	Label *string `json:"label,omitempty"`

	FromNode string  `json:"fromNode"`
	FromSide *string `json:"fromSide,omitempty"` // one of "top", "right", "bottom", "left"
	FromEnd  *string `json:"fromEnd,omitempty"`  // one of "none", "arrow"

	ToNode string  `json:"toNode"`
	ToSide *string `json:"toSide,omitempty"` // one of "top", "right", "bottom", "left"
	ToEnd  *string `json:"toEnd,omitempty"`  // one of "none", "arrow"
}

func (e *Edge) Validate() error {
	if e == nil {
		return nil
	}

	if e.FromNode == "" || e.ToNode == "" {
		return fmt.Errorf("fromNode and toNode are required")
	} else if e.FromNode == e.ToNode {
		return fmt.Errorf("fromNode and toNode cannot be the same node")
	}

	return nil
}

func NewEdge(from, to *Node, fromSide, toSide string) *Edge {
	e := Edge{
		ID:       util.NewID(),
		FromNode: from.ID,
		FromSide: &fromSide,
		ToNode:   to.ID,
		ToSide:   &toSide,
	}
	return &e
}

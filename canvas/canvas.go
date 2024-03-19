package canvas

import (
	"fmt"
	"strings"
)

type Canvas struct {
	Nodes []*Node `json:"nodes"`
	Edges []*Edge `json:"edges"`
}

func (c *Canvas) Validate() error {
	if c == nil {
		return nil
	}

	var nodeErrors, edgeErrors []error

	for _, node := range c.Nodes {
		if err := node.Validate(); err != nil {
			nodeErrors = append(nodeErrors, err)
		}
	}

	for _, edge := range c.Edges {
		if err := edge.Validate(); err != nil {
			edgeErrors = append(edgeErrors, err)
		}
	}

	if len(nodeErrors) > 0 || len(edgeErrors) > 0 {
		return fmt.Errorf("validation errors:\n%s\n%s", joinErrors(nodeErrors), joinErrors(edgeErrors))
	} else {
		return nil
	}
}

func joinErrors(errors []error) string {
	var sb strings.Builder

	for i, err := range errors {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, err))
	}

	return sb.String()
}

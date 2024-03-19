package canvas

import (
	"fmt"
)

type Color string

const (
	ColorRed    Color = "1"
	ColorOrange Color = "2"
	ColorYellow Color = "3"
	ColorGreen  Color = "4"
	ColorCyan   Color = "5"
	ColorPurple Color = "6"
)

func (c *Color) Validate() error {
	// check if color is a hex color or a preset color
	if c == nil || *c == "" {
		return nil
	}

	if (*c)[0] == '#' {
		hexCode := string(*c)[1:]
		return validateHexColor(hexCode)
	} else if *c == ColorRed || *c == ColorOrange || *c == ColorYellow || *c == ColorGreen || *c == ColorCyan || *c == ColorPurple {
		// preset color
		return nil
	} else {
		return fmt.Errorf("invalid color: %s", *c)
	}
}

func validateHexColor(code string) error {
	if len(code) != 6 {
		return fmt.Errorf("invalid hex color code length: #%s", code)
	}

	for _, c := range code {
		if c < '0' || c > '9' || c < 'a' || c > 'f' {
			return fmt.Errorf("invalid hex color code character: %c", c)
		}
	}

	return nil
}

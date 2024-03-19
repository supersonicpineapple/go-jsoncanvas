package jsoncanvas

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/supersonicpineapple/go-jsoncanvas/canvas"
)

var (
	disallowUnknownFields = true
)

func SetDisallowUnknownFields(v bool) {
	disallowUnknownFields = v
}

func Decode(r io.Reader) (*canvas.Canvas, error) {
	decoder := json.NewDecoder(r)
	if disallowUnknownFields {
		decoder.DisallowUnknownFields()
	}

	c := new(canvas.Canvas)
	if err := decoder.Decode(&c); err != nil {
		return nil, fmt.Errorf("can't decode canvas file: %w", err)
	}

	return c, nil
}

func Parse(path string) (*canvas.Canvas, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}
	defer f.Close()

	return Decode(f)
}

func Encode(c *canvas.Canvas, w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(c); err != nil {
		return fmt.Errorf("can't encode canvas object: %w", err)
	}
	return nil
}

// Serialize checks if the file exists
func Serialize(c *canvas.Canvas, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		// TODO: create file if it does not exist?
		return fmt.Errorf("can't stat file: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("got dir %s, please specify a file instead", path)
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_SYNC, 644)
	if err != nil {
		return fmt.Errorf("can't open file %s: %w", path, err)
	}
	defer f.Close()

	return Encode(c, f)
}

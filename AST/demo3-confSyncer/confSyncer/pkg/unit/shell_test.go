package unit

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestCopy(t *testing.T) {
	abs, err := filepath.Abs("$HOME/.config")
	if err != nil {
		panic(err)
	}
	fmt.Println(abs)
}

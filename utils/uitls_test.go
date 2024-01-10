package utils

import (
	"fmt"
	"testing"
)

func TestGenHash(t *testing.T) {
	str := "Cty"
	hashStr := GenHash(str)
	fmt.Println(hashStr)
}

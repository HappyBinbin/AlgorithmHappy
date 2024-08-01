package _42

import (
	"fmt"
	"testing"
)

func TestDuepulicate(t *testing.T) {

	duplicate := []int{4, 3, 2, 7, 4, 7, 3, 2}
	result := findDuplicates(duplicate)
	fmt.Printf("%+v", result)
}

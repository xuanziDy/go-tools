package slice

import (
	"fmt"
	"testing"
)

func TestIntSlice(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	res, err := Delete[int](s, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Printf("%+v\n", res)
}

func TestStringSlice(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	res, err := Delete[string](s, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Printf("%+v\n", res)
}

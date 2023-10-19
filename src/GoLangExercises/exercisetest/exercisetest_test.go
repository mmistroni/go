package exercisetest_test

import (
	"testing"
	"example.com/golang-exercises/exercisetest"
)

func TestQuack(t *testing.T) {
	f := "foo"
	b := "bar"

	str, num := exercisetest.Quack(f, b) 

	if  str !=  "foo_bar" && num !=1 {
		t.Fatal("Wrong Call :(")
	}
}



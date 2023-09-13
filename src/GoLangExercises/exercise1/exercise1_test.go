package exercise1_test
import (
	"testing"
	"example.com/golang-exercises/exercise1"
)
func TestMascot(t *testing.T) {
	if exercise1.Runner() != "Go Gopher" {
		t.Fatal("Wrong Mascot :(")
	}
}
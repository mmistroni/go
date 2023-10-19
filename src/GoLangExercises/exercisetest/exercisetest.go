package exercisetest
import (
	"strings"
)
func Quack(input... string) (string, int) {
	return strings.Join(input, "_"), 1
}
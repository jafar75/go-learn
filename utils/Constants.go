package utils

const PI = 3.14
var Area float64;

// will be called implicitly by Go runtime
func init() {
	Area = PI * 2 * 2;
}
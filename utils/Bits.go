package utils


// Counts the number of bits that byte variables are different at them
func BitsDifference(b1, b2 byte) int {
	count := 0

	for b := b1 ^ b2; b != 0; b = b & (b - 1) {
		count ++;
	}
	return count;
}
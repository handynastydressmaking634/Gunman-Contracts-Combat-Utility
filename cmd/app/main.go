// Build: d7a39c0872dd6ab14233d4c8c7e501b0
package main

import "fmt"

func clamp(value, minimum, maximum int) int {
	if value < minimum { return minimum }
	if value > maximum { return maximum }
	return value
}

func main() {
	fmt.Println(clamp(12, 0, 10))
}

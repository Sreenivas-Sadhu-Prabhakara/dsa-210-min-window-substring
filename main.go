package main

import "fmt"

// MinWindow returns the shortest substring of s containing every character
// of t with multiplicity ("" if none). Expand right until satisfied,
// contract left while still satisfied, track the best.
func MinWindow(s, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}
	var need, have [128]int
	required := 0 // distinct characters still unsatisfied
	for i := 0; i < len(t); i++ {
		if need[t[i]] == 0 {
			required++
		}
		need[t[i]]++
	}
	bestLen, bestStart := len(s)+1, 0
	left := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		have[c]++
		if need[c] > 0 && have[c] == need[c] {
			required--
		}
		for required == 0 { // window covers t: shrink to the minimal form
			if right-left+1 < bestLen {
				bestLen = right - left + 1
				bestStart = left
			}
			out := s[left]
			have[out]--
			if need[out] > 0 && have[out] < need[out] {
				required++
			}
			left++
		}
	}
	if bestLen > len(s) {
		return ""
	}
	return s[bestStart : bestStart+bestLen]
}

func main() {
	fmt.Println(MinWindow("ADOBECODEBANC", "ABC"))
}

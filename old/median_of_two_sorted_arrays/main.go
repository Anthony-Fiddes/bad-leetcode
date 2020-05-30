package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j int
	l1 := len(nums1)
	l2 := len(nums2)
	s := l1 + l2
	target := s / 2
	odd := s%2 == 1 // If it's odd, we need to look at the number that was before the target
	var prev int
	for {
		var lesser int
		switch {
		case i == l1:
			lesser = nums2[j]
			j++
		case j == l2:
			lesser = nums1[i]
			i++
		default:
			if nums1[i] < nums2[j] {
				lesser = nums1[i]
				i++
			} else {
				lesser = nums2[j]
				j++
			}
		}

		if i+j-1 == target {
			if !odd {
				return float64(prev+lesser) / 2.0
			}
			return float64(lesser)
		}
		prev = lesser
	}
}

func main() {
	findMedianSortedArrays([]int{1, 2}, []int{3})
}

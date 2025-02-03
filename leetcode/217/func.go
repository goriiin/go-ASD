package main

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]struct{}, len(nums))

	for _, elem := range nums {
		_, ok := numsMap[elem]
		if ok {
			return false
		}

		numsMap[elem] = struct{}{}
	}

	return true

}

package groupthepeople

// groupThePeople
// There are n people that are split into some unknown number of groups. Each
// person is labeled with a unique ID from 0 to n - 1.
// You are given an integer array groupSizes, where groupSizes[i] is the size of
// the group that person i is in. For example, if groupSizes[1] = 3, then person 1
// must be in a group of size 3.
// Return a list of groups such that each person i is in a group of size
// groupSizes[i].
// Each person should appear in exactly one group, and every person must be in a
// group. If there are multiple answers, return any of them. It is guaranteed that
// there will be at least one valid solution for the given input.
func groupThePeople(groupSizes []int) [][]int {
	// Each key in groupHash is a group size, and each associated value is a group
	// of all groups of that group size.
	groupHash := map[int][][]int{}
	for personID, groupSize := range groupSizes {
		// Fetch the list of groups for the given size
		groups, ok := groupHash[groupSize]
		if !ok {
			groups = [][]int{[]int{}}
			groupHash[groupSize] = groups
		}

		// See if the last group has room for this person. If not, make a new one.
		numGroups := len(groups)
		lastGroup := &groups[numGroups-1]
		if len(*lastGroup) == groupSize {
			groupHash[groupSize] = append(groups, []int{})
			lastGroup = &groupHash[groupSize][numGroups]
		}
		*lastGroup = append(*lastGroup, personID)
	}

	// Collect all the completed groups into a single slice and return
	resultGroups := [][]int{}
	for _, groups := range groupHash {
		for _, group := range groups {
			resultGroups = append(resultGroups, group)
		}
	}
	return resultGroups
}

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
package kidswithcandies

// kidsWithCandies
// Given the array candies and the integer extraCandies, where candies[i]
// represents the number of candies that the ith kid has.
// For each kid check if there is a way to distribute extraCandies among the kids
// such that he or she can have the greatest number of candies among them. Notice
// that multiple kids can have the greatest number of candies.
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for i, x := range candies {
		if i == 0 || x > max {
			max = x
		}
	}
	result := []bool{}
	for _, x := range candies {
		result = append(result, x+extraCandies >= max)
	}
	return result
}

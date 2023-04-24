package numrescueboats

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	lo, hi := 0, len(people)-1
	var boat int
	for hi >= lo {
		if people[lo]+people[hi] <= limit {
			lo++
		}
		hi--
		boat++
	}

	return boat
}

func main() {

}

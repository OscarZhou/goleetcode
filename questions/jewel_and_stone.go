package questions

import (
	"fmt"

	track "github.com/OscarZhou/gotrack"
)

type JewelAndStone struct {
	Question
	Track *track.Track
}

func (q *JewelAndStone) Init() {
	q.No = 771
	q.Title = "Jewels and Stones"
	q.FullTitle = "Jewels and Stones"
	q.URL = "https://leetcode.com/problems/jewels-and-stones/"
	q.Level = LevelEasy
	q.FuncName = "numJewelsInStones"
}

func (q JewelAndStone) Print() {
	fmt.Println(`
	func numJewelsInStones(J string, S string) int {
		count := 0
		for _, s := range []rune(S) {
			for _, j := range []rune(J) {
				if s == j {
					count++
					break
				}
			}
		}
		return count
	}
	`)

}

func (q JewelAndStone) PrintTitle() {
	fmt.Printf("%d, %s\n", q.No, q.Title)
}

func (q JewelAndStone) Run() error {
	return nil
}

func numJewelsInStones(J string, S string) int {
	count := 0
	for _, s := range []rune(S) {
		for _, j := range []rune(J) {
			if s == j {
				count++
				break
			}
		}
	}
	return count
}

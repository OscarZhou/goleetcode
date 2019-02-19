package questions

var (
	NumberMap map[int]Questioner
	TitleMap  map[string]Questioner
)

func init() {
	NumberMap = make(map[int]Questioner)
	TitleMap = make(map[string]Questioner)
}

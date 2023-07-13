package title

func MakeTitle(text string) string {
	var bar, result string
	bar = "================================="
	result = bar + "/n" + text + "/n" + bar
	return result
}

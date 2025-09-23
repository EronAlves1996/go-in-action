package textprocessor

type DummyProcessor struct{}

func (d DummyProcessor) ReverseWords(input string) string {
	return ""
}

func (d *DummyProcessor) CountyVowels(input string) int {
	return 0
}

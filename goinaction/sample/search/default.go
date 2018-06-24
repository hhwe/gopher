package search

type defaultMatch struct{}

func init() {
	var matcher defaultMatch
	Register("default", matcher)
}

func (m defaultMatch) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

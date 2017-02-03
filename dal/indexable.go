package dal

type Indexable interface {
	MatchKeywords(keywords []string) bool
}

//TODO: there should be a library for this already
func FilterCompany(s []Company, fn func(Company) bool) []Company {
	var p = make([]Company, 0) // More efficient?
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func FilterComment(s []Comment, fn func(Comment) bool) []Comment {
	var p = make([]Comment, 0) // More efficient?
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

package dal

import (
	"fmt"
	"strings"
)


//go:generate mockery -name=CommentService -outpkg=mocks
type CommentService interface {
	Get(id uint64) Comment
	Search(keywords string) []Comment
	Update(comment Comment) Comment
	Finalize()
}

type Comment struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

func (c *Comment) MatchKeywords(keywords []string) bool {
	return strings.Contains(c.Author, keywords[0]) || strings.Contains(c.Text, keywords[0]) //TODO: ghetto!
}

func (c *Comment) String() string {
	return fmt.Sprintf("Comment[%d,%s,%s]",c.ID,c.Author,c.Text)
}

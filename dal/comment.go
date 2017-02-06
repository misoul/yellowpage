package dal

import (
	"fmt"
	"strings"
	"github.com/jinzhu/gorm"
)


//go:generate mockery -name=CommentService -outpkg=mocks
type CommentService interface {
	Get(id uint64) (Comment, error)
	Create(comment Comment) (Comment, error)
	Search(keywords string) ([]Comment, error)
	Update(comment Comment) (Comment, error)
	Finalize()
}

type Comment struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	gorm.Model
}

func (c *Comment) MatchKeywords(keywords []string) bool {
	return strings.Contains(c.Author, keywords[0]) || strings.Contains(c.Text, keywords[0]) //TODO: ghetto!
}

func (c *Comment) String() string {
	return fmt.Sprintf("Comment[%d,%s,%s]",c.ID,c.Author,c.Text)
}

package mem

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"

	"github.com/misoul/yellowpage/dal"
)

const commentsFile = "./server/data/comments.json"

type CommentInMem struct {
	comments []dal.Comment
}

func InitComment() (*CommentInMem, error) {
	_, err := os.Stat(commentsFile)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to stat the data file (%s): %s", commentsFile, err)
		log.Fatal(errMsg, http.StatusInternalServerError)
		return nil, err
	}

	data, err := ioutil.ReadFile(commentsFile) //TODO: not load file for each REST request
	if err != nil {
		errMsg := fmt.Sprintf("Unable to read the data file (%s): %s", commentsFile, err)
		log.Fatal(errMsg, http.StatusInternalServerError)
		return nil, err
	}

	var comments []dal.Comment
	if err := json.Unmarshal(data, &comments); err != nil {
		errMsg := fmt.Sprintf("Unable to Unmarshal comments from data file (%s): %s", commentsFile, err)
		log.Fatal(errMsg, http.StatusInternalServerError)
		return nil, err
	}

	return &CommentInMem{comments:comments}, nil
}

func (cin CommentInMem) Finalize() {
	log.Println("Closing up CommentInMem: ", cin)
}

func (cin CommentInMem) Get(id uint64) dal.Comment {
	return dal.Comment{} //TODO
}

func (cin CommentInMem) Update(comment dal.Comment) dal.Comment {
	cin.comments = append(cin.comments, comment)

	return dal.Comment{} //TODO
}

func (cin CommentInMem) Search(keywords []string) []dal.Comment {
	result := cin.comments
	if keywords != nil {
		result = dal.FilterComment(cin.comments, func(c dal.Comment) bool {
			return c.MatchKeywords(keywords)
		})
	}
	return result
}

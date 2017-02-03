package mem

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"

	"github.com/misoul/yellowpage/dal"
	"sync"
)

const commentsFile = "./server/data/comments.json"

type CommentInMem struct {
	comments []dal.Comment
	fileMutex *sync.Mutex
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

	return &CommentInMem{comments:comments, fileMutex: new(sync.Mutex)}, nil
}

func (cin CommentInMem) Finalize() {
	log.Println("Closing up CommentInMem: ", cin)
}

func (cin CommentInMem) Get(id uint64) dal.Comment {
	return dal.Comment{} //TODO
}

func (cin *CommentInMem) Update(comment dal.Comment) (dal.Comment, error) {
	cin.comments = append(cin.comments, comment)

	// Marshal the comments to indented json.
	commentData, err := json.MarshalIndent(cin.comments, "", "    ")
	if err != nil {
		log.Printf("Unable to marshal comments to json: %s", err)
		return dal.Comment{}, err
	}

	cin.fileMutex.Lock()
	defer cin.fileMutex.Unlock()

	fi, err := os.Stat(commentsFile)
	if err != nil {
		log.Printf("Unable to stat the data file (%s): %s", commentsFile, err)
		return dal.Comment{}, err
	}

	// Write out the comments to the file, preserving permissions
	err = ioutil.WriteFile(commentsFile, commentData, fi.Mode())
	if err != nil {
		log.Printf("Unable to write comments to data file (%s): %s", commentsFile, err)
		return dal.Comment{}, err
	}

	return dal.Comment{}, nil
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

package mysql

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/misoul/yellowpage/dal"
)

type CommentMySql struct {
	db *gorm.DB
}

func InitComment(dbUrl string) (*CommentMySql, error) {
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Fatalf("Failed to create SQL.DB: %s", err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	db.LogMode(true)
	db.AutoMigrate(&dal.Comment{})

	return &CommentMySql{db: db}, err
}

func (cin CommentMySql) Finalize() {
	log.Println("Closing up CommentMySql: ", cin)
	cin.db.Close()
}

func (cin CommentMySql) Get(id uint64) (dal.Comment, error) {
	var comment dal.Comment
	err := cin.db.First(&comment, id).Error
	if err != nil {
		log.Printf("Failed to get comment [%d]\n", id)
		log.Println(err)
	}
	return comment, err
}

func (cin CommentMySql) Create(comment dal.Comment) (dal.Comment, error) {
	err := cin.db.Create(&comment).Error
	if err != nil {
		log.Printf("Failed to create comment [%s]\n", comment)
		log.Println(err)
	}
	return comment, err
}

func (cin CommentMySql) Update(comment dal.Comment) (dal.Comment, error) {
	err := cin.db.Save(&comment).Error
	if err != nil {
		log.Printf("Failed to update company [%s]\n", comment)
		log.Println(err)
	}
	return comment, err
}

func (cin CommentMySql) Search(keywords []string) ([]dal.Comment, error) {
	var comments []dal.Comment
	if len(keywords) > 0 && len(keywords[0]) > 0 {
		search := "%" + keywords[0] + "%"
		err := cin.db.Where("`author` LIKE ? OR `text` LIKE ?", search, search).Find(&comments).Error
		if err != nil {
			log.Printf("Failed to find %s: %s", keywords[0], err)
			return nil, err
		}
	} else {
		comments, _ = getAllComments(cin.db)
	}

	return comments, nil
}

func getAllComments(db *gorm.DB) ([]dal.Comment, error) {
	var comments []dal.Comment
	err := db.Find(&comments).Error
	if err != nil {
		log.Printf("Failed to query all: %s", err)
		return nil, err
	}
	return comments, nil
}

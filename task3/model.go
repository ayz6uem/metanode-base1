package task3

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Model() {
	db, err := gorm.Open(sqlite.Open("base1.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&BlogUser{})
	db.AutoMigrate(&BlogPost{})
	db.AutoMigrate(&BlogComment{})

	//user := &BlogUser{Name: "jack",
	//	BlogPosts: []BlogPost{{Title: "golang hello world", Body: "hello world go", BlogComments: []BlogComment{{Content: "very good"}}},
	//		{Title: "sql rumen", Body: "select * from user"}},
	//}
	//db.Create(user)

	user := &BlogUser{}
	result := db.Model(user).Preload("BlogPosts").Preload("BlogPosts.BlogComments").First(user, 1)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("user name:%v\n", user.Name)
	for i := range user.BlogPosts {
		fmt.Println("posts:", user.BlogPosts[i].Title, user.BlogPosts[i].Body)
		for c := range user.BlogPosts[i].BlogComments {
			fmt.Println("comments:", user.BlogPosts[i].BlogComments[c].Content)
			//db.Delete(&user.BlogPosts[i].BlogComments[c])
		}
	}
}

type BlogUser struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	BlogPosts []BlogPost
	PostCount int64 `gorm:"not null"`
}

type BlogPost struct {
	gorm.Model
	Title        string `gorm:"type:varchar(255);not null"`
	Body         string `gorm:"type:varchar(255);not null"`
	BlogUserID   uint   `gorm:"not null"`
	BlogComments []BlogComment
	CommentState string `gorm:"type:varchar(255);not null;default:'无评论'"`
}

type BlogComment struct {
	gorm.Model
	Content    string `gorm:"type:varchar(255);not null"`
	BlogPostID uint   `gorm:"not null"`
}

func (p *BlogPost) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&BlogUser{}).Where("id=?", p.BlogUserID).Update("PostCount", gorm.Expr("post_count + 1"))
	return nil
}

func (c *BlogComment) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&BlogPost{}).Where("id=?", c.BlogPostID).Update("CommentState", "已评论")
	return nil
}

func (c *BlogComment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&BlogComment{}).Where("blog_post_id=?", c.BlogPostID).Count(&count)
	if count == 0 {
		tx.Model(&BlogPost{}).Where("id=?", c.BlogPostID).Update("CommentState", "无评论")
	}
	return nil
}

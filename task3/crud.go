package task3

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Crud() {
	db, err := gorm.Open(sqlite.Open("base1.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Student{})

	//student := &Student{
	//	Name:  "张三",
	//	Age:   20,
	//	Grade: "三年级",
	//}
	//db.Create(student)

	//var students []Student
	//db.Where("age > ?", 18).Find(&students)
	//fmt.Println("students:", students)

	//db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	//db.Where("age < ?", 15).Delete(&Student{})

}

type Student struct {
	ID    uint
	Name  string
	Age   uint8
	Grade string
}

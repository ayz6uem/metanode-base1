package task3

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TransactionRun() {
	db, err := gorm.Open(sqlite.Open("base1.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = db.Debug()

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	//db.Create(&Account{ID: 1, Balance: 100})
	//db.Create(&Account{ID: 2, Balance: 0})

	err = db.Transaction(func(tx *gorm.DB) error {
		a := &Account{ID: 1}
		b := &Account{ID: 2}

		t := Transaction{
			FromAccountId: a.ID,
			ToAccountId:   b.ID,
			Amount:        100,
		}

		db.Model(&a).Clauses(clause.Locking{Strength: "UPDATE"}).First(&a)
		db.Model(&b).Clauses(clause.Locking{Strength: "UPDATE"}).First(&b)
		if a.Balance < t.Amount {
			return errors.New("not enough balance")
		}
		a.Balance -= t.Amount
		b.Balance += t.Amount
		fmt.Println(a, b)
		db.Save(&a)
		db.Save(&b)
		db.Save(&t)
		return nil
	})
	if err != nil {
		panic(err)
	}

}

type Account struct {
	ID      int64
	Balance int64
}

type Transaction struct {
	ID            int64
	FromAccountId int64
	ToAccountId   int64
	Amount        int64
}

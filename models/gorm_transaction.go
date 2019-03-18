package models

import (
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/jinzhu/gorm"
)

func GormTransaction(db *gorm.DB, txFunc func(*gorm.DB) *sample_error.SampleError) (err *sample_error.SampleError) {
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = txFunc(tx)
	return err
}

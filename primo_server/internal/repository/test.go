package repository

import "gorm.io/gorm"

type TestRepository interface {
	// not use
}
type testRepository struct {
	db *gorm.DB
}

func NewTestRepository(db *gorm.DB) TestRepository {
	return &testRepository{db: db}
}

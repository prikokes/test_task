package models

import "github.com/jinzhu/gorm"

type DB interface {
	Where(query interface{}, args ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Begin() *gorm.DB
	Table(name string) *gorm.DB
	Count(value interface{}) *gorm.DB
	Update(attrs ...interface{}) *gorm.DB
	Commit() *gorm.DB
	Rollback() *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Joins(query string, args ...interface{}) *gorm.DB
	Scan(dest interface{}) *gorm.DB
}

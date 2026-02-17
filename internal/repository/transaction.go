package repository

import "gorm.io/gorm"

// WithTransaction runs fn inside a DB transaction.
// Commit on success, rollback on any error or panic.
func WithTransaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	return db.Transaction(fn)
}

package repository

import (
	"strings"
)

// Checks for PostgreSQL unique constraint violation (code 23505)
func isDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "23505") ||
		strings.Contains(err.Error(), "duplicate key")
}

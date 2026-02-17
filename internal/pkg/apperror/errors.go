package apperror

import "errors"

var (
	ErrNotFound          = errors.New("data not found")
	ErrDuplicateEntry    = errors.New("data already exists")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrAlreadyRated      = errors.New("gift already rated")
	ErrNotRedeemed       = errors.New("gift has not been redeemed by this user")
)

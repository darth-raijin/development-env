package internal

import "github.com/google/uuid"

type User struct {
	id      uuid.UUID
	balance float32
}

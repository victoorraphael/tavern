package tavern

import (
	"github.com/google/uuid"
)

//Person represent a person
type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}

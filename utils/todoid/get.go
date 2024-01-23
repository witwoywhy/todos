package todoid

import "github.com/google/uuid"

type GetIDFunc func() string

var GetID GetIDFunc = func() string {
	return uuid.New().String()
}

package domain

import (
	"fmt"
	"strings"
)


type NotFoundError struct {
	MissingItemIDs []string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("these items not found in stock %s", strings.Join(e.MissingItemIDs, ", "))	
}
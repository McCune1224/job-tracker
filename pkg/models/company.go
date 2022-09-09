package models

import (
	"github.com/google/uuid"
)


type Company struct {
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
    Office string `json:"office"`

}




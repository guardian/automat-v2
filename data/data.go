package data

import "time"

type Slot struct {
	ID, Description string
}

var Slots = []Slot{
	{"body-end", "Slot following article body."},
}

func FindSlot(ID string) (Slot, bool) {
	for _, s := range Slots {
		if s.ID == ID {
			return s, true
		}
	}

	return Slot{}, false
}

type Owner struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Test struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsEnabled   bool      `json:"isEnabled"`
	Variants    []string  `json:"variants"`
	Owner       Owner     `json:"author"`
	Expiry      time.Time `json:"expiry"`
}

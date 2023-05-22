package go_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Yoga",
		MiddleName: "Meleniawan",
		LastName:   "Pamungkas",
	}

	encoder.Encode(customer)
}

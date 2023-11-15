package functional_options

import (
	"log"
	"testing"
)

func TestNewUser(t *testing.T) {
	u := NewUser(WithName("TOM"), WithAge(18), WithHobbies([]string{"bike packing", "walk off road"}))

	log.Printf("u:%#v\n", u)
}

func TestNewUser2(t *testing.T) {
	u := NewUser2(&NameAgeOption{Age: 18, Name: "Jack"}, &HobbiesOption{Hobbies: []string{"A", "B"}})
	log.Printf("u:%#v\n", u)
}

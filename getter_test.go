package main

import (
	"fmt"
	"testing"
)

func TestKadaiGetter_GetAll(t *testing.T) {
	getter := NewKadaiGetter(STUDENT_ID, PASSWORD)

	kadais, err := getter.GetAll()
	if err != nil {
		t.Fatalf("Failed to get kadais: %v", err)
	}

	for _, k := range kadais {
		fmt.Printf("Type: '%s'\n", k._type)
		fmt.Printf("Title: '%s'\n", k.title)
		fmt.Printf("Course: '%s'\n", k.course)
		if k.start != nil {
			fmt.Printf("Start: '%s'\n", k.start)
		} else {
			fmt.Println("Start: 'なし'")
		}
		if k.deadline != nil {
			fmt.Printf("Deadline: '%s'\n", k.deadline)
		} else {
			fmt.Println("Deadline: 'なし'")
		}
		fmt.Println("-------------------------")
	}
}

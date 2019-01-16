package usecase

import (
	"fmt"
	"testing"
)

func TestGetWord(t *testing.T) {
	var v WordUseCase
	fmt.Println(v.GetWord())
	// fmt.Println(GetWord())
}

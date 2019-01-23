package usecase

import (
	"fmt"
	"testing"
)

func TestGetWord(t *testing.T) {
	v := wordUseCase{}
	fmt.Println(v.GetWord(1))
}

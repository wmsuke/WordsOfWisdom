package repositores

import (
	"fmt"
	"testing"
)

func TestFindOne(t *testing.T) {
	var s = NewWordRepository()
	result, error := s.FindOne(1)
	if error != nil {
		t.Fatalf("failed test %#v", error)
	}
	fmt.Println(result)
}

func TestRandomOne(t *testing.T) {
	var v = NewWordRepository()
	fmt.Println(v.RandomOne())
}

func TestGetNiceCount(t *testing.T) {
	fmt.Println(getNiceCount(1))
}

func TestGetFavoriteCount(t *testing.T) {
	fmt.Println(getFavoriteCount(1))
}

func TestIsFavorite(t *testing.T) {
	fmt.Println(isFavorite(1, "abc1"))
}

func TestIsNicd(t *testing.T) {
	fmt.Println(isNice(1, "abc1"))
}

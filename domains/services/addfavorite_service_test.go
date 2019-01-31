package services

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	ctrl := NewAddFavoriteServices()
	result := ctrl.Add(1, "dVachSk3sK1NpN3dri0wp0KOh1gqsSZMODrqo7Zv3WF3kzgXWCAfMCwIN6gYCdHI")
	fmt.Println(result)
}

package rokuremote

import (
	"os"
	"testing"
)

func TestPlayer_Home(t *testing.T) {
	r, _ := Connect(os.Getenv("ROKUIP"))
	err := r.Home()
	if err != nil {
		t.Errorf("Home: %s\n", err)
	}
}

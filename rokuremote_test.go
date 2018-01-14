package rokuremote



import (
	"testing"
	"os"
)


func TestPlayer_Home(t *testing.T) {
	r, _ := Connect(os.Getenv("ROKUIP"))
	r.Home()
}

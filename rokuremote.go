package rokuremote

import (
	"net/http"
	"strconv"
)

type Player struct {
	Address string
	Port string
	NickName string
	Client *http.Client
}

const (
	Hulu    = 2285
	Netflix = 12
	HBO = 61322
	Plex = 13535
	AmazonPrime = 13
	USTVNow = 2026
)

func Connect(ip string) (Player, error) {
	return Player{ip, "8060", "", &http.Client{}}, nil
}

// Given player nickname connect and return player
func ConnectName(name string) (Player, error) {

	return Player{}, nil
}

func (p *Player) Post(path string) error {
	req, err := http.NewRequest("POST", "http://"+p.Address+":8060"+path, nil)
	if err != nil {
		return err
	}
	_, err =  p.Client.Do(req)
	return err
}

func (p *Player) Home() error {
	return p.Post("/keypress/home")
}
func (p *Player) Up() error {
	return p.Post("/keypress/up")
}
func (p *Player) Down() error {
	return p.Post("/keypress/down")
}
func (p *Player) Left() error {
	return p.Post("/keypress/left")
}
func (p *Player) Right() error {
	return p.Post("/keypress/right")
}
func (p *Player) OK() error {
	return p.Post("/keypress/select")
}
func (p *Player) Back() error {
	return p.Post("/keypress/back")
}

func (p *Player) StartChannel(chanid int) error {
	return p.Post("/launch/" + strconv.Itoa(chanid))
}
package main

import (
	"errors"
	"fmt"
	"log"
)

type Player struct {
	Name       string
	Games      []string
	Age        int
	Height     int
	Weight     int
	Achievment []string
}

type PlayerBuilder interface {
	SetName(string) PlayerBuilder
	AddGames([]string) PlayerBuilder
	SetAge(int) PlayerBuilder
	AddHeight(int) PlayerBuilder
	AddWeight(int) PlayerBuilder
	AddAchievment([]string) PlayerBuilder
	Build() (*Player, error)
}

type playerBuilder struct {
	player Player
}

func NewPlayerBuilder() *playerBuilder {
	return &playerBuilder{
		player: Player{},
	}
}

func (p *playerBuilder) SetName(name string) PlayerBuilder {
	p.player.Name = name
	return p
}

func (p *playerBuilder) AddGames(games []string) PlayerBuilder {
	p.player.Games = games
	return p
}

func (p *playerBuilder) SetAge(age int) PlayerBuilder {
	p.player.Age = age
	return p
}

func (p *playerBuilder) AddHeight(h int) PlayerBuilder {
	p.player.Height = h
	return p
}

func (p *playerBuilder) AddWeight(w int) PlayerBuilder {
	p.player.Weight = w
	return p
}

func (p *playerBuilder) AddAchievment(ach []string) PlayerBuilder {
	p.player.Achievment = ach
	return p
}

func (p *playerBuilder) Build() (*Player, error) {
	if p.player.Weight < 50 {
		return nil, errors.New("player under Weight")
	}
	return &p.player, nil
}

func main() {
	player, err := NewPlayerBuilder().SetName("junaid").SetAge(18).AddGames([]string{"cricket", "football"}).
		AddAchievment([]string{"state player"}).AddHeight(160).AddWeight(60).Build()
	if err != nil {
		log.Print(err)
	}

	fmt.Println(player)

}

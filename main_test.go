package main

import (
	"awesomeProject/pkg/PFC"
	"testing"
)

func TestPFCIsWorking(t *testing.T) {
	player := PFC.NewPlayer()
	game := PFC.NewGame()
	game.Play(player.AiPlay())
}

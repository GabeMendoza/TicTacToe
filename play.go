package main

import (
  "fmt"
  "./ttt"
)

func NewGame() ttt.Game {
  players := InitPlayers()
  board   := ttt.MakeBoard()

  return ttt.Game{board, players}
}

func InitPlayers() [2]ttt.Player {
  var players [2]ttt.Player
  marks := [2]byte{'X', 'O'}

  for i := 0; i < 2; i++ {
    fmt.Println("Name of player", string(marks[i]))
    name := ttt.PromptGetInput()
    players[i] = ttt.Player{name, marks[i]}
  }

  return players
}

func PlayGame() bool {
  game := NewGame()

  for PlayRound(game.Board, game.Players) {}

  if game.Board.IsTie() {
    fmt.Println("Oops, it seems that there is no winner...")
  }
  if game.Board.IsWon() {
    for _, player := range game.Players {
      if game.Board.IsWinner(player) {
        fmt.Printf("%s wins!\n", player.Name)
        break
      }
    }
  }

  return ttt.PromptPlayAgain()
}

func PlayRound(board ttt.Board, players [2]ttt.Player) bool {
  for _, player := range players {
    ttt.PrintBoard(board)
    fmt.Printf("%s's turn (%s) \n", player.Name, string(player.Symbol))

    cell := ttt.PromptGetPlayableCell(board)
    board.PlaceMove(cell, player)

    fmt.Printf("\n\n")

    if board.IsOver() {
      return false
    }
  }
  return true
}

func main() {
  for PlayGame() {}
}

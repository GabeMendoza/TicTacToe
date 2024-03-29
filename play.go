package main

import "./ttt"

func NewGame() ttt.Game {
  players := InitPlayers()
  board   := ttt.MakeBoard()

  return ttt.Game{board, players}
}

func InitPlayers() []ttt.Player {
  player1 := ttt.Player{ttt.PromptGetName("X", 34), 'X', 34}
  player2 := ttt.Player{ttt.PromptGetName("O", 36), 'O', 36}

  return []ttt.Player{player1, player2}
}

func PlayNewGame() bool {
  game := NewGame()

  PlayGame(game)
  EndGame(game)

  return ttt.PromptPlayAgain()
}

func PlayGame(game ttt.Game) {
  player := game.FirstPlayer()

  ttt.PrintBoard(game.Board, game.Players)
  ttt.PrintBreak()

  for !game.Board.IsOver() {
    cell := ttt.PromptGetCell(game.Board, player)
    game.Board.PlaceMove(cell, player)

    ttt.PrintBoard(game.Board, game.Players)
    ttt.PrintBreak()
    player = game.NextPlayer(player)
  }
}

func EndGame(game ttt.Game) {
 if game.Board.IsTie() {
    ttt.PrintTie()
  } else if game.Board.IsWon() {
    ttt.PrintWin(*game.Winner())
  }
}

func main() {
  for PlayNewGame() {}
}

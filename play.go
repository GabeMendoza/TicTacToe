package main

import "./ttt"

func NewGame() ttt.Game {
  players := InitPlayers()
  board   := ttt.MakeBoard()

  return ttt.Game{board, players}
}

func InitPlayers() []ttt.Player {
  player1 := ttt.Player{ttt.PromptGetName("X"), 'X'}
  player2 := ttt.Player{ttt.PromptGetName("O"), 'O'}

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

  for !game.Board.IsOver() {
    ttt.PrintBoard(game.Board)
    ttt.PrintPlayerTurn(player)

    cell := ttt.PromptGetPlayableCell(game.Board)
    game.Board.PlaceMove(cell, player)

    ttt.PrintTurnBreak()
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

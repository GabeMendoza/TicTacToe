package main

import (
  "fmt"
  "strconv"
  "strings"
  "./ttt"
)

func PrintBoard(board ttt.Board) {
  for l := 0; l < 3 ; l++ {
    line := []string{"", "", ""}

    for c := 0 ; c < 3 ; c++ {
      index := l * 3 + c
      if board.Cells[index] == 0 {
        line[c] = strconv.Itoa(index + 1)
      } else {
        line[c] = string(board.Cells[index])
      }
    }
    fmt.Println(strings.Join(line, " | "))
  }
}

func PrintPrompt() {
  fmt.Print("-> ")
}

func GetInput() string {
  var input string
  fmt.Scan(&input)
  return input
}

func PlayRound(board ttt.Board, players []ttt.Player) {
  for _, player := range players {
    PrintBoard(board)
    fmt.Printf("%s's turn (%s) \n", player.Name, string(player.Symbol))
    for {
      PrintPrompt()
      input := GetInput()
      cell, err := strconv.Atoi(input)
      if err == nil && board.IsCellPlayable(cell) {
        index := cell - 1
        if board.IsCellPlayable(index) {
          board.PlaceMove(index, player)
          break
        }
      }
    }

    fmt.Printf("\n\n")

    if board.IsOver() {
      break
    }
  }
}

func main() {
  player1 := ttt.Player{"John", 'X'}
  player2 := ttt.Player{"Bruce", 'O'}
  players := []ttt.Player{player1, player2}

  board := ttt.MakeBoard()

  for {
    PlayRound(board, players)
    if board.IsOver() {
      break
    }
  }

  if board.IsTie() {
    fmt.Println("No more moves possible. Let's play again!")
  }
  if board.IsWon() {
    for _, player := range players {
      if board.IsWinner(player) {
        fmt.Printf("%s wins!\n", player.Name)
        break
      }
    }
  }
}

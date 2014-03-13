package ttt

import (
  "fmt"
  "strings"
  "strconv"
)

func PrintBoard(board Board) {
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

func PromptGetInput() string {
  var input string
  fmt.Print("-> ")
  fmt.Scan(&input)
  return input
}

func PromptGetInt() int {
  input := PromptGetInput()
  val, err := strconv.Atoi(input)
  if err != nil {
    val = PromptGetInt()
  }
  return val
}

func PromptGetBool() bool {
  input := PromptGetInput()
  return input == "y" || input == "Y" || input == "YES" || input == "yes"
}

func PromptGetPlayableCell(board Board) int {
  cell := PromptGetInt() - 1
  if !board.IsCellPlayable(cell) {
    cell = PromptGetPlayableCell(board)
  }
  return cell
}

func PromptPlayAgain() bool {
  fmt.Println("Play again? (y/n)")
  return ttt.PromptGetBool()
}

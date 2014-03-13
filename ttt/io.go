package ttt

import (
  "fmt"
  "strings"
  "strconv"
)

func color(c int) string {
  return "\x1b[" + strconv.Itoa(c) + ";1m"
}

func colorReset() string {
  return "\x1b[0m"
}

func PrintBoard(board Board, players []Player) {
  for l := 0; l < 3 ; l++ {
    line := []string{"", "", ""}

    for c := 0 ; c < 3 ; c++ {
      index := l * 3 + c
      if board.Cells[index] == 0 {
        line[c] = strconv.Itoa(index + 1)
      } else {
        for _, player := range players {
          if player.Symbol == board.Cells[index] {
            line[c] = color(player.Color) + string(player.Symbol) + colorReset()
            break
          }
        }
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
  return PromptGetBool()
}

func PrintTie() {
  fmt.Println("Oops, it seems that there is no winner...")
}

func PrintWin(winner Player) {
  fmt.Printf("%s%s%s wins!\n", color(winner.Color), winner.Name, colorReset())
}

func PrintPlayerTurn(player Player) {
  fmt.Printf("%s%s%s's turn (%s%s%s) \n", color(player.Color), player.Name, colorReset(), color(player.Color), string(player.Symbol), colorReset())
}

func PrintTurnBreak() {
  fmt.Printf("\n\n")
}

func PromptGetName(mark string, c int) string {
  fmt.Printf("Name of %splayer %s%s:\n", color(c), mark, colorReset())
  return PromptGetInput()
}

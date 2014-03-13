package ttt

type Board struct{
  Cells []byte
}

func MakeBoard() Board {
  return Board{make([]byte, 9)}
}

func (b *Board) PlaceMove(cell int, player Player) {
  b.Cells[cell] = player.Symbol
}

func (b Board) WinningAligns() [][]int {
  return [][]int {
    {0, 1, 2},
    {3, 4, 5},
    {6, 7, 8},
    {0, 3, 6},
    {1, 4, 7},
    {2, 5, 8},
    {0, 4, 8},
    {2, 4, 6},
  }
}

func (b Board) IsWinningAlign(align []int) bool {
  firstCellEmpty := b.Cells[align[0]] == 0
  first2Equals   := b.Cells[align[0]] == b.Cells[align[1]]
  last2Equals    := b.Cells[align[1]] == b.Cells[align[2]]

  if !firstCellEmpty && first2Equals && last2Equals {
    return true
  }
  return false
}

func (b Board) IsWon() bool {
  for _, align := range b.WinningAligns() {
    if b.IsWinningAlign(align) {
      return true
    }
  }

  return false
}

func (b Board) IsOver() bool {
  return b.IsWon() || b.IsTie()
}

func (b Board) HasEmptyCells() bool {
  for _, cell := range b.Cells {
    if cell == 0 {
      return true
    }
  }
  return false
}

func (b Board) IsTie() bool {
  if !b.HasEmptyCells() && !b.IsWon() {
    return true
  }
  return false
}

func (b Board) IsWinner(player Player) bool {
  if !b.IsWon() {
    return false
  }

  for _, align := range b.WinningAligns() {
    if player.Symbol == b.Cells[align[0]] && b.IsWinningAlign(align) {
      return true
    }
  }

  return false
}

func (b Board) IsCellPlayable(cell int) bool {
  if b.Cells[cell] == 0 {
    return true
  }
  return false
}

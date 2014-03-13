package ttt

type Game struct{
  Board   Board
  Players []Player
}

func (game *Game) FirstPlayer() Player {
  return game.Players[0]
}

func (game *Game) NextPlayer(current Player) Player {
  if current == game.Players[0] {
    return game.Players[1]
  } else {
    return game.Players[0]
  }
}

func (game *Game) Winner() *Player {
  for _, player := range game.Players {
    if game.Board.IsWinner(player) {
      return &player
    }
  }
  return nil
}

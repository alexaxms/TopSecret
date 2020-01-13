package domain

type Board struct {
	Name    string
	Active  bool
}

type BoardUsecase interface {
	GetBoards(b []*Board) ([]*Board, error)
}

type BoardRepository interface {
	Boards(u []*Board) ([]*Board, error)
}

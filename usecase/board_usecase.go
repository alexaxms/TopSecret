package usecase

import (
	"github.com/alexaxms/TopSecret/domain"
)

type BoardUsecase struct {
	boardRepo domain.BoardRepository
}

func NewBoardUsecase(repo domain.BoardRepository) *BoardUsecase {
	return &BoardUsecase{
		boardRepo: repo,
	}
}

func (br *BoardUsecase) GetBoards(b []*domain.Board) ([]*domain.Board, error) {
	b, err := br.boardRepo.Boards(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

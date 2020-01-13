package controller

import (
	"github.com/alexaxms/TopSecret/domain"
	"net/http"
)

type boardController struct {
	boardUsecase domain.BoardUsecase
}

type BoardController interface {
	GetBoards(c Context) error
}

func NewBoardController(us domain.BoardUsecase) BoardController {
	return &boardController{us}
}

func (uc *boardController) GetBoards(c Context) error {
	var b []*domain.Board

	b, err := uc.boardUsecase.GetBoards(b)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, b)
}

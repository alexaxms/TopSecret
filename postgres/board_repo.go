package postgres

import (
	"github.com/alexaxms/TopSecret/domain"
	"github.com/jinzhu/gorm"
)

type boardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) domain.BoardRepository {
	return &boardRepository{db}
}
func (bs boardRepository) Boards(u []*domain.Board) ([]*domain.Board, error) {
	//var boards = []domain.Board{{
	//	Name:   "x",
	//	Active: false,
	//}, {
	//	Name:   "d",
	//	Active: false,
	//}}
	err := bs.db.Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

package registry

import (
	"github.com/alexaxms/TopSecret/controller"
	d "github.com/alexaxms/TopSecret/domain"
	br "github.com/alexaxms/TopSecret/postgres"
	bu "github.com/alexaxms/TopSecret/usecase"
)

func (r *registry) NewBoardController() controller.BoardController {
	return controller.NewBoardController(r.NewBoardUseCase())
}

func (r *registry) NewBoardUseCase() d.BoardUsecase {
	return bu.NewBoardUsecase(r.NewBoardRepository())
}

func (r *registry) NewBoardRepository() d.BoardRepository {
	return br.NewBoardRepository(r.db)
}

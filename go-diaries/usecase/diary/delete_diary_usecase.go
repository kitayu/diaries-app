package diary

import (
	"context"

	"github.com/kitayu/go-diaries/usecase/repository"
)

type DeleteDiaryInputPort struct {
	ID int64
}

type DeleteDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

func NewDeleteDiaryUsecase(dr repository.DiaryRepository) *DeleteDiaryUsecase {
	return &DeleteDiaryUsecase{dr}
}

func (du DeleteDiaryUsecase) Execute(ctx context.Context, in *DeleteDiaryInputPort) error {

	if err := du.diaryRepo.Delete(ctx, in.ID); err != nil {
		return err
	}

	return nil
}

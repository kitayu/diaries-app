package diary

import (
	"context"

	"github.com/kitayu/go-diaries/domain/model"
	"github.com/kitayu/go-diaries/usecase/repository"
)

type GetDiaryInputPort struct {
	ID int64
}

type GetDiaryOutputPort struct {
	Diary *model.Diary
}

type GetDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

func NewGetDiaryUsecase(dr repository.DiaryRepository) *GetDiaryUsecase {
	return &GetDiaryUsecase{dr}
}

func (du GetDiaryUsecase) Execute(ctx context.Context, in *GetDiaryInputPort) (*GetDiaryOutputPort, error) {
	diary, err := du.diaryRepo.FindByID(ctx, in.ID)

	if err != nil {
		return nil, err
	}

	return &GetDiaryOutputPort{diary}, nil
}

package diary

import (
	"context"

	"github.com/kitayu/go-diaries/domain/model"
	"github.com/kitayu/go-diaries/usecase/repository"
)

type CreateDiaryInputPort struct {
	Title       string
	Description string
}

type CreateDiaryOutputPort struct {
	Diary *model.Diary
}

type CreateDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

func NewCreateDiaryUsecase(dr repository.DiaryRepository) *CreateDiaryUsecase {
	return &CreateDiaryUsecase{dr}
}

func (du CreateDiaryUsecase) Execute(ctx context.Context, in *CreateDiaryInputPort) (*CreateDiaryOutputPort, error) {
	diary := &model.Diary{
		Title:       in.Title,
		Description: in.Description,
	}

	diary, err := du.diaryRepo.Store(ctx, diary)
	if err != nil {
		return nil, err
	}

	return &CreateDiaryOutputPort{diary}, nil
}

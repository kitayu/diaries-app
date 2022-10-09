package diary

import (
	"context"

	"github.com/kitayu/go-diaries/domain/model"
	"github.com/kitayu/go-diaries/usecase/repository"
)

type ListDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

type ListDiaryOutputPort struct {
	Diaries []*model.Diary
}

func NewListDiaryUsecase(dr repository.DiaryRepository) *ListDiaryUsecase {
	return &ListDiaryUsecase{dr}
}

func (du ListDiaryUsecase) Execute(ctx context.Context) (*ListDiaryOutputPort, error) {
	diaries, err := du.diaryRepo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return &ListDiaryOutputPort{diaries}, nil
}

package diary

import (
	"context"

	"github.com/kitayu/go-diaries/domain/model"
	"github.com/kitayu/go-diaries/usecase/repository"
)

type UpdateDiaryInputPort struct {
	ID          int64
	Title       string
	Description string
}

type UpdateDiaryOutputPort struct {
	Diary *model.Diary
}

type UpdateDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

func NewUpdateDiaryUsecase(dr repository.DiaryRepository) *UpdateDiaryUsecase {
	return &UpdateDiaryUsecase{dr}
}

func (du UpdateDiaryUsecase) Execute(ctx context.Context, in *UpdateDiaryInputPort) (*UpdateDiaryOutputPort, error) {
	diary := &model.Diary{
		ID:          in.ID,
		Title:       in.Title,
		Description: in.Description,
	}

	diary, err := du.diaryRepo.Update(ctx, diary)
	if err != nil {
		return nil, err
	}

	return &UpdateDiaryOutputPort{diary}, nil
}

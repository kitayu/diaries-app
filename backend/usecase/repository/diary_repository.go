package repository

import (
	"context"

	"github.com/kitayu/go-diaries/domain/model"
)

type DiaryRepository interface {
	Store(ctx context.Context, diary *model.Diary) (*model.Diary, error)
	Update(ctx context.Context, diary *model.Diary) (*model.Diary, error)
	Delete(ctx context.Context, id int64) error
	FindAll(ctx context.Context) ([]*model.Diary, error)
	FindByID(ctx context.Context, id int64) (*model.Diary, error)
}

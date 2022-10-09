package repository

import (
	"context"
	"database/sql"

	"github.com/kitayu/go-diaries/domain/model"
)

type diaryRepository struct {
	db *sql.DB
}

func NewDiaryRepository(db *sql.DB) *diaryRepository {
	return &diaryRepository{db}
}

func (dr *diaryRepository) Store(ctx context.Context, diary *model.Diary) (*model.Diary, error) {
	stmt, err := dr.db.PrepareContext(ctx, "INSERT INTO diaries(title, description) VALUES (?, ?);")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	r, err := stmt.ExecContext(ctx, diary.Title, diary.Description)
	if err != nil {
		return nil, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}

	diary.ID = id
	return diary, nil
}

func (dr *diaryRepository) Update(ctx context.Context, diary *model.Diary) (*model.Diary, error) {
	stmt, err := dr.db.PrepareContext(ctx, "UPDATE diaries SET title = ?, description = ? WHERE id = ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(diary.Title, diary.Description, diary.ID)

	if err != nil {
		return nil, err
	}

	return diary, nil
}

func (dr *diaryRepository) Delete(ctx context.Context, id int64) error {
	stmt, err := dr.db.PrepareContext(ctx, "DELETE FROM diaries WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (dr *diaryRepository) FindAll(ctx context.Context) ([]*model.Diary, error) {
	stmt, err := dr.db.PrepareContext(ctx, "SELECT id, title, description FROM diaries")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows *sql.Rows
	rows, err = stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var diaries []*model.Diary
	for rows.Next() {
		var diary model.Diary
		if err := rows.Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
			return nil, err
		}
		diaries = append(diaries, &diary)
	}

	return diaries, nil
}

func (dr *diaryRepository) FindByID(ctx context.Context, id int64) (*model.Diary, error) {
	stmt, err := dr.db.PrepareContext(ctx, "SELECT id, title, description FROM diaries WHERE id = ?")
	if err != nil {
		return nil, err
	}

	diary := &model.Diary{}
	if err := stmt.QueryRow(id).Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
		return nil, err
	}

	return diary, nil
}

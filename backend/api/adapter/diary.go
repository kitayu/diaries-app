package adapter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kitayu/go-diaries/usecase/diary"
)

type createDiaryRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewCreateDiaryInputPortRequest(r *http.Request) (*diary.CreateDiaryInputPort, error) {
	var input createDiaryRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	return &diary.CreateDiaryInputPort{
		Title:       input.Title,
		Description: input.Description,
	}, nil
}

type updateDiaryRequestBody struct {
	ID          int64  `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUpdateDiaryInputPortRequest(r *http.Request) (*diary.UpdateDiaryInputPort, error) {
	var input updateDiaryRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		return nil, err
	}

	input.ID = int64(ID)

	return &diary.UpdateDiaryInputPort{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
	}, nil
}

func NewDeleteDiaryInputPortRequest(r *http.Request) (*diary.DeleteDiaryInputPort, error) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		return nil, err
	}

	return &diary.DeleteDiaryInputPort{
		ID: int64(ID),
	}, nil
}

func NewGetDiaryInputPortRequest(r *http.Request) (*diary.GetDiaryInputPort, error) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		return nil, err
	}

	return &diary.GetDiaryInputPort{
		ID: int64(ID),
	}, nil
}

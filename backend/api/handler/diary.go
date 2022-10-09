package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kitayu/go-diaries/api/adapter"
	"github.com/kitayu/go-diaries/api/presenter"
	"github.com/kitayu/go-diaries/usecase/diary"
)

func NewCreateDiaryHandler(du *diary.CreateDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adapter
		input, err := adapter.NewCreateDiaryInputPortRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Usecase
		output, err := du.Execute(r.Context(), input)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Presenter
		if err := json.NewEncoder(w).Encode(
			presenter.NewCreateDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewUpdateDiaryHandler(du *diary.UpdateDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adapter
		input, err := adapter.NewUpdateDiaryInputPortRequest(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Usecase
		output, err := du.Execute(r.Context(), input)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Presenter
		if err := json.NewEncoder(w).Encode(
			presenter.NewUpdateDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewDeleteDiaryHandler(du *diary.DeleteDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adapter
		input, err := adapter.NewDeleteDiaryInputPortRequest(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Usecase
		err = du.Execute(r.Context(), input)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Presenter
	})
}

func NewGetDiaryHandler(du *diary.GetDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adapter
		input, err := adapter.NewGetDiaryInputPortRequest(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Usecase
		output, err := du.Execute(r.Context(), input)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Presenter
		if err := json.NewEncoder(w).Encode(
			presenter.NewGetDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewListDiaryHandler(du *diary.ListDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adapter

		// Usecase
		output, err := du.Execute(r.Context())

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Presenter
		if err := json.NewEncoder(w).Encode(
			presenter.NewListDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

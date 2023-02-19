package api

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/kitayu/go-diaries/api/handler"
	"github.com/kitayu/go-diaries/interface/repository"
	"github.com/kitayu/go-diaries/usecase/diary"
)

func BuildRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	buildProjectRoutes(r, db)
	return r
}

func buildProjectRoutes(r *mux.Router, db *sql.DB) {
	dr := repository.NewDiaryRepository(db)
	r.Handle("/diary", handler.NewCreateDiaryHandler(
		diary.NewCreateDiaryUsecase(dr))).Methods("POST", "OPTIONS")
	r.Handle("/diary/{id}", handler.NewUpdateDiaryHandler(
		diary.NewUpdateDiaryUsecase(dr))).Methods("PUT", "OPTIONS")
	r.Handle("/diary/{id}/delete", handler.NewDeleteDiaryHandler(
		diary.NewDeleteDiaryUsecase(dr))).Methods("DELETE", "OPTIONS")
	r.Handle("/diary/{id}", handler.NewGetDiaryHandler(
		diary.NewGetDiaryUsecase(dr))).Methods("GET")
	r.Handle("/diaries", handler.NewListDiaryHandler(
		diary.NewListDiaryUsecase(dr))).Methods("GET")
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kitayu/go-diaries/config"
	model "github.com/kitayu/go-diaries/domain/model"
	db "github.com/kitayu/go-diaries/infrastructure/db"
	server "github.com/kitayu/go-diaries/infrastructure/server"
	repository "github.com/kitayu/go-diaries/interface/repository"
)

type DiaryRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer conn.Close()

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Handle("/diary", addDiary(conn)).Methods("POST")
	r.Handle("/diary/{id}", editDiary(conn)).Methods("PUT")
	r.Handle("/diary/{id}/delete", deleteDiary(conn)).Methods("DELETE")
	r.Handle("/diary/{id}", getDiary(conn)).Methods("GET")
	r.Handle("/diaries", getDiaryList(conn)).Methods("GET")

	srv := server.NewServer(r)

	log.Printf("Serving on localhost:%v\n", config.Config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}

func addDiary(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input DiaryRequest
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			log.Fatalf("jsonデコードに失敗しました。 %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dr := repository.NewDiaryRepository(db)
		d := &model.Diary{
			Title:       input.Title,
			Description: input.Description,
		}
		_, err := dr.Store(r.Context(), d)

		if err != nil {
			log.Fatalf("diariesの登録に失敗しました。 %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func editDiary(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Fatalf("idの取得に失敗しました。 %v", err)
			panic("パラメーターエラー")
		}

		var input DiaryRequest
		if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
			log.Fatalf("jsonデコードに失敗しました。 %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dr := repository.NewDiaryRepository(db)
		d := &model.Diary{
			ID:          int64(ID),
			Title:       input.Title,
			Description: input.Description,
		}
		_, err = dr.Update(r.Context(), d)

		if err != nil {
			log.Fatalf("diariesの更新に失敗しました。 %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func deleteDiary(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Fatalf("idの取得に失敗しました。 %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dr := repository.NewDiaryRepository(db)
		err = dr.Delete(r.Context(), int64(ID))

		if err != nil {
			log.Fatalf("diariesの削除に失敗しました。 %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func getDiary(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Fatalf("idの取得に失敗しました。 %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dr := repository.NewDiaryRepository(db)
		diary, err := dr.FindByID(r.Context(), int64(ID))

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		fmt.Fprintln(w, diary)
	})
}

func getDiaryList(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dr := repository.NewDiaryRepository(db)
		diaries, err := dr.FindAll(r.Context())

		if err != nil {
			log.Fatalf("diariesの取得に失敗しました。 %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(diaries)
		if err != nil {
			log.Fatalf("jsonデコードに失敗しました。: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if b := string(json); b == "null" {
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprintln(w, "")
		} else {
			fmt.Fprintln(w, string(json))
		}
	})
}

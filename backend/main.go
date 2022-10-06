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
	db "github.com/kitayu/go-diaries/infrastructure/db"
	server "github.com/kitayu/go-diaries/infrastructure/server"
)

type Diary struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

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

		_, err := db.Exec(
			"INSERT INTO diaries(title, description) VALUES (?, ?)",
			input.Title,
			input.Description,
		)

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

		_, err = db.Exec(
			"UPDATE diaries SET title = ?, description = ? WHERE id = ?",
			input.Title,
			input.Description,
			ID,
		)

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

		_, err = db.Exec(
			"DELETE FROM diaries WHERE id = ?",
			ID,
		)

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

		rows := db.QueryRow(
			"SELECT id, title, description FROM diaries WHERE id = ?",
			ID,
		)
		var diary Diary
		if err = rows.Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, diary)
	})
}

func getDiaryList(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, title, description FROM diaries")
		if err != nil {
			log.Fatalf("diariesの取得に失敗しました。 %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var diaries []*Diary
		for rows.Next() {
			var diary Diary
			if err := rows.Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
				log.Fatalf("diariesの取得に失敗しました。 %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			diaries = append(diaries, &diary)
		}

		json, err := json.Marshal(diaries)
		if err != nil {
			log.Fatalf("jsonデコードに失敗しました。: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, string(json))
	})
}

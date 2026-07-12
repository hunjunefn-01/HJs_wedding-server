package handler

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"hjs-wedding-server/env"
	"hjs-wedding-server/httphandler"
	"hjs-wedding-server/sqldb"
	_ "github.com/lib/pq" // Postgres 드라이버를 import 합니다.
	"github.com/rs/cors"
)

var db *sql.DB

// Handler 함수는 Vercel에 의해 호출되는 진입점입니다.
func Handler(w http.ResponseWriter, r *http.Request) {
	// 데이터베이스 연결이 없으면 새로 생성합니다.
	// Vercel의 서버리스 환경에서는 함수가 "warm" 상태일 때 이 연결이 재사용될 수 있습니다.
	if db == nil {
		log.Println("Initializing database connection...")
		connStr := os.Getenv("POSTGRES_URL") // Vercel이 주입해주는 환경 변수
		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Failed to open database connection: %v", err)
		}
		sqldb.SetDb(db)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/guestbook", new(httphandler.GuestbookHandler))
	mux.Handle("/api/attendance", new(httphandler.AttendanceHandler))

	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{env.AllowOrigin},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut},
		AllowCredentials: true,
	})

	handler := corHandler.Handler(mux)
	handler.ServeHTTP(w, r)
}
package httphandler

import (
	"encoding/json"
	"net/http"

	"hjs-wedding-server/sqldb"
	"hjs-wedding-server/types"
)

type AttendanceHandler struct {
	http.Handler
}

func (h *AttendanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var attendance types.AttendanceCreate
		err := decoder.Decode(&attendance)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("BadRequest"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		err = sqldb.CreateAttendance(attendance.Side, attendance.Name, attendance.Meal, attendance.Count)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("InternalServerError"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

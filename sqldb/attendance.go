package sqldb

import (
	"fmt"
	"time"
)

func CreateAttendance(side, name, meal string, count int) error {
	_, err := sqlDb.Exec(`
		INSERT INTO attendance (side, name, meal, count, timestamp)
		VALUES ($1, $2, $3, $4, $5)
	`, side, name, meal, count, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

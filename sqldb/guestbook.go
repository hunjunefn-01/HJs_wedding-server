package sqldb

import (
	"fmt"
	"time"

	"hjs-wedding-server/env"
	"hjs-wedding-server/types"
	"hjs-wedding-server/util"
)

func GetGuestbook(offset, limit int) (*types.GuestbookGetResponse, error) {
	rows, err := sqlDb.Query(`
		SELECT id, name, content, timestamp
		FROM guestbook
		WHERE valid = TRUE
		ORDER BY timestamp DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	guestbookGetResponse := &types.GuestbookGetResponse{
		Posts: []types.GuestbookPostForGet{},
	}

	// 전체 개수를 세는 쿼리를 QueryRow로 변경하여 더 효율적으로 처리합니다.
	err = sqlDb.QueryRow(`SELECT COUNT(*) FROM guestbook WHERE valid = TRUE`).Scan(&guestbookGetResponse.Total)
	if err != nil {
		// count 쿼리가 실패해도 목록은 보여줄 수 있도록 에러를 로깅만 할 수 있습니다.
		// 여기서는 기존 로직과 동일하게 에러를 반환합니다.
		return nil, err
	}

	for rows.Next() {
		guestbookPost := types.GuestbookPostForGet{}
		err := rows.Scan(&guestbookPost.Id, &guestbookPost.Name, &guestbookPost.Content, &guestbookPost.Timestamp)
		if err != nil {
			return nil, err
		}
		guestbookGetResponse.Posts = append(guestbookGetResponse.Posts, guestbookPost)
	}

	return guestbookGetResponse, nil
}

func CreateGuestbookPost(name, content, password string) error {
	phash, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	result, err := sqlDb.Exec(`
		INSERT INTO guestbook (name, content, password, timestamp)
		VALUES ($1, $2, $3, $4)
	`, name, content, phash, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("NO_ROWS_AFFECTED")
	}

	return nil
}

func DeleteGuestbookPost(id int, password string) error {
	passwordMatch := false
	if env.AdminPassword != "" && env.AdminPassword == password {
		passwordMatch = true
	} else {
		var phash string
		// 한 개의 결과만 예상되므로 QueryRow를 사용하는 것이 더 효율적입니다.
		err := sqlDb.QueryRow(`SELECT password FROM guestbook WHERE id = $1 AND valid = TRUE`, id).Scan(&phash)
		if err != nil {
			// 해당 id의 게시글이 없는 경우 sql.ErrNoRows 에러가 발생합니다.
			if err == sql.ErrNoRows {
				return fmt.Errorf("NO_GUESTBOOK_POST_FOUND")
			}
			return err
		}

		if util.CheckPasswordHash(password, phash) {
			passwordMatch = true
		}
	}

	if !passwordMatch {
		return fmt.Errorf("INCORRECT_PASSWORD")
	}

	result, err := sqlDb.Exec(`
		UPDATE guestbook
		SET valid = FALSE
		WHERE id = $1
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("NO_ROWS_AFFECTED")
	}

	return nil
}

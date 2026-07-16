import time

import config
import security
from db import get_db


class IncorrectPasswordError(Exception):
    pass


class PostNotFoundError(Exception):
    pass


def get_guestbook(offset: int, limit: int) -> dict:
    db = get_db()

    cursor = db.execute(
        """
        SELECT id, name, content, timestamp
        FROM guestbook
        WHERE valid = TRUE
        ORDER BY timestamp DESC
        LIMIT ? OFFSET ?
        """,
        (limit, offset),
    )
    rows = cursor.fetchall()

    total = db.execute("SELECT COUNT(*) FROM guestbook WHERE valid = TRUE").fetchone()[0]

    posts = [
        {"id": row[0], "name": row[1], "content": row[2], "timestamp": row[3]}
        for row in rows
    ]

    return {"posts": posts, "total": total}


def create_guestbook_post(name: str, content: str, password: str) -> None:
    db = get_db()
    phash = security.hash_password(password)

    db.execute(
        """
        INSERT INTO guestbook (name, content, password, timestamp)
        VALUES (?, ?, ?, ?)
        """,
        (name, content, phash, int(time.time())),
    )
    db.commit()


def delete_guestbook_post(post_id: int, password: str) -> None:
    db = get_db()

    password_match = False
    if config.ADMIN_PASSWORD and config.ADMIN_PASSWORD == password:
        password_match = True
    else:
        row = db.execute(
            "SELECT password FROM guestbook WHERE id = ? AND valid = TRUE", (post_id,)
        ).fetchone()
        if row is None:
            raise PostNotFoundError()

        if security.check_password_hash(password, row[0]):
            password_match = True

    if not password_match:
        raise IncorrectPasswordError()

    db.execute("UPDATE guestbook SET valid = FALSE WHERE id = ?", (post_id,))
    db.commit()

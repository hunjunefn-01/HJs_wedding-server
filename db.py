import os
import sqlite3

_db: sqlite3.Connection | None = None

_DEFAULT_DB_PATH = os.path.join(os.path.dirname(os.path.abspath(__file__)), "sql.db")


def init_db(path: str = _DEFAULT_DB_PATH) -> None:
    global _db
    _db = sqlite3.connect(path, check_same_thread=False)

    _db.execute(
        """
        CREATE TABLE IF NOT EXISTS guestbook (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name VARCHAR(20),
            content VARCHAR(200),
            password VARCHAR(20),
            timestamp INTEGER,
            valid BOOLEAN DEFAULT TRUE
        )
        """
    )
    _db.execute("CREATE INDEX IF NOT EXISTS guestbook_timestamp ON guestbook (timestamp)")
    _db.execute("CREATE INDEX IF NOT EXISTS guestbook_valid ON guestbook (valid)")

    _db.execute(
        """
        CREATE TABLE IF NOT EXISTS attendance (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            side VARCHAR(10),
            name VARCHAR(20),
            meal VARCHAR(20),
            count INTEGER,
            timestamp INTEGER
        )
        """
    )
    _db.commit()


def get_db() -> sqlite3.Connection:
    if _db is None:
        raise RuntimeError("Database not initialized — call init_db() first")
    return _db

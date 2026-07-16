import time

from db import get_db


def create_attendance(side: str, name: str, meal: str, count: int) -> None:
    db = get_db()
    db.execute(
        """
        INSERT INTO attendance (side, name, meal, count, timestamp)
        VALUES (?, ?, ?, ?, ?)
        """,
        (side, name, meal, count, int(time.time())),
    )
    db.commit()

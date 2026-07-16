from flask import Flask, jsonify, request
from flask_cors import CORS

import attendance
import config
import db
import guestbook

app = Flask(__name__)
CORS(
    app,
    origins=[config.ALLOW_ORIGIN],
    methods=["GET", "POST", "PUT"],
    allow_headers=["Content-Type"],
)


@app.route("/api/guestbook", methods=["GET", "POST", "PUT"])
def guestbook_route():
    if request.method == "GET":
        offset_q = request.args.get("offset")
        limit_q = request.args.get("limit")

        try:
            offset = int(offset_q)
            limit = int(limit_q)
        except (TypeError, ValueError):
            return "Bad Request: invalid offset or limit", 400

        try:
            result = guestbook.get_guestbook(offset, limit)
        except Exception:
            return "Internal Server Error", 500

        return jsonify(result)

    if request.method == "POST":
        body = request.get_json(silent=True)
        if body is None or not all(k in body for k in ("name", "content", "password")):
            return "Bad Request", 400

        try:
            guestbook.create_guestbook_post(body["name"], body["content"], body["password"])
        except Exception:
            return "Internal Server Error", 500

        return "", 200

    if request.method == "PUT":
        body = request.get_json(silent=True)
        if body is None or not all(k in body for k in ("id", "password")):
            return "Bad Request", 400

        try:
            guestbook.delete_guestbook_post(body["id"], body["password"])
        except guestbook.IncorrectPasswordError:
            return "Forbidden: incorrect password", 403
        except Exception:
            return "Internal Server Error", 500

        return "", 200

    return "Method Not Allowed", 405


@app.route("/api/attendance", methods=["POST"])
def attendance_route():
    body = request.get_json(silent=True)
    if body is None or not all(k in body for k in ("side", "name", "meal", "count")):
        return "Bad Request", 400

    try:
        attendance.create_attendance(body["side"], body["name"], body["meal"], body["count"])
    except Exception:
        return "Internal Server Error", 500

    return "", 200


if __name__ == "__main__":
    db.init_db()
    app.run(host="0.0.0.0", port=8080)

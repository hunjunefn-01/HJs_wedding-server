import os

from dotenv import load_dotenv

load_dotenv(os.path.join(os.path.dirname(os.path.abspath(__file__)), ".env"))

ALLOW_ORIGIN = os.environ.get("ALLOW_ORIGIN", "")
ADMIN_PASSWORD = os.environ.get("ADMIN_PASSWORD", "")

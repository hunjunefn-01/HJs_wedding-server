import os

from dotenv import load_dotenv

load_dotenv()

ALLOW_ORIGIN = os.environ.get("ALLOW_ORIGIN", "")
ADMIN_PASSWORD = os.environ.get("ADMIN_PASSWORD", "")

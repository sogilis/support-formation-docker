#!/usr/bin/env bash

PORT="${PORT:-8080}"
STORAGE_FOLDER="${APP_STORAGE_FOLDER:-"/app/data"}"
BASE_URL="${APP_BASE_URL:-"/"}"

echo "ENV VAR ARE: $PORT $STORAGE_FOLDER"

exec "$@" "--port" "$PORT" "--storage-path" "$STORAGE_FOLDER"
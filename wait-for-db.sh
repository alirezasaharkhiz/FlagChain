#!/bin/sh
set -e

host="db"
port=3306

echo "Waiting for MySQL..."

while ! nc -z $host $port; do
  echo "Waiting for MySQL at $host:$port..."
  sleep 2
done

echo "MySQL is up - executing command"

exec "$@"

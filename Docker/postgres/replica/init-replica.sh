#!/bin/bash
set -e

# 1) Đợi master sẵn sàng
until pg_isready -h postgres-master-event -p 5432 -U admin; do
  echo "Waiting for master to become available..."
  sleep 2
done

# 2) Xoá dữ liệu cũ của replica (nếu có)
rm -rf "$PGDATA"/*

# 3) Lấy base‑backup từ master
PGPASSWORD=replicator123 pg_basebackup \
  -h postgres-master-event \
  -p 5432 \
  -U replicator \
  -D "$PGDATA" \
  --wal-method=stream

# 4) Bật chế độ standby
touch "$PGDATA/standby.signal"
echo "primary_conninfo = 'host=postgres-master-event port=5432 user=replicator password=replicator123 dbname=eventdb'" \
  >> "$PGDATA/postgresql.auto.conf"

chown -R postgres:postgres "$PGDATA"

# 5) Khởi động PostgreSQL
exec docker-entrypoint.sh postgres -c config_file=/etc/postgresql/postgresql.conf

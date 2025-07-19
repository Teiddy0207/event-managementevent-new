-- Chuyển sang database postgres
\connect postgres;

-- Tạo user replicator nếu chưa tồn tại
DO
$$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_roles WHERE rolname = 'replicator'
   ) THEN
CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'replicator123';
END IF;
END
$$;

-- Tạo database eventdb nếu chưa tồn tại (phần này phải nằm ngoài DO $$)
SELECT 'CREATE DATABASE eventdb'
    WHERE NOT EXISTS (
    SELECT FROM pg_database WHERE datname = 'eventdb'
)\gexec

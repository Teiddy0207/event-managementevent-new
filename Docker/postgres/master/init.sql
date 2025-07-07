-- Chạy lệnh này để đảm bảo không bị lỗi khi khởi tạo
-- Nhưng phải tách riêng, không được nằm trong DO $$ ... $$

\connect postgres;

-- Kiểm tra có tồn tại trước rồi mới tạo
SELECT 'CREATE DATABASE mydb'
WHERE NOT EXISTS (
    SELECT FROM pg_database WHERE datname = 'mydb'
)\gexec

CREATE TABLE login_users (
                             id SERIAL PRIMARY KEY,
                             user_id INTEGER NOT NULL,
                             token TEXT NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
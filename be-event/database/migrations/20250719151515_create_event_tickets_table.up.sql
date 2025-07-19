CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    quantity_available INTEGER NOT NULL
);

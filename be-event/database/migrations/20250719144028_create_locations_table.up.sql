CREATE TABLE  locations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    city_id INTEGER NOT NULL,
    address TEXT NOT NULL
);

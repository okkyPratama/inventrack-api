-- +migrate Up
DROP TABLE IF EXISTS warehouses;

-- +migrate Down
CREATE TABLE warehouses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    capacity INTEGER NOT NULL
);
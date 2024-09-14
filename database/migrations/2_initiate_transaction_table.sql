-- +migrate Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    category_id INT REFERENCES categories(id),
    supplier_id INT REFERENCES suppliers(id)
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id),
    quantity INT NOT NULL,
    type VARCHAR(10) NOT NULL, -- 'in' or 'out'
    date TIMESTAMP NOT NULL,
    user_id INT REFERENCES users(id)
);

-- +migrate Down
DROP TABLE transactions;
DROP TABLE products;
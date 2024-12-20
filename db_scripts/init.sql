CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS products (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT,
    description TEXT,
    price DECIMAL,
    quantity integer
);

TRUNCATE TABLE products;
INSERT INTO products(name, description, price, quantity) VALUES
                ('T-shirt', 'Coolest t-shirt', 19.0, 8),
                ('Headset', 'Comfortable', 10.0, 5),
                ('Hat', 'Coolest hat', 9.5, 50);
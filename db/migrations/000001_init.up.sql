CREATE TABLE IF NOT EXISTS customers (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    email VARCHAR UNIQUE NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    address VARCHAR NOT NULL
);

create table if not exists products (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    title  VARCHAR NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    qty INT NOT NULL
);

create table if not exists transactions (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    customer_id uuid NOT NULL REFERENCES customers(id),
    product_id uuid NOT NULL REFERENCES products(id),
    qty INT NOT NULL,
    total_price DOUBLE PRECISION NOT NULL
);
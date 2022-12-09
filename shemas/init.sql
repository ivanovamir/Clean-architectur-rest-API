CREATE TABLE category (
    id serial primary key,
    title varchar(255),
    image varchar(500)
);

CREATE TABLE product (
    id serial primary key,
    category_id int references category(id),
    title varchar(255),
    vendor_code varchar(255),
    description varchar(255),
    price varchar(255),
    image varchar(500)
);

CREATE TABLE "user" (
    id serial primary key,
    created_at time not null,
    email varchar(255),
    password_hash varchar(255)
);


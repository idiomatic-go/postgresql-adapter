create table public.person
(
    id            integer      not null
        constraint person_pk
            primary key,
    first_name    varchar(100) not null,
    last_name     varchar(100) not null,
    date_of_birth date         not null
);

alter table public.person
    owner to postgres;

ALTER USER postgres PASSWORD 'mypassword';

created_at timestamp default now(): This final column definition will provide us with some extra meta-information
    about their user. If not specified explicitly, the created_at timestamp will default to the time the row was
    inserted.

INSERT INTO products (product_no, name, price) VALUES
    (1, 'Cheese', 9.99),
    (2, 'Bread', 1.99),
    (3, 'Milk', 2.99);
create database test;

create table if not exists products(
id serial primary key,
description varchar,
price real,
quantity integer,
createdAt timestamp default current_timestamp
);

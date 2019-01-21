create database test;

\c test;

create table if not exists images(
id bigserial primary key,
image bytea null
);

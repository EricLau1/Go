create database test;

create table if not exists users (
id int auto_increment primary key,
name varchar(50) not null,
email varchar(50) not null,
password varchar(255) not null,
created_at timestamp default current_timestamp());
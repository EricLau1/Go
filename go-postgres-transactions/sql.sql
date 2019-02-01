create database dbtransactions;

\c dbtransactions;

create table if not exists users(
id bigserial primary key,
name varchar(100) not null
);

create table if not exists feedback(
id bigserial primary key,
usr bigint not null,
comment varchar(255) not null,
constraint feedback_usr_fk foreign key(usr)
references users
);

create table if not exists rating(
id bigserial primary key,
feedback bigint unique,
likes integer default 0,
dislikes integer default 0,
reports integer default 0,
constraint rating_feedback_fk foreign key(feedback)
references feedback
on delete cascade
on update cascade
);


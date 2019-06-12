\c postgres;
drop database if exists distributed;

create database distributed;

\c distributed;

CREATE TABLE sensor (
    id serial primary key,
    name character varying(50) NOT NULL,
    serial_no character varying(50) NOT NULL,
    unit_type character varying(50) NOT NULL,
    max_safe_value double precision NOT NULL,
    min_safe_value double precision NOT NULL
);

CREATE TABLE sensor_reading (
    id bigserial primary key,
    value double precision NOT NULL,
    sensor_id integer,
    taken_on timestamp,
    constraint sr_sensor_id_fk foreign key(sensor_id) references sensor(id)
);

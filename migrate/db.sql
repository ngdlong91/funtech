drop database if exists funtech;

create database funtech;

use funtech;
drop table if exists product;

create table product
(
    id         bigint unsigned auto_increment
        primary key,
    quantity   int not null ,
    created_at int not null ,
    updated_at int not null
);

INSERT INTO product (quantity, created_at, updated_at) values (2, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO product (quantity, created_at, updated_at) values (18, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO product (quantity, created_at, updated_at) values (3, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
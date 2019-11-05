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
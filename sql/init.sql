create database if not exists go_signup;

use go_signup;

show tables;

create table if not exists authentication_users (
    id         varchar(36) not null unique,
    email      varchar(50) not null unique,
    username   varchar(20) not null unique,
    phone      varchar(15) not null unique,
    password   varchar(100) not null,
    created_at datetime not null,
    updated_at datetime not null,
    constraint authentication_users_id_pk primary key (id)
    )
    engine = INNODB
    default char set = UTF8;

desc authentication_users;

# drop table authentication_users;
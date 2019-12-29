
create database dropwizcrud;
use dropwizcrud;

create table users (
	id bigint primary key not null auto_increment,
    first_name varchar(255) not null,
    middle_name varchar(255) not null,
    last_name varchar(255) not null,
    username varchar(255) not null
);

create table departments (
	id bigint primary key not null auto_increment,
    dept_name varchar(255) not null    
);
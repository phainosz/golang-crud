CREATE DATABASE IF NOT EXISTS golang_crud;

USE golang_crud;

CREATE TABLE IF NOT EXISTS users (
  id int auto_increment primary key,
  name varchar(50) not null,
  email varchar(50) not null unique
) ENGINE=InnoDB;
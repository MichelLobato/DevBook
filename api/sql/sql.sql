CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email varchar(50) not null unique,
    senha varchar(50) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE = INNODB;
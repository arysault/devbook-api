CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nickname VARCHAR(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    criadoEm timestamp default current_timestamp()
 ) ENGINE=INNODB;
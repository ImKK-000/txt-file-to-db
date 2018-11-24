CREATE DATABASE txt_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use txt_db;
CREATE TABLE txt_table(
    id int NOT NULL AUTO_INCREMENT,
    data varchar(100),
    PRIMARY KEY (id)
);

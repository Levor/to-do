CREATE SCHEMA IF NOT EXISTS `todo`;
CREATE TABLE IF NOT EXISTS `todo`.`users`
(
    id         int auto_increment,
    login      CHAR(255) not null,
    password   CHAR(255) not null,
    name       CHAR(255) null,
    first_name CHAR(255) null,
    constraint users_pk
        primary key (id)
);
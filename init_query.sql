-- create database productdb;

CREATE TABLE Products (
    id    integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    model varchar(30) NOT NULL,
    company varchar(30) NOT NULL,
    price integer NOT NULL
);

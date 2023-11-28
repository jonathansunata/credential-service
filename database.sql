/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/*
*/
CREATE TABLE users (
                       id serial primary key,
                       full_name varchar(50) not null,
                       phone_number varchar(50),
                       password varchar(100),
                       salt VARCHAR(50),
                       count_login INT
);

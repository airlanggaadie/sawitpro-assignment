/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */

-- just table with all about user information
CREATE TABLE IF NOT EXISTS "users" (
  id                  UUID PRIMARY KEY NOT NULL,
  fullname            VARCHAR(60) NOT NULL,
  phonenumber				  VARCHAR(13) NOT NULL UNIQUE,
  -- also can add another field for user information
  created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- table with user authentication. this table relation 1:1 with users
CREATE TABLE IF NOT EXISTS "user_password" (
  user_id             UUID NOT NULL UNIQUE,
  password			      VARCHAR(100) NOT NULL, -- hashed password
  created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES "users" (id)
);

-- table with daily count of user login. this table relation 1:many with users
CREATE TABLE IF NOT EXISTS "user_login" (
  user_id             UUID NOT NULL UNIQUE,
  date                DATE NOT NULL DEFAULT CURRENT_DATE,
  count               BIGINT NOT NULL DEFAULT 1,
  FOREIGN KEY (user_id) REFERENCES "users" (id),
  PRIMARY KEY (user_id, date)
);

/* for cleanup database */
-- DROP TABLE IF EXISTS "user_login";
-- DROP TABLE IF EXISTS "user_password";
-- DROP TABLE IF EXISTS "users";
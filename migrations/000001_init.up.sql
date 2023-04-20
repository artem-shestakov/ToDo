CREATE TABLE users
(
    id serial PRIMARY KEY,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL
);

CREATE TABLE lists
(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL UNIQUE,
    description text
);

CREATE TABLE users_lists
(
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    list_id integer REFERENCES lists(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE tasks
(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL UNIQUE,
    description text,
    is_done boolean NOT NULL DEFAULT false,
    list_id integer,
    FOREIGN KEY (list_id) REFERENCES lists(id)
);


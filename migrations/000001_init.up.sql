CREATE TABLE users
(
    user_id int PRIMARY KEY,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL
);

CREATE TABLE lists
(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE tasks
(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text,
    is_done boolean NOT NULL DEFAULT false,
    list_id integer,
    FOREIGN KEY (list_id) REFERENCES lists(id)
);


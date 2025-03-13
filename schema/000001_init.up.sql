CREATE TABLE users (
                       id serial not null unique,
                       name varchar(255) not null,
                       username varchar(255) not null unique,
                       password_hash varchar(255) not null
);

CREATE TABLE todo_lists (
                            id serial not null unique,
                            title varchar(255) not null,
                            description varchar(255)
);

CREATE TABLE users_lists (
                             id serial not null unique,
                             user_id int NOT NULL REFERENCES users (id) ON DELETE CASCADE,
                             list_id int NOT NULL REFERENCES todo_lists (id) ON DELETE CASCADE
);

CREATE TABLE todo_items (
                            id serial not null unique,
                            title varchar(255) NOT NULL,
                            description varchar(255),
                            done boolean not null default false
);

CREATE TABLE lists_items (
                             id serial not null unique,
                             item_id int NOT NULL REFERENCES todo_items (id) ON DELETE CASCADE,
                             list_id int NOT NULL REFERENCES todo_lists (id) ON DELETE CASCADE
);

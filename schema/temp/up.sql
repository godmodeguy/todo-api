CREATE TABLE Users (
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE Tasks (
    id          serial         not null unique,
    list_id     int            references Users(id) on delete cascade not null,
    title       varchar(255)   not null,
    description varchar(255),
    status      bool default   false,
    priority    int default    0
);

CREATE TABLE LinkType (
    id      serial          not null unique,
    type    varchar(255)    not null
);

CREATE TABLE TodoList (
    id          serial          not null unique,
    title       varchar(255)    not null,
    description varchar(255)
);

CREATE TABLE UserList (
    id          serial not null unique,
    list_id     int references TodoList(id) on delete cascade not null,
    user_id     int references Users(id) on delete cascade not null,
    type_id     int references LinkType(id) on delete cascade not null
);


DROP TABLE sessions;
DROP TABLE users;

CREATE TABLE users
(
    nickname varchar(30) unique,
    email varchar(30) unique,
    password varchar not null,
    phone varchar(20) not null,
    primary key (nickname, email)
);

CREATE TABLE sessions
(
    nickname varchar(30) not null,
    cookie varchar(50) not null primary key,
    expiration timestamp not null default current_timestamp,
    foreign key (nickname) references users (nickname) on delete cascade
);


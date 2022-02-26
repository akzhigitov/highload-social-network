CREATE TABLE IF NOT EXISTS user
(
    user_id int   not null primary key  AUTO_INCREMENT,
    email      varchar(255) not null,
    password   binary(60)   not null,
    last_name  varchar(255) not null,
    first_name varchar(255) not null,
    city    varchar(100)   not null,
    interests varchar(255) not null,
    sex        tinyint      not null,
    birthday   date         not null,
     UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS friends
(
    user_id1   int not null,
    user_id2  int not null,
    PRIMARY KEY (user_id1, user_id2),
    FOREIGN KEY (user_id1)
        REFERENCES user(user_id),
    FOREIGN KEY (user_id2)
        REFERENCES user(user_id)
);

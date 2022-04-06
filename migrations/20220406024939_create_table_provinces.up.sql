DROP TABLE IF EXISTS provinces;

CREATE TABLE IF NOT EXISTS provinces
(
    id   char(2)      not null,
    name varchar(255) not null,
    primary key (id)
);

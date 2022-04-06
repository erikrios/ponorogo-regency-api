DROP TABLE IF EXISTS districts;

CREATE TABLE IF NOT EXISTS districts
(
    id         char(7)      not null,
    regency_id char(4)      not null,
    name       varchar(255) not null,
    primary key (id),
    constraint districts_regency_id_foreign
        foreign key (regency_id)
            references regencies (id) on delete cascade
);

DROP TABLE IF EXISTS villages;

CREATE TABLE IF NOT EXISTS villages
(
    id          char(10)     not null,
    district_id char(7)      not null,
    name        varchar(255) not null,
    primary key (id),
    constraint villages_district_id_foreign
        foreign key (district_id)
            references districts (id) on delete cascade
);

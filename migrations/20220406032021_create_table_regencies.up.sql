DROP TABLE IF EXISTS regencies;

CREATE TABLE IF NOT EXISTS regencies
(
    id          char(4)      not null,
    province_id char(2)      not null,
    name        varchar(255) not null,
    primary key (id),
    constraint regencies_province_id_foreign
        foreign key (province_id)
            references provinces (id) on delete cascade
);

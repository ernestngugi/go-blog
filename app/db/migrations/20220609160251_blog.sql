-- +goose Up
CREATE TABLE blogs(
    id int primary key not null auto_increment,
    title varchar(50) not null,
    description text not null,
    status varchar(10) not null,
    date_created timestamp default CURRENT_TIMESTAMP,
    date_modified timestamp null
);

-- +goose Down
drop table if exists blogs;

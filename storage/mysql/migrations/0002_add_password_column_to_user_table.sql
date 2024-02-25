-- +migrate Up
ALTER TABLE users add column password varchar(455) NOT NULL;

-- +migrate Down
ALTER TABLE users drop column password;


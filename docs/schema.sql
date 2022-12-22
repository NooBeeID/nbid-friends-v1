CREATE TABLE auth (
    id SERIAL PRIMARY KEY,
    email varchar,
    password varchar
);

CREATE TABLE follows (
    id SERIAL PRIMARY KEY,
    authId int,
    followingId int,
    created_at timestamp
);

ALTER TABLE auth 
ADD COLUMN img_url varchar;
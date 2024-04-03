CREATE TABLE events
(
    id    uuid primary key,
    title varchar,
    date  date,
    user_id uuid
);
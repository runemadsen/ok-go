CREATE TABLE posts (
  id serial primary key,
  title text,
  body text,
  created_at timestamp with time zone
  updated_at timestamp with time zone
);
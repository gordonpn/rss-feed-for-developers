\connect app_database
CREATE TABLE posts (
  title TEXT NOT NULL,
  link TEXT NOT NULL,
  summary TEXT,
  published TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  author TEXT,
  id TEXT NOT NULL UNIQUE
);

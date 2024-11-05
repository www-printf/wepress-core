CREATE TABLE IF NOT EXISTS users (
    id VARCHAR PRIMARY KEY,
    fullname VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO users (id, fullname, password, email)
VALUES ('4b793c1a-06ea-4ea0-a2b0-19e3c04c3f1d', 'Nguyen Van A', '$2a$10$B6QS.AyoNoK3FezoX8lsNOmuEIc0VaBNl6lxB8cMiSyL0NNl5PvrK', 'example@email.com');
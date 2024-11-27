CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fullname VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    privkey VARCHAR,
    pubkey VARCHAR
);

INSERT INTO users (id, fullname, password, email, privkey, pubkey)
VALUES ('4b793c1a-06ea-4ea0-a2b0-19e3c04c3f1d', 'Nguyen Van A', '$2a$10$B6QS.AyoNoK3FezoX8lsNOmuEIc0VaBNl6lxB8cMiSyL0NNl5PvrK', 'example@email.com', 'XQHNFZsKNhdJdywJ9xYioFEfkZKSnvk5BmfTNeXbQyhHe+hpggA2mjTLog7p1yw895NgDpTYfV9OTzMrS84IdA==','R3voaYIANpo0y6IO6dcsPPeTYA6U2H1fTk8zK0vOCHQ=');
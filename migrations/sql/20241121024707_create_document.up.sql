CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    object_key VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    owner_id UUID NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users(id) 
        ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS metadata (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    document_id UUID NOT NULL,
    mime_type VARCHAR NOT NULL,
    size INT NOT NULL,
    extension VARCHAR NOT NULL,
    FOREIGN KEY (document_id) REFERENCES documents(id) 
        ON DELETE CASCADE ON UPDATE CASCADE
);
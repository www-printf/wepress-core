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

INSERT INTO documents (id, object_key, owner_id) 
VALUES ('de8f8e50-50b4-4080-acc3-0c0cf40259e9' ,'4b793c1a06ea4ea0a2b019e3c04c3f1d/15a3171b27f7429a81c7e58ba82c2a15', '4b793c1a-06ea-4ea0-a2b0-19e3c04c3f1d');

INSERT INTO metadata (id, name, document_id, mime_type, size, extension)
VALUES ('e0b930ba-b1ff-4c00-802d-6f2567e2d246', 'doc.pdf', 'de8f8e50-50b4-4080-acc3-0c0cf40259e9', 'application/pdf', 1024, 'pdf');
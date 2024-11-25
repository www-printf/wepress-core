CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS print_jobs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    status VARCHAR,
    start_at TIMESTAMP DEFAULT NOW(),
    end_at TIMESTAMP DEFAULT NOW(),
    num_of_copies INTEGER,
    paper_size VARCHAR,
    orientation VARCHAR,
    mode VARCHAR,
    pages_to_print VARCHAR,
    document_id UUID NOT NULL,
    FOREIGN KEY (document_id) REFERENCES documents(id) 
);
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS print_histories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    job_id VARCHAR NOT NULL,
    printer_id  INTEGER NOT NULL,
    cluster_id INTEGER NOT NULL,

    FOREIGN KEY (printer_id) REFERENCES printers(id) 
        ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (cluster_id) REFERENCES clusters(id)
        ON DELETE SET NULL ON UPDATE CASCADE
);
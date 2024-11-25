CREATE TABLE IF NOT EXISTS clusters (
    id SERIAL PRIMARY KEY,
    status VARCHAR,
    added_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    building VARCHAR,
    room VARCHAR,
    campus VARCHAR,
    longitude FLOAT,
    latitude FLOAT
);

INSERT INTO clusters (status, building, room, campus, longitude, latitude)
VALUES 
('active', 'A5', '109', 'Ly Thuong Kiet', 106.676292, 10.762622),
('active', 'B3', '103', 'Ly Thuong Kiet', 106.676180, 10.762500),
('active', 'H6', '104', 'Di An', 106.680920, 10.762890),
('active', 'H1', '105', 'Di An', 106.681120, 10.763000);

CREATE TABLE IF NOT EXISTS printers (
    id SERIAL PRIMARY KEY,
    status VARCHAR,
    added_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    cluster_id INTEGER NOT NULL,
    manufacturer VARCHAR,
    model VARCHAR,
    serial_number VARCHAR,
    ip_address VARCHAR UNIQUE,
    mac_address VARCHAR UNIQUE,

    FOREIGN KEY (cluster_id) REFERENCES clusters(id) 
        ON DELETE SET NULL ON UPDATE CASCADE
);

INSERT INTO printers (status, cluster_id, manufacturer, model, serial_number, ip_address, mac_address)
VALUES
('active', 1, 'HP', 'LaserJet Pro M404dn', 'SN123456789', '192.168.1.10', '3C:52:82:00:01:01'),
('inactive', 1, 'HP', 'DeskJet 2331', 'SN1234432112', '192.168.1.18', '3C:52:82:00:01:02'),
('active', 1, 'Canon', 'i-SENSYS LBP623Cdw', 'SN987654321', '192.168.1.11', '00:1E:8F:00:01:01'),
('active', 2, 'Canon', 'i-SENSYS MF244dw', 'SN3213213210', '192.168.1.19', '00:1E:8F:00:01:02'),
('active', 3, 'Canon', 'PIXMA TS5370', 'SN1231231234', '192.168.1.15', '00:1E:8F:00:01:03'),
('active', 2, 'Brother', 'HL-L2350DW', 'SN1122334455', '192.168.1.12', '40:16:7E:00:01:01'),
('active', 4, 'Brother', 'DCP-T420W', 'SN4445556667', '192.168.1.16', '40:16:7E:00:01:02'),
('active', 2, 'Epson', 'EcoTank L3150', 'SN6677889900', '192.168.1.13', '00:26:33:00:01:01'),
('active', 4, 'Epson', 'WorkForce WF-7710', 'SN7778889990', '192.168.1.17', '00:26:33:00:01:02'),
('inactive', 3, 'HP', 'OfficeJet Pro 8020', 'SN1234509876', '192.168.1.14', '3C:52:82:00:01:03');

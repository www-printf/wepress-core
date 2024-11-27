CREATE TABLE IF NOT EXISTS clusters (
    id SERIAL PRIMARY KEY,
    added_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    building VARCHAR,
    room VARCHAR,
    campus VARCHAR,
    longitude FLOAT,
    latitude FLOAT
);

INSERT INTO clusters (building, room, campus, longitude, latitude)
VALUES 
('A5', '109', 'Ly Thuong Kiet', 106.676292, 10.762622),
('B3', '103', 'Ly Thuong Kiet', 106.676180, 10.762500),
('H6', '104', 'Di An', 106.680920, 10.762890),
('H1', '105', 'Di An', 106.681120, 10.763000);

CREATE TABLE IF NOT EXISTS printers (
    id SERIAL PRIMARY KEY,
    added_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    cluster_id INTEGER NOT NULL,
    manufacturer VARCHAR,
    model VARCHAR,
    serial_number VARCHAR,
    uri VARCHAR UNIQUE,

    FOREIGN KEY (cluster_id) REFERENCES clusters(id) 
        ON DELETE SET NULL ON UPDATE CASCADE
);

-- INSERT INTO printers (cluster_id, manufacturer, model, serial_number, uri)
-- VALUES
-- (1, 'HP', 'LaserJet Pro M404dn', 'SN123456789', '192.168.1.10:50051'),
-- (1, 'HP', 'DeskJet 2331', 'SN1234432112', '192.168.1.18:50051'),
-- (1, 'Canon', 'i-SENSYS LBP623Cdw', 'SN987654321', '192.168.1.11:50051'),
-- (2, 'Canon', 'i-SENSYS MF244dw', 'SN3213213210', '192.168.1.19:50051'),
-- (3, 'Canon', 'PIXMA TS5370', 'SN1231231234', '192.168.1.15:50051'),
-- (2, 'Brother', 'HL-L2350DW', 'SN1122334455', '192.168.1.12:50051'),
-- (4, 'Brother', 'DCP-T420W', 'SN4445556667', '192.168.1.16:50051'),
-- (2, 'Epson', 'EcoTank L3150', 'SN6677889900', '192.168.1.13:50051'),
-- (4, 'Epson', 'WorkForce WF-7710', 'SN7778889990', '192.168.1.17:50051'),
-- (3, 'HP', 'OfficeJet Pro 8020', 'SN1234509876', '192.168.1.14:50051');

INSERT INTO printers (cluster_id, manufacturer, model, serial_number, uri)
VALUES
(1, 'HP', 'LaserJet Pro M404dn', 'SN123456789', 'localhost:50001'),
(1, 'HP', 'DeskJet 2331', 'SN1234432112', 'localhost:50002'),
(1, 'Canon', 'i-SENSYS LBP623Cdw', 'SN987654321', 'localhost:50003'),
(2, 'Canon', 'i-SENSYS MF244dw', 'SN3213213210', 'localhost:50004'),
(3, 'Canon', 'PIXMA TS5370', 'SN1231231234', 'localhost:50005'),
(2, 'Brother', 'HL-L2350DW', 'SN1122334455', 'localhost:50006'),
(4, 'Brother', 'DCP-T420W', 'SN4445556667', 'localhost:50007'),
(2, 'Epson', 'EcoTank L3150', 'SN6677889900', 'localhost:50008'),
(4, 'Epson', 'WorkForce WF-7710', 'SN7778889990', 'localhost:50009'),
(3, 'HP', 'OfficeJet Pro 8020', 'SN1234509876', 'localhost:50010');

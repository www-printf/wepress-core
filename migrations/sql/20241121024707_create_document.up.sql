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

-- Documents for user@email.com

INSERT INTO documents (id, object_key, owner_id)
VALUES 
('06807dcc-c515-4f03-8dbd-26317df02ab0', 'c1d2e3f45678901234567890abcdef12/0467e6f559c74e45a98df7bd1c06ecd0', 'c1d2e3f4-5678-9012-3456-7890abcdef12'),
('1eb692d9-e80d-4a01-9672-0e6c4814da0e', 'c1d2e3f45678901234567890abcdef12/649cc04604654931a7f87d5a8ce2e319', 'c1d2e3f4-5678-9012-3456-7890abcdef12'),
('b12e9ed0-d27c-4441-a138-d68955617fa3', 'c1d2e3f45678901234567890abcdef12/c63982cbfc1e433080e416736469030b', 'c1d2e3f4-5678-9012-3456-7890abcdef12'),
('b294273a-5c77-46d3-9383-e64a8957b22f', 'c1d2e3f45678901234567890abcdef12/6c556717050f467699bf1b34ac93229c', 'c1d2e3f4-5678-9012-3456-7890abcdef12'),
('a6013537-bc6a-4ae6-a371-409e28d8ed21', 'c1d2e3f45678901234567890abcdef12/7d429f5fa5b84fb89453ba6a890839b8', 'c1d2e3f4-5678-9012-3456-7890abcdef12'),
('6d4b79db-4a63-4161-89c7-3e32ffedb532', 'c1d2e3f45678901234567890abcdef12/25637743b966406084f4cf3f593dbdf7', 'c1d2e3f4-5678-9012-3456-7890abcdef12'),
('b1c36fbc-d9bd-4af7-8bb8-ee58a90f4e18', 'c1d2e3f45678901234567890abcdef12/b2da2e9547214b3eb4bb9ed8ea869f45', 'c1d2e3f4-5678-9012-3456-7890abcdef12'),
('029853b2-8423-42f5-b55c-4305094c4046', 'c1d2e3f45678901234567890abcdef12/f0c47c73d5704f42a35aeb4ff2753af3', 'c1d2e3f4-5678-9012-3456-7890abcdef12');

INSERT INTO metadata (id, name, document_id, mime_type, size, extension)
VALUES 
('3c2aa3f9-b997-470a-a752-fdd603a2a131', 'Capstone_Project_Autumn_2023.pdf', '06807dcc-c515-4f03-8dbd-26317df02ab0', 'application/pdf', 600000, 'pdf'),
('d5e3ba40-67f6-4521-8784-929f3ce8712d', 'HK241- Assignment 1-Network Application P2P File Sharing.pdf', '1eb692d9-e80d-4a01-9672-0e6c4814da0e', 'application/pdf', 700000, 'pdf'),
('9bf8b47c-46c5-47d8-ae00-f79136259d1e', 'HK241-_NetworkDesignForABank_Assig2.pdf', 'b12e9ed0-d27c-4441-a138-d68955617fa3', 'application/pdf', 750000, 'pdf'),
('74183f9c-655b-4e53-b56a-bb2bb1473db7', 'Mô tả BTL1-HK241.pdf', 'b294273a-5c77-46d3-9383-e64a8957b22f', 'application/pdf', 750000, 'pdf'),
('2b87fc86-5479-4987-9671-d156dc07bb70', 'Mô tả BTL2-HK241.pdf', 'a6013537-bc6a-4ae6-a371-409e28d8ed21', 'application/pdf', 750000, 'pdf'),
('bf94083f-490f-4af2-aa20-4b074d22ca2d', 'Lab01.pdf', '6d4b79db-4a63-4161-89c7-3e32ffedb532', 'application/pdf', 750000, 'pdf'),
('9cf687dc-ff5f-4750-bc1b-49ace9e062c3', 'Lab03.pdf', 'b1c36fbc-d9bd-4af7-8bb8-ee58a90f4e18', 'application/pdf', 750000, 'pdf'),
('59c5aae6-07a8-4137-b71f-6730db626c1f', 'Lab05.pdf', '029853b2-8423-42f5-b55c-4305094c4046', 'application/pdf', 750000, 'pdf');
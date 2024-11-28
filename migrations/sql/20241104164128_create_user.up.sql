CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        CREATE TYPE user_role AS ENUM ('user', 'spso');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fullname VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    privkey VARCHAR,
    pubkey VARCHAR,
    role user_role DEFAULT 'user'
);

INSERT INTO users (id, fullname, password, email, privkey, pubkey, role)
VALUES 
('4b793c1a-06ea-4ea0-a2b0-19e3c04c3f1d', 'Hoàng Anh Hùng', '$2a$10$B6QS.AyoNoK3FezoX8lsNOmuEIc0VaBNl6lxB8cMiSyL0NNl5PvrK', 'hoanganhhung@email.com', 'XQHNFZsKNhdJdywJ9xYioFEfkZKSnvk5BmfTNeXbQyhHe+hpggA2mjTLog7p1yw895NgDpTYfV9OTzMrS84IdA==','R3voaYIANpo0y6IO6dcsPPeTYA6U2H1fTk8zK0vOCHQ=', 'spso'),
('5f7d8c2b-7e9b-4e8e-8c5e-19e3c04c3f2e', 'Bùi Minh Khôi', '$2a$10$qFQi/UrEhiEXyQTBMEzc2u0j8A3mZM8E/EBvc0OEkMIgpQN9hFSE2', 'buiminhkhoi@email.com', 'zc5xmohUaZT6VC7uzKTW2b7K0/erNGaPZrw7kwc/GtVmjhbSop8LWmf+eqKkWoFAH1NBCyWt4ofLq4p/Qd5Tdg==', 'Zo4W0qKfC1pn/nqipFqBQB9TQQslreKHy6uKf0HeU3Y=', 'spso'),
('6a8e9d3c-8f0c-4f9f-9d6f-29e3c04c3f3f', 'Trần Anh Khôi', '$2a$10$gsra4hjZRfrIVAso4NaVMu8JYj4cXYQLaqU0.GPdjU.nVCPKnYrF.', 'trananhkhoi@email.com', 'DLNiBAzcinNatiQTx99arGhwLueT3eckLE/JnxDjAQuJxrLLqrFxBoNASUwk6uuQogZldLiD7OFmNvcvBDeXug==', 'icayy6qxcQaDQElMJOrrkKIGZXS4g+zhZjb3LwQ3l7o=', 'spso'),
('7b9f0e4d-9f1d-5f0f-0e7f-39e3c04c3f4f', 'Nguyễn Trung Kiên', '$2a$10$r6RwbmHzITe0TvHrFetOaeOudc9k.ZOl0AoGW2vkhLYc.ZE7lxWC.', 'nguyentrungkien@email.com', 'LTPR/IsCr23+bHm04iYWlUEHZgzES4qqrPZK4n9qHfjwDCBDiYc+3vRIFisykezc4bNW7zjyErHlq5y7BA8cYg==', '8AwgQ4mHPt70SBYrMpHs3OGzVu848hKx5aucuwQPHGI=', 'user'),
('8c0f1f5e-0f2e-6f1f-1f8f-49e3c04c3f5f', 'Hồ Quốc Khương', '$2a$10$vAPuxfkCiNFlXEFL/WRpZOtS/Kmnjmhn.XJ6ZUmdZEqF72VcHEqD.', 'hoquockhuong@email.com', 'uwKCq863PoRgev8ZzMku8Y+IvzRv+DlPQD997MBJh2lRJKaFiLQ3XQs8X6KX6iJsAeyFesr/r7pPXnd3kjcMag==', 'USSmhYi0N10LPF+il+oibAHshXrK/6+6T153d5I3DGo=', 'user'),
('9d1f206f-1f3f-7f2f-2f9f-59e3c04c3f6f', 'Huỳnh Ngọc Duy Khương', '$2a$10$ehvWTNvqzahKhZ/tqOSRbuwu5t6wmA7p7bMyQiQE7K1Jvzn4aAQkG', 'huynhngocduykhuong@email.com', 'Ts4QdddwS8dx321whMEV6I+RnMLNadnvAWOaqqJyxXxeKRFRwTcchVN25jP4tFxx0n1CZ0tXhsMpXJDn98uCWA==', 'XikRUcE3HIVTduYz+LRccdJ9QmdLV4bDKVyQ5/fLglg=', 'user'),
('ae2f3170-2f4f-8f3f-3fa0-69e3c04c3f7f', 'Nguyễn Phan Duy Bảo', '$2a$10$PLcNPTBN1T4nfvJzoKkZjOb4qNG80NzPzkm5T.Zv7VuZ8MlMsgpYW', 'nguyenphanduybao@email.com', '7Oxrd/Sp3BXKlwwkKlDCA23IM7Wu7LjRY4wMNjsfjX5XGU5P+qWr2wjLCrzm/7UOF/0DzLBxT7lvcgFQkcf5CQ==', 'VxlOT/qlq9sIywq85v+1Dhf9A8ywcU+5b3IBUJHH+Qk=', 'user'),
('bf3f4281-3f5f-9f4f-4fb0-79e3c04c3f8f', 'Đặng Quốc Bảo', '$2a$10$q9DT2ItyHcBht.aZ2ubWDumLXX.W/pjDfTiWiljHNNYv0wBN/Os6W', 'dangquocbao@email.com', 'hIozlPfP2TAzKqVmYdAN/R0d5BfwFDJ+2OYEt3S0uargyxPCDWDXqJ09RuMYQVTtvRNT/54VYltL++8DPcbAeg==', '4MsTwg1g16idPUbjGEFU7b0TU/+eFWJbS/vvAz3GwHo=', 'spso'),
('c1d2e3f4-5678-9012-3456-7890abcdef12', 'Nguyễn Văn A', '$2a$10$zDF/3MeQHTjEpa6ISvCvaegWbJsRtWbAZLT40PEgzqX/F8cAofeM6', 'user@email.com', 'mbVA3m+ier4vvHgZ4ODXFdGXdghK/ROyzVTJGT4ZHJqYDInWsO96ADeb1yjaDE4wD1TQ2puK1j68CgP1ij33TA==', 'mAyJ1rDvegA3m9co2gxOMA9U0NqbitY+vAoD9Yo990w=', 'user'),
('d3e4f5a6-7890-1234-5678-abcdef123456', 'Nguyễn Văn B', '$2a$10$dS4qWb7lI7sF1KeZmAG/NOXceCyONidK.MVxW7nqAc6JUOB4I3XGO', 'spso@email.com', 'UwI8ULwNqFLqKEaUgKEcOjUBtT2UMlmpkKP1lIEdu16mqWrlEsOLn6ZKazxMicc7Qf9tkEzHcb38U+YfZJQcqQ==', 'pqlq5RLDi5+mSms8TInHO0H/bZBMx3G9/FPmH2SUHKk=', 'spso');
-- Extensionlar
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table (store, admin, warehouse operator)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name TEXT NOT NULL,
    phone TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('admin', 'store', 'warehouse_operator')),
    created_at TIMESTAMP DEFAULT now()
);
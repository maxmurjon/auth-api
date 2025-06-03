CREATE TABLE stores (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    store_name TEXT NOT NULL,
    address TEXT,
    is_active BOOLEAN DEFAULT TRUE
);
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(20) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role VARCHAR(10) DEFAULT 'user',
    isBlocked BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW()
);
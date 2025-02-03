-- Enable the UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "User" (
    id BIGINT PRIMARY KEY,
    first_name TEXT NOT NULL, 
    last_name TEXT NOT NULL, 
    phone_number TEXT NOT NULL, 
    referal_code BIGINT NOT NULL, 
    referal_count INT DEFAULT 0,
    referred_by BIGINT REFERENCES "User"(id) ON DELETE SET NULL, -- New column to track referrals
    language TEXT,
    profile_pic TEXT,
    net_balance NUMERIC(18,2) DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "Task" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    link TEXT NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    reward INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "TaskUser" (
    user_id BIGINT REFERENCES "User"(id) ON DELETE CASCADE,
    task_id UUID REFERENCES "Task"(id) ON DELETE CASCADE,
    total_tries INTEGER DEFAULT 0,
    started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    finished_at TIMESTAMP,
    PRIMARY KEY (user_id, task_id) -- Composite primary key
);

CREATE TABLE "Withdrawal" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id BIGINT REFERENCES "User"(id) ON DELETE CASCADE,
    amount NUMERIC(18,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    status TEXT CHECK (status IN ('completed', 'onprogress')) NOT NULL
);

CREATE TABLE "Admin" (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    profile_pic TEXT
);

CREATE TABLE "FAQ" (
    id SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    answer TEXT NOT NULL
);

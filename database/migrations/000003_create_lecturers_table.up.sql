-- ============================================
-- Migration 3: Lecturers Table
-- File: database/migrations/000003_create_lecturers_table.up.sql
-- ============================================

CREATE TABLE IF NOT EXISTS lecturers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nip VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    date_of_birth DATE,
    gender VARCHAR(10) CHECK (gender IN ('male', 'female')),
    department VARCHAR(100) NOT NULL,
    position VARCHAR(50),
    specialization VARCHAR(100),
    education_level VARCHAR(50),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'retired')),
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_lecturers_nip ON lecturers(nip);
CREATE INDEX idx_lecturers_email ON lecturers(email);
CREATE INDEX idx_lecturers_department ON lecturers(department);
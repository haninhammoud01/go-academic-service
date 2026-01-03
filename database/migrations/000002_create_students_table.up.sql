-- ============================================
-- Migration 2: Students Table
-- File: database/migrations/000002_create_students_table.up.sql
-- ============================================

CREATE TABLE IF NOT EXISTS students (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nim VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    date_of_birth DATE,
    gender VARCHAR(10) CHECK (gender IN ('male', 'female')),
    major VARCHAR(100) NOT NULL,
    enrollment_year INTEGER NOT NULL,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'graduated', 'dropped')),
    gpa DECIMAL(3,2) DEFAULT 0.00,
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_students_nim ON students(nim);
CREATE INDEX idx_students_email ON students(email);
CREATE INDEX idx_students_major ON students(major);
CREATE INDEX idx_students_status ON students(status);
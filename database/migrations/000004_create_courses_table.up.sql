-- ============================================
-- Migration 4: Courses Table
-- File: database/migrations/000004_create_courses_table.up.sql
-- ============================================

CREATE TABLE IF NOT EXISTS courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    credits INTEGER NOT NULL CHECK (credits > 0),
    semester INTEGER NOT NULL CHECK (semester > 0),
    department VARCHAR(100) NOT NULL,
    course_type VARCHAR(50) CHECK (course_type IN ('mandatory', 'elective')),
    max_students INTEGER DEFAULT 40,
    lecturer_id UUID REFERENCES lecturers(id) ON DELETE SET NULL,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_courses_code ON courses(code);
CREATE INDEX idx_courses_semester ON courses(semester);
CREATE INDEX idx_courses_department ON courses(department);
CREATE INDEX idx_courses_lecturer_id ON courses(lecturer_id);
-- ============================================
-- Migration 5: Enrollments Table (KRS)
-- File: database/migrations/000005_create_enrollments_table.up.sql
-- ============================================

CREATE TABLE IF NOT EXISTS enrollments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id UUID NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    academic_year VARCHAR(10) NOT NULL,
    semester INTEGER NOT NULL CHECK (semester > 0),
    enrollment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) DEFAULT 'enrolled' CHECK (status IN ('enrolled', 'completed', 'dropped', 'failed')),
    grade VARCHAR(2) CHECK (grade IN ('A', 'AB', 'B', 'BC', 'C', 'D', 'E')),
    score DECIMAL(5,2),
    attendance_percentage DECIMAL(5,2),
    remarks TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE(student_id, course_id, academic_year, semester)
);

CREATE INDEX idx_enrollments_student_id ON enrollments(student_id);
CREATE INDEX idx_enrollments_course_id ON enrollments(course_id);
CREATE INDEX idx_enrollments_semester ON enrollments(semester);
CREATE INDEX idx_enrollments_status ON enrollments(status);

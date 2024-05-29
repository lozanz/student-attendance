-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE attendance (
  id SERIAL PRIMARY KEY ,
  student_id INT REFERENCES student(id),
  date DATE NOT NULL,
  status VARCHAR(50) NOT NULL CHECK (status IN ('hadir', 'absen', 'terlambat', 'izin', 'sakit'))

)

-- +migrate StatementEnd
-- FindAllExams :many
SELECT * FROM exam_library.exams;

-- FindAllExamQuestions :many
SELECT * FROM exam_library.exam_questions;

-- FindExamByID :one
SELECT * FROM exam_library.exams WHERE exam_id = $1;

-- FindExamQuestionByID :one
SELECT * FROM exam_library.exam_questions WHERE exam_question_id = $1;
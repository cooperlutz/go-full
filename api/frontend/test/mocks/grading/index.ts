import { http, HttpResponse } from "msw";
import { type Exam, type Question } from "~/grading/services";

const gradingHandlers = [
  http.get("/api/grading/v1/exams/ungraded", () => {
    const exams: Array<Exam> = [
      {
        examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
        totalPointsPossible: 100,
        totalPointsEarned: 85,
        gradingCompleted: false,
        questions: [
          {
            questionId: "1",
            examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
            questionIndex: 0,
            questionType: "short-answer",
            graded: false,
            pointsPossible: 50,
          },
          {
            questionId: "2",
            examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
            questionIndex: 1,
            questionType: "short-answer",
            pointsPossible: 50,
            pointsEarned: 45,
            graded: false,
          },
        ],
      },
    ];
    return HttpResponse.json(exams);
  }),

  http.get("/api/grading/v1/exams/:examId/questions/:questionIndex", () => {
    const question: Question = {
      questionId: "1",
      examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
      questionIndex: 1,
      questionType: "short-answer",
      feedback: "What is Go?",
      graded: false,
      pointsPossible: 50,
      providedAnswer: "a command to start something",
      pointsEarned: 0,
    };
    return HttpResponse.json(question);
  }),

  http.post(
    "/api/grading/v1/exams/:examId/questions/:questionIndex/grade",
    () => {
      return HttpResponse.json();
    },
  ),

  http.get("/api/grading/v1/exams/:examId", () => {
    const exams: Exam = {
      examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
      totalPointsPossible: 100,
      totalPointsEarned: 85,
      gradingCompleted: false,
      questions: [
        {
          questionId: "1",
          examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
          questionIndex: 0,
          questionType: "short-answer",
          graded: false,
          pointsPossible: 50,
        },
        {
          questionId: "2",
          examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
          questionIndex: 1,
          questionType: "short-answer",
          pointsPossible: 50,
          pointsEarned: 45,
          graded: false,
        },
      ],
    };
    return HttpResponse.json(exams);
  }),
];

export const gradingMockHandlers = [...gradingHandlers];

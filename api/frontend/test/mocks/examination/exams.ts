import { http, HttpResponse } from "msw";

const examinationHandlers = [
  http.get("/api/examination/v1/exams", () => {
    return HttpResponse.json([
      {
        examId: "5d9abb80-0706-42ad-8131-33627d3e6b17",
        studentId: "123e4567-e89b-12d3-a456-426614174000",
      },
      {
        examId: "f660452b-4075-4eac-b87a-a5b1ce7bd428",
        studentId: "123e4567-e89b-12d3-a456-426614174000",
      },
    ]);
  }),

  http.post("/api/examination/v1/exams", async ({ request }) => {
    const newExam = await request.json();
    return HttpResponse.json({
      examId: crypto.randomUUID(),
      studentId: (newExam as { studentId: string }).studentId || "",
      examLibraryId: (newExam as { examLibraryId: string }).examLibraryId || "",
    });
  }),

  http.get("/api/examination/v1/exams/:examId", (req) => {
    const { examId } = req.params;
    return HttpResponse.json({
      examId,
      studentId: "123e4567-e89b-12d3-a456-426614174000",
      examLibraryId: "f660452b-4075-4eac-b87a-a5b1ce7bd428",
      questions: [
        {
          questionIndex: 1,
          answered: true,
        },
        {
          questionIndex: 2,
          answered: false,
        },
      ],
    });
  }),

  http.get("/api/examination/v1/exams/:examId/progress", () => {
    return HttpResponse.json({
      answeredQuestions: 23,
      totalQuestions: 28,
    });
  }),

  http.get("/api/examination/v1/exams/:examId/questions/:index", (req) => {
    const { examId, index } = req.params;
    if (index === "1") {
      return HttpResponse.json({
        examId,
        questionIndex: index,
        answered: false,
        questionText: `What is going on in question ${index}?`,
        questionType: "multiple-choice",
        providedAnswer: null,
        responseOptions: ["Option A", "Option B", "Option C", "Option D"],
      });
    }
    if (index === "2") {
      return HttpResponse.json({
        examId,
        questionIndex: index,
        answered: false,
        questionText: `Solve the equation: 2x + 3 = 7.`,
        questionType: "short-answer",
        providedAnswer: null,
        responseOptions: null,
      });
    }
    if (index === "3") {
      return HttpResponse.json({
        examId,
        questionIndex: index,
        answered: false,
        questionText: `Explain the theory of relativity.`,
        questionType: "essay",
        providedAnswer: null,
        responseOptions: null,
      });
    }
  }),

  http.post(
    "/api/examination/v1/exams/:examId/questions/:index",
    async ({ request, params }) => {
      const answer = await request.json();
      const { examId, index } = params;
      return HttpResponse.json({
        examId,
        questionIndex: index,
        providedAnswer: (answer as { providedAnswer: string }).providedAnswer,
      });
    },
  ),
];

export default examinationHandlers;

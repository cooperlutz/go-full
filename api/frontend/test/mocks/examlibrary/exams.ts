import { http, HttpResponse } from "msw";

const examLibraryHandlers = [
  http.get("/api/examlibrary/v1/exams", () => {
    return HttpResponse.json([
      {
        gradeLevel: 3,
        id: "5d9abb80-0706-42ad-8131-33627d3e6b17",
        name: "Midterm Exam",
      },
      {
        gradeLevel: 4,
        id: "73a2acee-9266-410c-97f9-cd7a8f24e7a9",
        name: "Midterm Exam",
      },
      {
        gradeLevel: 5,
        id: "8b0a98c0-3591-440f-8c65-408e33deb000",
        name: "Midterm Exam",
      },
      {
        gradeLevel: 6,
        id: "49ce4c63-d8db-4f32-957b-37afecd873fe",
        name: "Sample Exam",
      },
      {
        gradeLevel: 7,
        id: "0ac39757-38da-43cf-8b84-4bd5e0fc8f32",
        name: "Sample Exam",
      },
      {
        gradeLevel: 8,
        id: "cb19ecd4-2952-416f-bcac-6cc059c0f2d3",
        name: "Sample Exam",
      },
      {
        gradeLevel: 9,
        id: "60b7aae6-1710-40d1-b405-8284e47cf164",
        name: "Sample Exam",
      },
      {
        gradeLevel: 10,
        id: "e3e5ce2f-f098-46b7-8bf6-46aa15f10d6b",
        name: "Sample Exam",
      },
      {
        gradeLevel: 11,
        id: "e0eb0311-f318-4640-8a93-a76d703f4e48",
        name: "Sample Exam",
      },
      {
        gradeLevel: 12,
        id: "eabfdc0f-465c-4344-bdd9-db68119fce5f",
        name: "Sample Exam",
      },
    ]);
  }),

  http.post("/api/examlibrary/v1/exams", async ({ request }) => {
    const reqBody = (await request.json()) as {
      name: string;
      gradeLevel: number;
      questions: Array<{
        index: number;
        questionText: string;
        questionType: string;
        possiblePoints: number;
        possibleAnswers?: string[];
      }>;
    } | null;

    return HttpResponse.json({
      gradeLevel: reqBody?.gradeLevel || 0,
      id: crypto.randomUUID(),
      name: reqBody?.name || "Untitled Exam",
      questions: reqBody?.questions || [],
    });
  }),

  http.get("/api/examlibrary/v1/exams/:id", (req) => {
    const { id } = req.params;
    return HttpResponse.json({
      id,
      name: "Sample Exam",
      gradeLevel: 6,
    });
  }),
];

export default examLibraryHandlers;

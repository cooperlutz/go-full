import { http, HttpResponse } from 'msw'

const examinationHandlers = [

  http.get('/examination/api/v1/exams', () => {
    return HttpResponse.json([
  {
    "examId": "5d9abb80-0706-42ad-8131-33627d3e6b17",
    "studentId": "123e4567-e89b-12d3-a456-426614174000",
  },
  {
    "examId": "f660452b-4075-4eac-b87a-a5b1ce7bd428",
    "studentId": "123e4567-e89b-12d3-a456-426614174000",
  }
]
    )
  }),

  http.post('/examination/api/v1/exams', async ({ request }) => {
    const reqBody = await request.json() as { 
      examId: string; 
      studentId: string;
      } | null;

    return HttpResponse.json({
      "examId": reqBody?.examId || "",
      "studentId": reqBody?.studentId || ""
    })
  }),
]

export default examinationHandlers;
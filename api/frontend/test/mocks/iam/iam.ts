import { http, HttpResponse } from "msw";

const iamHandlers = [
  http.get("/api/iam/profile", () => {
    return HttpResponse.json({
      id: "1f23abc456def7890ghi",
      email: "email@example.com",
    });
  }),
];

export default iamHandlers;

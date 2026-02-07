import { http, HttpResponse } from "msw";

const authHandlers = [
  http.post("/auth/register", () => {
    return HttpResponse.json({
      id: "1f23abc456def7890ghi",
      email: "email@example.com",
    });
  }),

  http.post("/auth/login", () => {
    return HttpResponse.json({
      accessToken: "mocked-access-token",
      refreshToken: "mocked-refresh-token",
    });
  }),

  http.post("/auth/refresh", () => {
    return HttpResponse.json({
      accessToken: "mocked-access-token",
    });
  }),
];

export default authHandlers;

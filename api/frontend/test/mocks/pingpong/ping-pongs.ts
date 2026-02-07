import { http, HttpResponse } from "msw";

const pingPongHandlers = [
  http.get("/api/pingpong/v1/ping-pongs", () => {
    return HttpResponse.json({
      pingpongs: [
        {
          createdAt: "2025-10-30T13:19:44.127908Z",
          deleted: false,
          deletedAt: null,
          id: "f660452b-4075-4eac-b87a-a5b1ce7bd428",
          message: "pong",
          updatedAt: "0001-01-01T00:00:00Z",
        },
      ],
    });
  }),

  http.post("/api/pingpong/v1/ping-pongs", async ({ request }) => {
    const reqBody = (await request.json()) as { message: string } | null;

    if (reqBody?.message === "pong") {
      return HttpResponse.json({
        createdAt: "2025-10-30T13:19:44.127908Z",
        deleted: false,
        deletedAt: null,
        id: "f660452b-4075-4eac-b87a-a5b1ce7bd428",
        message: "Ping!",
        updatedAt: "0001-01-01T00:00:00Z",
      });
    }

    if (reqBody?.message === "ping") {
      return HttpResponse.json({
        createdAt: "2025-10-30T13:19:44.127908Z",
        deleted: false,
        deletedAt: null,
        id: "f660452b-4075-4eac-b87a-a5b1ce7bd428",
        message: "Pong!",
        updatedAt: "0001-01-01T00:00:00Z",
      });
    }
  }),

  /* STEP 5.1. Implement Frontend Mock Endpoints
here, we create a mock response for hitting the endpoint for a particular id
*/
  http.get("/api/pingpong/v1/ping-pongs/:id", (req) => {
    const { id } = req.params;
    if (id === "f660452b-4075-4eac-b87a-a5b1ce7bd428") {
      return HttpResponse.json({
        createdAt: "2025-10-30T13:19:44.127908Z",
        deleted: false,
        deletedAt: null,
        id: "f660452b-4075-4eac-b87a-a5b1ce7bd428",
        message: "pong",
        updatedAt: "0001-01-01T00:00:00Z",
      });
    } else {
      return HttpResponse.json({ error: "Not Found" }, { status: 404 });
    }
  }),
];

export default pingPongHandlers;

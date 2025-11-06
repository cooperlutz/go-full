import { http, HttpResponse } from 'msw'

const pingPongHandlers = [

  http.get('/pingpong/api/v1/ping-pongs', () => {
    return HttpResponse.json({
      "pingpongs": [
        {
          "createdAt": "2025-10-30T13:19:44.127908Z",
          "deleted": false,
          "deletedAt": null,
          "id": "f660452b-4075-4eac-b87a-a5b1ce7bd428",
          "message": "pong",
          "updatedAt": "0001-01-01T00:00:00Z"
        }
      ]
    })
  }),

  http.post('/pingpong/api/v1/ping-pongs', async ({ request }) => {
    const reqBody = await request.json() as { message: string } | null;

    if (reqBody?.message === 'pong') {
      return HttpResponse.json({
        "createdAt": "2025-10-30T13:19:44.127908Z",
        "deleted": false,
        "deletedAt": null,
        "id": "f660452b-4075-4eac-b87a-a5b1ce7bd428",
        "message": "Ping!",
        "updatedAt": "0001-01-01T00:00:00Z"
      })
    }

    if (reqBody?.message === 'ping') {
    return HttpResponse.json({
      "createdAt": "2025-10-30T13:19:44.127908Z",
      "deleted": false,
      "deletedAt": null,
      "id": "f660452b-4075-4eac-b87a-a5b1ce7bd428",
      "message": "Pong!",
      "updatedAt": "0001-01-01T00:00:00Z"
    })
    } 
  })
]

export default pingPongHandlers;
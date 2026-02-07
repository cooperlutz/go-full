import { beforeAll, afterEach, afterAll } from "vitest";
import { server } from "./test/mocks/node";
import "./test/mocks/localStorage";

beforeAll(() => {
  server.listen();
});

afterEach(() => {
  server.resetHandlers();
});

afterAll(() => {
  server.close();
});

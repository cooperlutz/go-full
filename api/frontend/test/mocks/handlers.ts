import { pingPongMockHandlers } from "./pingpong";
import { examLibraryMockHandlers } from "./examlibrary";
import { examinationMockHandlers } from "./examination";

export const handlers = [
  ...pingPongMockHandlers,
  ...examLibraryMockHandlers,
  ...examinationMockHandlers,
];

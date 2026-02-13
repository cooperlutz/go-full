import { pingPongMockHandlers } from "./pingpong";
import { examLibraryMockHandlers } from "./examlibrary";
import { examinationMockHandlers } from "./examination";
import { iamMockHandlers } from "./iam";
import { gradingMockHandlers } from "./grading";

export const handlers = [
  ...pingPongMockHandlers,
  ...examLibraryMockHandlers,
  ...examinationMockHandlers,
  ...iamMockHandlers,
  ...gradingMockHandlers,
];

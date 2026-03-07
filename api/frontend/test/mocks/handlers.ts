import { pingPongMockHandlers } from "./pingpong";
import { examLibraryMockHandlers } from "./examlibrary";
import { examinationMockHandlers } from "./examination";
import { iamMockHandlers } from "./iam";
import { gradingMockHandlers } from "./grading";
import { reportingMockHandlers } from "./reporting";

export const handlers = [
  ...pingPongMockHandlers,
  ...examLibraryMockHandlers,
  ...examinationMockHandlers,
  ...iamMockHandlers,
  ...gradingMockHandlers,
  ...reportingMockHandlers,
];

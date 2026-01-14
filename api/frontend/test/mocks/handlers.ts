import { pingPongMockHandlers } from './pingpong'
import { examLibraryMockHandlers } from './examlibrary';

export const handlers = [
  ...pingPongMockHandlers,
  ...examLibraryMockHandlers,
]

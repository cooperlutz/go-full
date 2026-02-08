import authHandlers from "./auth";
import iamHandlers from "./iam";

export const iamMockHandlers = [...authHandlers, ...iamHandlers];

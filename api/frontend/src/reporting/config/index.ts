import { authRefreshMiddleware } from "~/app/utils/middleware";

import { Configuration } from "../services/runtime";

export const BackendConfig = new Configuration({
  basePath: "/api/reporting",
  credentials: "include",
  middleware: [authRefreshMiddleware],
});

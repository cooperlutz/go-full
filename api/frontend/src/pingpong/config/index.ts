import { authRefreshMiddleware } from "~/app/utils/middleware";

import { Configuration } from "~/pingpong/services/runtime";

export const BackendConfig = new Configuration({
  basePath: "/api/pingpong/v1",
  credentials: "include",
  middleware: [authRefreshMiddleware],
});

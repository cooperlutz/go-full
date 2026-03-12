import { getAuthorizationHeader } from "~/app/utils/authHeader";
import { authRefreshMiddleware } from "~/app/utils/middleware";

import { Configuration } from "../services/runtime";

export const BackendConfig = new Configuration({
  basePath: "/api/retail-sales",
  headers: {
    Authorization: getAuthorizationHeader,
  },
  middleware: [authRefreshMiddleware],
});

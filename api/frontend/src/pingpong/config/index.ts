import { Configuration } from "~/pingpong/services/runtime";
import { getAuthorizationHeader } from "~/app/utils/authHeader";
import { authRefreshMiddleware } from "~/app/utils/middleware";

export const BackendConfig = new Configuration({
  basePath: "/api/pingpong/v1",
  headers: {
    Authorization: getAuthorizationHeader,
  },
  middleware: [authRefreshMiddleware],
});

import { Configuration } from "~/examlibrary/services/runtime";
import { getAuthorizationHeader } from "~/app/utils/authHeader";
import { authRefreshMiddleware } from "~/app/utils/middleware";

export const BackendConfig = new Configuration({
  basePath: "/api/examlibrary/v1",
  headers: {
    Authorization: getAuthorizationHeader,
  },
  middleware: [authRefreshMiddleware],
});

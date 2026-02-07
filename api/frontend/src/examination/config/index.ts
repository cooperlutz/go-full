import { Configuration } from "~/examination/services/runtime";
import { getAuthorizationHeader } from "~/app/utils/authHeader";
import { authRefreshMiddleware } from "~/app/utils/middleware";

export const BackendConfig = new Configuration({
  basePath: "/api/examination",
  headers: {
    Authorization: getAuthorizationHeader,
  },
  middleware: [authRefreshMiddleware],
});

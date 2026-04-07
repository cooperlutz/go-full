import { BackendConfig } from "~/iam/config";
import { DefaultApi } from "~/iam/services";

type FetchAPI = WindowOrWorkerGlobalScope["fetch"];

interface FetchParams {
  url: string;
  init: RequestInit;
}

interface ResponseContext {
  fetch: FetchAPI;
  url: string;
  init: RequestInit;
  response: Response;
}

interface RequestContext {
  fetch: FetchAPI;
  url: string;
  init: RequestInit;
}

interface ErrorContext {
  fetch: FetchAPI;
  url: string;
  init: RequestInit;
  error: unknown;
  response?: Response;
}

interface Middleware {
  pre?(context: RequestContext): Promise<FetchParams | void>;
  post?(context: ResponseContext): Promise<Response | void>;
  onError?(context: ErrorContext): Promise<Response | void>;
}

export const authRefreshMiddleware: Middleware = {
  post: async (context: ResponseContext) => {
    if (context.response.status === 401) {
      try {
        const api = new DefaultApi(BackendConfig);
        await api.refreshToken();
        // Retry the original request — cookies are sent automatically
        const newResponse = await context.fetch(context.url, context.init);
        return newResponse;
      } catch {
        // Refresh failed; redirect to login
        window.location.href = "/login";
      }
    }
  },
};

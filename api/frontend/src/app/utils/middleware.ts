import { useRefreshToken } from "~/iam/composables/useIam";
import { useLocalTokenStore } from "~/iam/stores/useToken";

const { refreshToken } = useRefreshToken();
const tokenStore = useLocalTokenStore();

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
        await refreshToken();
        // retry the original request after refreshing the token
        const newInit = { ...context.init };
        newInit.headers = {
          ...newInit.headers,
          Authorization: `Bearer ${tokenStore.getAccessToken()}`,
        };
        console.log("Retrying request with new access token...");
        try {
          const newResponse = await context.fetch(context.url, newInit);
          return newResponse;
        } catch (error) {
          console.error(
            "Failed to retry request after refreshing token:",
            error,
          );
          tokenStore.clear();
        }
      } catch (error) {
        console.error("Failed to refresh token:", error);
        tokenStore.clear();
      }
    }
  },
};

import { ref } from "vue";

import { useLocalTokenStore } from "~/iam/stores/useToken";
import { BackendConfig } from "~/iam/config";
import {
  DefaultApi,
  type LoginUserRequest,
  type RefreshTokenRequest,
  type RegisterUserRequest,
} from "~/iam/services";

const iamAPI = new DefaultApi(BackendConfig);
const tokenStore = useLocalTokenStore();

export function useLoginUser() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const loginUser = async (email: string, password: string) => {
    loading.value = true;
    error.value = null;
    try {
      const loginReq: LoginUserRequest = {
        loginRequest: {
          email,
          password,
        },
      };
      const loggedInUser = await iamAPI.loginUser(loginReq);
      tokenStore.setAccessToken(loggedInUser.accessToken);
      tokenStore.setRefreshToken(loggedInUser.refreshToken);
      return loggedInUser;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    loginUser,
    error,
    loading,
  };
}

export function useRefreshToken() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const refreshToken = async () => {
    const refreshToken = tokenStore.getRefreshToken();
    loading.value = true;
    error.value = null;

    try {
      const refreshReq: RefreshTokenRequest = {
        refreshRequest: {
          refreshToken,
        },
      };
      const refreshedToken = await iamAPI.refreshToken(refreshReq);
      tokenStore.setAccessToken(refreshedToken.token);
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    refreshToken,
    error,
    loading,
  };
}

export function useRegister() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const registerUser = async (email: string, password: string) => {
    loading.value = true;
    error.value = null;
    try {
      const registerReq: RegisterUserRequest = {
        registerRequest: {
          email,
          password,
        },
      };
      const registeredUser = await iamAPI.registerUser(registerReq);
      return registeredUser;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    registerUser,
    error,
    loading,
  };
}

export function useLogout() {
  const logout = () => {
    tokenStore.clear();
  };

  return {
    logout,
  };
}

export function useProfile() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const getProfile = async () => {
    loading.value = true;
    error.value = null;

    try {
      const userProfile = await iamAPI.getUserProfile();
      return userProfile;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    getProfile,
    error,
    loading,
  };
}

import { ref } from 'vue'

import { BackendConfig } from '../config'
import { DefaultApi, type LoginUserRequest, type RegisterUserRequest } from '../services'
import { useAuthState } from '../stores/useAuthState'

const iamAPI = new DefaultApi(BackendConfig)
const { setAuthenticated, clearAuthenticated } = useAuthState()

export function useLoginUser() {
  const error = ref<Error | null>(null)
  const loading = ref(false)

  const loginUser = async (email: string, password: string) => {
    loading.value = true
    error.value = null
    try {
      const loginReq: LoginUserRequest = {
        loginRequest: {
          email,
          password,
        },
      }
      await iamAPI.loginUser(loginReq)
      setAuthenticated()
    } catch (err) {
      if (err instanceof Error) {
        error.value = err
      }
    } finally {
      loading.value = false
    }
  }

  return {
    loginUser,
    error,
    loading,
  }
}

export function useRefreshToken() {
  const error = ref<Error | null>(null)
  const loading = ref(false)

  const refreshToken = async () => {
    loading.value = true
    error.value = null

    try {
      await iamAPI.refreshToken()
    } catch (err) {
      if (err instanceof Error) {
        error.value = err
      }
    } finally {
      loading.value = false
    }
  }

  return {
    refreshToken,
    error,
    loading,
  }
}

export function useRegister() {
  const error = ref<Error | null>(null)
  const loading = ref(false)

  const registerUser = async (email: string, password: string) => {
    loading.value = true
    error.value = null
    try {
      const registerReq: RegisterUserRequest = {
        registerRequest: {
          email,
          password,
        },
      }
      const registeredUser = await iamAPI.registerUser(registerReq)
      return registeredUser
    } catch (err) {
      if (err instanceof Error) {
        error.value = err
      }
    } finally {
      loading.value = false
    }
  }

  return {
    registerUser,
    error,
    loading,
  }
}

export function useLogout() {
  const logout = async () => {
    await iamAPI.logoutUser()
    clearAuthenticated()
  }

  return {
    logout,
  }
}

export function useProfile() {
  const error = ref<Error | null>(null)
  const loading = ref(false)

  const getProfile = async () => {
    loading.value = true
    error.value = null

    try {
      const userProfile = await iamAPI.getUserProfile()
      return userProfile
    } catch (err) {
      if (err instanceof Error) {
        error.value = err
      }
    } finally {
      loading.value = false
    }
  }

  return {
    getProfile,
    error,
    loading,
  }
}

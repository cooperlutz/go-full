<script lang="ts" setup>
import { ref } from "vue";

import { useLoginUser } from "~/iam/composables/useIam";

const email = ref("");
const password = ref("");
const { loginUser, error, loading } = useLoginUser();

const login = async () => {
  await loginUser(email.value, password.value);
  password.value = ""; // Clear password after login attempt
  if (!error.value) {
    window.location.href = "/dashboard";
  }
};
</script>

<template>
  <section class="bg-info">
    <div
      class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0"
    >
      <a class="flex items-center mb-6 text-2xl font-semibold">
        <img class="w-8 h-8 mr-2" src="/logo.png" alt="logo" />
        Go Full
      </a>
      <div class="card bg-base-100 shadow-sm">
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1 class="text-xl font-bold leading-tight tracking-tight">
            Login to your account
          </h1>
          <form class="space-y-4 md:space-y-6" action="#">
            <div>
              <label for="email" class="block mb-2 text-sm">email</label>
              <input
                class="input validator"
                type="email"
                v-model="email"
                required
                placeholder="mail@site.com"
                id="login-input-email"
              />
            </div>
            <div>
              <label for="password" class="block mb-2 text-sm">password</label>
              <input
                type="password"
                v-model="password"
                class="input validator"
                required
                placeholder="Password"
                minlength="8"
                pattern="(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}"
                title="Must be more than 8 characters, including number, lowercase letter, uppercase letter"
                id="login-input-password"
              />
            </div>
            <div class="flex items-center justify-between">
              <a
                href="#"
                class="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500"
                >Forgot password?</a
              >
            </div>
            <button
              class="w-full btn btn-primary"
              @click="login"
              :disabled="loading"
              id="login-button"
            >
              Login
            </button>
            <div v-if="error" style="color: red">{{ error.message }}</div>
            <p class="text-sm font-light text-gray-500 dark:text-gray-400">
              Don't have an account yet?
              <a
                href="/register"
                class="font-medium text-primary-600 hover:underline dark:text-primary-500"
                id="login-register-link"
                >Register here</a
              >
            </p>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>

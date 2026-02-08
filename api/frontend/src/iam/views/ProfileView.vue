<script setup lang="ts">
import { onMounted, ref } from "vue";
import { Mail, School } from "lucide-vue-next";

import { type UserProfile } from "~/iam/services";
import { useProfile } from "~/iam/composables/useIam";

const userProfile = ref<UserProfile | null>(null);

onMounted(async () => {
  const { getProfile } = useProfile();
  const profile = await getProfile();
  if (profile) {
    userProfile.value = profile;
  }
});
</script>

<template>
  <section>
    <div class="mx-auto max-w-5xl px-4 2xl:px-0">
      <div class="py-4 md:py-8">
        <div class="mb-4 grid gap-4 sm:grid-cols-2 sm:gap-8 lg:gap-16">
          <div class="space-y-4">
            <div class="flex space-x-4">
              <div class="skeleton h-16 w-16"></div>
              <div>
                <span
                  class="mb-2 inline-block rounded px-2.5 py-0.5 text-xs"
                  id="user-id"
                >
                  {{ userProfile?.id }}
                </span>
                <h2 class="flex items-center text-xl font-bold leading-none">
                  TODO: Name
                </h2>
              </div>
            </div>
            <dl class="space-y-2">
              <dt class="flex font-semibold">
                <Mail />
                <span class="ml-2">Email</span>
              </dt>
              <dd id="user-email">
                {{ userProfile?.email }}
              </dd>
            </dl>
          </div>
          <div class="space-y-4">
            <dl class="space-y-2">
              <dt class="flex font-semibold">
                <School />
                <span class="ml-2">School</span>
              </dt>
              <dd class="flex items-center gap-1">TODO: School name</dd>
            </dl>
            <dl class="space-y-2">
              <dt class="font-semibold">Grade Level</dt>
              <dd>TODO: 3</dd>
            </dl>
          </div>
        </div>
        <button class="btn">Edit your profile</button>
      </div>
      <div class="skeleton h-32 w-full justify-center flex items-center">
        TODO: Prior exams list goes here
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { Check } from "lucide-vue-next";

import { type Question } from "../services";

defineProps<{
  examId: string;
  questions: Array<Question>;
}>();
</script>

<template>
  <div
    class="card card-border border-secondary border-solid bg-base-100 shadow-lg"
  >
    <div class="overflow-y-auto h-96">
      <!-- menu -->
      <ul
        v-for="question in questions"
        :key="question.questionIndex"
        class="menu rounded-box w-full"
      >
        <li>
          <a
            :href="`/exam/${examId}/question/${question.questionIndex}`"
            :id="`question-nav-item-${question.questionIndex}`"
          >
            Question {{ question.questionIndex?.toString() }}
            <span v-if="question.answered" class="badge badge-outline">
              <Check />
            </span>
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

import RecordAnswerButton from "./RecordAnswerButton.vue";
import { type Question } from "../../services";

const props = defineProps<{
  question: Question;
}>();

const providedAnswer = ref<string>("");
providedAnswer.value = props.question.providedAnswer || "";
</script>

<template>
  <div>
    <div class="badge badge-outline badge-primary">Short Answer</div>
    <h2 class="card-title mt-4">{{ props.question.questionText }}</h2>
    <div class="divider"></div>
    <p>{{ props.question.questionText }}</p>
    <label class="cursor-pointer flex items-center space-x-2 mt-6 space-y-4">
      <input
        type="text"
        class="input"
        v-model="providedAnswer"
        id="short-answer-input"
      />
    </label>
    <RecordAnswerButton
      :examId="props.question.examId"
      :questionIndex="props.question.questionIndex"
      :answer="providedAnswer"
    />
  </div>
</template>

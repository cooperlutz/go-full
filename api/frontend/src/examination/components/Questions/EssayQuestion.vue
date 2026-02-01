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
    <h2 class="card-title">Essay Question</h2>
    <div class="divider"></div>
    <p class="mt-6">{{ props.question.questionText }}</p>
    <ul class="mt-6 space-y-4">
      <label class="cursor-pointer flex items-center space-x-2">
        <textarea
          class="textarea"
          v-model="providedAnswer"
          id="essay-question-input"
        />
      </label>
    </ul>
    <RecordAnswerButton
      :examId="props.question.examId"
      :questionIndex="props.question.questionIndex"
      :answer="providedAnswer"
    />
  </div>
</template>

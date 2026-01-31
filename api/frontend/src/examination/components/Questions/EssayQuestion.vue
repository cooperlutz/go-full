<script setup lang="ts">
import { ref } from "vue";

import SubmitAnswerButton from "../SubmitAnswerButton.vue";
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
    <p>{{ props.question.questionText }}</p>
    <ul>
      <label class="cursor-pointer flex items-center space-x-2">
        <textarea
          class="textarea"
          v-model="providedAnswer"
          id="essay-question-input"
        />
      </label>
    </ul>
    <SubmitAnswerButton
      :examId="props.question.examId"
      :questionIndex="props.question.questionIndex"
      :answer="providedAnswer"
    />
  </div>
</template>

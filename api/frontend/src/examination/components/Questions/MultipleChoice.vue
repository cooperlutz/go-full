<script setup lang="ts">
import { ref } from "vue";

import RecordAnswerButton from "./RecordAnswerButton.vue";
import { type Question } from "../../services";

const props = defineProps<{
  question: Question;
}>();

const selectedAnswer = ref<string>("");
selectedAnswer.value = props.question.providedAnswer || "";
</script>

<template>
  <div id="multiple-choice-question">
    <div class="badge badge-outline badge-primary">Multiple Choice</div>
    <h2 class="card-title mt-4">{{ props.question.questionText }}</h2>
    <div class="divider"></div>
    <ul class="mt-6 space-y-4">
      <li v-for="option in props.question.responseOptions" :key="option">
        <label class="cursor-pointer flex items-center space-x-2">
          <input
            :id="`multiple-choice-radio-option-${props.question.responseOptions?.indexOf(option)}`"
            type="radio"
            :value="option"
            :name="option"
            class="radio"
            v-model="selectedAnswer"
          />
          <span>{{ option }}</span>
        </label>
      </li>
    </ul>
    <RecordAnswerButton
      :examId="props.question.examId"
      :questionIndex="props.question.questionIndex"
      :answer="selectedAnswer"
    />
  </div>
</template>

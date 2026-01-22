<script setup lang="ts">
import { ref } from "vue";

import { useExamination } from "~/examination/composables/useExamination";

const props = defineProps<{
  examId: string;
}>();
const studentId = ref("");
const { startExam, exam, error } = useExamination();

// clickStart starts a new exam with the given examId and studentId
const clickStart = async (examId: string, studentId: string) => {
  await startExam(examId, studentId);
  if (error.value) {
    console.error("Error starting exam:", error.value);
    return;
  }
  window.location.href = `/exam/${exam.value?.examId}`;
};
</script>

<template>
  <button
    id="start_exam_button"
    class="btn"
    onclick="start_exam_modal.showModal()"
  >
    Start Exam
  </button>
  <dialog id="start_exam_modal" class="modal">
    <div class="modal-box">
      <h3 class="text-lg font-bold">Start Exam</h3>
      <label class="input">
        <input
          id="student-id-input"
          type="text"
          class="grow"
          placeholder="Student ID"
          v-model="studentId"
        />
      </label>
      <div class="card-actions">
        <div
          id="start-button"
          class="btn btn-m text-xs"
          @click="clickStart(props.examId, studentId)"
        >
          Start Exam
        </div>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

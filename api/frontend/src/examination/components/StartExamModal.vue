<script setup lang="ts">
import { ref } from "vue";

import { useStartExam } from "../composables/useExamination";

const props = defineProps<{
  libraryExamId: string;
}>();
const studentId = ref("");
const { startExam, exam, error } = useStartExam();

// clickStart starts a new exam with the given examId and studentId
const clickStart = async (libraryExamId: string, studentId: string) => {
  await startExam(libraryExamId, studentId);
  if (error.value) {
    console.error("Error starting exam:", error.value);
    return;
  }
  window.location.href = `/exam/${exam.value?.examId}`;
};
</script>

<template>
  <button
    id="start-exam-modal-button"
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
          id="confirm-start-exam-button"
          class="btn btn-m text-xs"
          @click="clickStart(props.libraryExamId, studentId)"
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

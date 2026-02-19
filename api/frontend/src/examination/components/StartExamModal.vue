<script setup lang="ts">
import { ref } from "vue";

import { useStartExam } from "../composables/useExamination";

const props = defineProps<{
  libraryExamId: string;
}>();
const emit = defineEmits(["exam-started"]);

const studentId = ref("");
const { startExam, exam, error } = useStartExam();

// clickStart starts a new exam with the given examId and studentId
const clickStart = async (libraryExamId: string, studentId: string) => {
  await startExam(libraryExamId, studentId);
  if (error.value) {
    console.error("Error starting exam:", error.value);
    return;
  }
  emit("exam-started");
  window.location.href = `/exam/${exam.value?.examId}`;
};
</script>

<template>
  <button
    id="start-exam-modal-button"
    class="btn btn-info"
    onclick="start_exam_modal.showModal()"
  >
    Take Exam
  </button>
  <dialog id="start_exam_modal" class="modal">
    <div class="modal-box">
      <h3 class="text-lg font-bold">Take Exam</h3>
      Enter your Student ID to begin the exam.
      <label class="input mt-4">
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
          class="btn btn-m btn-success text-xs mt-4"
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

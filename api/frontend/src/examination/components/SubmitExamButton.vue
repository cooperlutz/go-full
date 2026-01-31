<script setup lang="ts">
import { useSubmitExam } from "../composables/useExamination";

defineProps<{
  examId: string;
}>();

const { submitExam } = useSubmitExam();

function submitExamHandler(examId: string) {
  submitExam(examId);
}
</script>

<template>
  <button
    class="btn"
    onclick="valiation_modal.showModal()"
    id="exam-submission-button"
  >
    Submit Exam
  </button>
  <dialog id="valiation_modal" class="modal">
    <div class="modal-box">
      Are you sure you want to submit your exam?
      <button
        class="btn"
        @click="submitExamHandler(examId)"
        id="confirm-exam-submission-button"
      >
        Yes, I'm sure
      </button>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

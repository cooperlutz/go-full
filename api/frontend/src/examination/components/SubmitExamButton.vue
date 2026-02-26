<script setup lang="ts">
import { useRouter } from "vue-router";

import { useSubmitExam } from "../composables/useExamination";

defineProps<{
  examId: string;
}>();
const emit = defineEmits(["exam-submitted"]);

const router = useRouter();
const { submitExam } = useSubmitExam();

function submitExamHandler(examId: string) {
  submitExam(examId);
  emit("exam-submitted");
  router.push({ name: "ExamSubmitted" });
}
</script>

<template>
  <button
    class="btn w-full btn-success"
    onclick="valiation_modal.showModal()"
    id="exam-submission-button"
  >
    Submit Exam
  </button>
  <dialog id="valiation_modal" class="modal">
    <div class="modal-box">
      Are you sure you want to submit your exam?
      <button
        class="btn btn-error mt-4"
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

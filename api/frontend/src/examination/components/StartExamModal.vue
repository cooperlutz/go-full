<script setup lang="ts">
import { useProfile } from "~/iam/composables/useIam";

import { useStartExam } from "../composables/useExamination";

const props = defineProps<{
  libraryExamId: string;
}>();
const emit = defineEmits(["exam-started"]);

const { startExam, exam, error } = useStartExam();
const { getProfile } = useProfile();

// clickStart starts a new exam with the given examId and studentId
const clickStart = async (libraryExamId: string) => {
  const profile = await getProfile();
  await startExam(libraryExamId, profile?.id || "");
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
      <div class="card-actions">
        <div
          id="confirm-start-exam-button"
          class="btn btn-m btn-success text-xs mt-4"
          @click="clickStart(props.libraryExamId)"
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

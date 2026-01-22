import { ref } from "vue";

import { BackendConfig } from "~/examination/config";
import {
  DefaultApi,
  type Exam,
  type StartExam,
  type StartNewExamRequest,
} from "~/examination/services";

const examinationAPI = new DefaultApi(BackendConfig);

// Composable for sending a ping pong message
export function useExamination() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const exam = ref<Exam | null>(null);

  const startExam = async (libraryExamId: string, studentId: string) => {
    loading.value = true;
    error.value = null;
    const data: StartExam = {
      libraryExamId,
      studentId,
    };
    const req: StartNewExamRequest = {
      startExam: data,
    };
    try {
      const response = await examinationAPI.startNewExam(req);
      exam.value = response;
      return response;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    startExam,
    exam,
  };
}

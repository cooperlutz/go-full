import { ref } from "vue";

import { BackendConfig } from "../config";
import { DefaultApi, type Exam, type Question } from "../services";

const gradingAPI = new DefaultApi(BackendConfig);

export function useGetFindIncompleteExams() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const incompleteExams = ref<Array<Exam> | null>(null);

  const getFindIncompleteExams = async () => {
    loading.value = true;
    error.value = null;
    try {
      const response = await gradingAPI.getFindIncompleteExams();
      incompleteExams.value = response;
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
    incompleteExams,
    getFindIncompleteExams,
  };
}

export function useGradeExamQuestion() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const gradeExamQuestion = async (
    examId: string,
    questionIndex: number,
    pointsEarned: number,
    feedback: string,
  ) => {
    loading.value = true;
    error.value = null;
    try {
      await gradingAPI.gradeExamQuestion({
        examId,
        questionIndex,
        gradeQuestion: {
          points: pointsEarned,
          feedback,
        },
      });
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
    gradeExamQuestion,
  };
}

export function useGetExamQuestion() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const question = ref<Question | null>(null);

  const getExamQuestion = async (examId: string, questionIndex: number) => {
    loading.value = true;
    error.value = null;
    try {
      const response = await gradingAPI.getExamQuestion({
        examId,
        questionIndex,
      });
      question.value = response;
      return question.value;
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
    question,
    getExamQuestion,
  };
}

export function useGetExam() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const exam = ref<Exam | null>(null);

  const getExam = async (examId: string) => {
    loading.value = true;
    error.value = null;
    try {
      const response = await gradingAPI.getExam({
        examId,
      });
      exam.value = response;
      return exam.value;
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
    exam,
    getExam,
  };
}

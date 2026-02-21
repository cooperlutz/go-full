import { ref } from "vue";

import { BackendConfig } from "../config";
import {
  DefaultApi,
  type Exam,
  type StartExam,
  type StartNewExamRequest,
  type Question,
  type GetExamQuestionRequest,
  type GetExamRequest,
  type SubmitExamRequest,
} from "../services";

const examinationAPI = new DefaultApi(BackendConfig);

export function useStartExam() {
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

export function useSubmitExam() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const submitExam = async (examId: string) => {
    loading.value = true;
    error.value = null;
    const req: SubmitExamRequest = {
      examId,
    };
    try {
      await examinationAPI.submitExam(req);
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
    submitExam,
  };
}

export function useGetExam() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const exam = ref<Exam | null>(null);

  const getExam = async (examId: string) => {
    loading.value = true;
    error.value = null;
    const req: GetExamRequest = {
      examId,
    };
    try {
      const response = await examinationAPI.getExam(req);
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
    getExam,
    exam,
  };
}

export function useGetQuestion() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const examQuestion = ref<Question | null>(null);

  const getQuestion = async (examId: string, index: number) => {
    loading.value = true;
    error.value = null;
    const req: GetExamQuestionRequest = {
      examId,
      questionIndex: index,
    };
    try {
      const response = await examinationAPI.getExamQuestion(req);
      examQuestion.value = response;
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
    getQuestion,
    examQuestion,
  };
}

export function useSubmitAnswer() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const submitAnswer = async (
    examId: string,
    questionIndex: number,
    answer: string,
  ) => {
    loading.value = true;
    error.value = null;
    try {
      await examinationAPI.answerQuestion({
        examId,
        questionIndex,
        answer: {
          providedAnswer: answer,
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
    submitAnswer,
  };
}

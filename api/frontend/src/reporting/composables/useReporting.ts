import { ref } from "vue";

import { BackendConfig } from "../config";
import { DefaultApi, type Metric, type GetMetricRequest } from "../services";

const examinationAPI = new DefaultApi(BackendConfig);

export function useGetMetrics() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const numberOfExamsInProgress = ref<Metric | null>(null);
  const numberOfExamsCompleted = ref<Metric | null>(null);
  const numberOfExamsBeingGraded = ref<Metric | null>(null);
  const numberOfExamsGradingCompleted = ref<Metric | null>(null);

  const getMetrics = async () => {
    loading.value = true;
    error.value = null;
    try {
      const inProgressResponse = await examinationAPI.getMetric({
        metricName: "number_of_exams_in_progress",
      });
      numberOfExamsInProgress.value = inProgressResponse;

      const completedResponse = await examinationAPI.getMetric({
        metricName: "number_of_exams_completed",
      });
      numberOfExamsCompleted.value = completedResponse;

      const beingGradedResponse = await examinationAPI.getMetric({
        metricName: "number_of_exams_being_graded",
      });
      numberOfExamsBeingGraded.value = beingGradedResponse;

      const gradingCompletedResponse = await examinationAPI.getMetric({
        metricName: "number_of_exams_grading_completed",
      });
      numberOfExamsGradingCompleted.value = gradingCompletedResponse;
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
    getMetrics,
    numberOfExamsInProgress,
    numberOfExamsCompleted,
    numberOfExamsBeingGraded,
    numberOfExamsGradingCompleted,
  };
}

export function useGetMetric() {
  const error = ref<Error | null>(null);
  const loading = ref(false);
  const metric = ref<Metric | null>(null);

  const getMetric = async (metricName: string) => {
    loading.value = true;
    error.value = null;
    const req: GetMetricRequest = {
      metricName,
    };
    try {
      const response = await examinationAPI.getMetric(req);
      metric.value = response;
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
    getMetric,
    metric,
  };
}

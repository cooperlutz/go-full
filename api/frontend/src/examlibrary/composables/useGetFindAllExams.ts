import { ref } from "vue";

import { BackendConfig } from "../config";
import { ExamlibraryApi, type ExamMetadata } from "../services";

const examLibraryAPI = new ExamlibraryApi(BackendConfig);

/**
 * Composable for retrieving all exams from the exam library.
 **/
export function useGetFindAllExams() {
  const allExams = ref<ExamMetadata[] | null>(null);
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const fetchData = async () => {
    loading.value = true;
    error.value = null;
    try {
      const getAll = await examLibraryAPI.getFindAllExams();
      allExams.value = getAll;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    allExams,
    error,
    loading,
    fetchData,
  };
}

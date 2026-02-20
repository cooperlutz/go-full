import { ref } from "vue";

import { BackendConfig } from "../config";
import { ExamlibraryApi, type GetFindOneByIDRequest } from "../services";

const examLibraryAPI = new ExamlibraryApi(BackendConfig);

/**
 * Composable for fetching a single exam by ID
 **/
export function useFindExamByID() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const findExam = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      const req: GetFindOneByIDRequest = {
        examID: id,
      };
      const getByID = await examLibraryAPI.getFindOneByID(req);
      return getByID;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    findExam,
    error,
    loading,
  };
}

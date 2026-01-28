import { ref } from "vue";

import { BackendConfig } from "~/examlibrary/config";
import {
  ExamlibraryApi,
  type PostAddExamToLibraryRequest,
  type Exam,
} from "~/examlibrary/services";

const examLibraryAPI = new ExamlibraryApi(BackendConfig);

/**
 * Composable for adding a single exam to the library
 **/
export function useAddExamToLibrary() {
  const error = ref<Error | null>(null);
  const loading = ref(false);

  const addExam = async (exam: Exam) => {
    loading.value = true;
    error.value = null;
    try {
      const req: PostAddExamToLibraryRequest = {
        exam,
      };
      const addedExam = await examLibraryAPI.postAddExamToLibrary(req);
      return addedExam;
    } catch (err) {
      if (err instanceof Error) {
        error.value = err;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    addExam,
    error,
    loading,
  };
}

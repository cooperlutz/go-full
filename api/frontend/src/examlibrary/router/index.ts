import ExamLibrary from "~/examlibrary/views/ExamLibrary.vue";
import ExamOverview from "../views/ExamOverview.vue";

const examLibraryRoutes = [
  {
    path: "/exam-library",
    component: ExamLibrary,
  },
  {
    path: "/exam-library/:id",
    component: ExamOverview,
  },
];

export default examLibraryRoutes;

import ExamLibrary from "../views/ExamLibrary.vue";
import ExamOverview from "../views/ExamOverview.vue";
import ExamCreator from "../views/ExamCreator.vue";

const examLibraryRoutes = [
  {
    path: "/exam-library",
    component: ExamLibrary,
  },
  {
    path: "/exam-library/:id",
    component: ExamOverview,
    name: "ExamOverview",
  },
  {
    path: "/exam-library/creator",
    component: ExamCreator,
    name: "ExamCreator",
  },
];

export default examLibraryRoutes;

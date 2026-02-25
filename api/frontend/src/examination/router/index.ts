import ExaminationView from "../views/ExaminationView.vue";
import ExaminationOverview from "../views/ExaminationOverview.vue";
import ExamSubmitted from "../views/ExaminationSubmitted.vue";

const examinationRoutes = [
  {
    path: "/exam/:id",
    component: ExaminationOverview,
  },
  {
    path: "/exam/submitted",
    component: ExamSubmitted,
    name: "ExamSubmitted",
  },
  {
    path: "/exam/:id/question/:index",
    component: ExaminationView,
  },
];

export default examinationRoutes;

import ExaminationView from "../views/ExaminationView.vue";
import ExamSubmission from "../components/SubmitExamButton.vue";
import ExaminationOverview from "../views/ExaminationOverview.vue";

const examinationRoutes = [
  {
    path: "/exam/:id",
    component: ExaminationOverview,
  },
  {
    path: "/exam/:id/submit",
    component: ExamSubmission,
  },
  {
    path: "/exam/:id/question/:index",
    component: ExaminationView,
  },
];

export default examinationRoutes;

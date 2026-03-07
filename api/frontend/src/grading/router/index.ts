import ExamGrading from "../views/ExamGrading.vue";
import QuestionGrading from "../views/QuestionGrading.vue";
import GradingView from "../views/GradingView.vue";

const gradingRoutes = [
  {
    path: "/grading/exam/:examId/question/:questionIndex",
    component: QuestionGrading,
  },
  {
    path: "/grading/exam/:examId",
    component: ExamGrading,
  },
  {
    path: "/grading",
    component: GradingView,
    name: "GradingView",
  },
];

export default gradingRoutes;

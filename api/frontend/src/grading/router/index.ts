import ExamGrading from "../views/ExamGrading.vue";
import QuestionGrading from "../views/QuestionGrading.vue";
import ExamsToGrade from "../views/GradingView.vue";

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
    component: ExamsToGrade,
    name: "ExamsToGrade",
  },
];

export default gradingRoutes;

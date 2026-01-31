<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";

import { QuestionType, type Exam, type ExamQuestion } from "../services/models";
import { useAddExamToLibrary } from "../composables/useAddExamToLibrary";

const router = useRouter();

const addedExam = ref<Exam | null>(null);

const newExam = ref<Exam>({
  name: "",
  gradeLevel: 1,
  questions: [],
});

const possibleAnswers = ref<string[]>([]);

const newExamQuestion = ref<ExamQuestion>({
  index: 1,
  questionType: QuestionType.MultipleChoice,
  questionText: "",
  correctAnswer: "",
  possibleAnswers: possibleAnswers.value,
  possiblePoints: 1,
});

function addQuestion(questionType: QuestionType) {
  newExamQuestion.value.questionType = questionType;
  newExam.value.questions?.push({ ...newExamQuestion.value });
  // Reset the newExamQuestion for the next question
  possibleAnswers.value = [];
  newExamQuestion.value = {
    index: newExam.value.questions?.length
      ? newExam.value.questions.length + 1
      : 1,
    questionType: QuestionType.MultipleChoice,
    questionText: "",
    correctAnswer: "",
    possibleAnswers: possibleAnswers.value,
    possiblePoints: 1,
  };
}

const navigateToNewlyCreatedExam = () => {
  if (addedExam.value) {
    try {
      router.push({ name: "ExamOverview", params: { id: addedExam.value.id } });
    } catch (error) {
      console.error("Navigation error:", error);
    }
  }
};

const saveExamToLibrary = async () => {
  const { addExam } = useAddExamToLibrary();
  const ex = await addExam(newExam.value);
  if (ex) {
    addedExam.value = ex;
  } else {
    console.error("Failed to add exam to library");
  }
  navigateToNewlyCreatedExam();
};
</script>

<template>
  <div>
    Name:
    <input
      id="new-exam-name-input"
      type="text"
      placeholder="Type here"
      class="input"
      v-model="newExam.name"
    />
    Grade Level:
    <input
      id="new-exam-grade-level-input"
      type="number"
      placeholder="Type here"
      class="input"
      v-model="newExam.gradeLevel"
    />
    <button
      class="btn btn-success"
      @click="saveExamToLibrary()"
      id="save-exam-to-library-button"
    >
      Save Exam to Library
    </button>
  </div>

  <button
    class="btn"
    onclick="add_multiple_choice.showModal()"
    id="add-multiple-choice-button"
  >
    Add Multiple Choice
  </button>
  <dialog id="add_multiple_choice" class="modal">
    <div class="modal-box">
      Question:
      <input
        id="mc-question-text-input"
        type="text"
        placeholder="Type here"
        class="input"
        v-model="newExamQuestion.questionText"
      />
      Option A:
      <input
        id="mc-option-a-input"
        type="text"
        placeholder="Type here"
        class="input"
        v-model="possibleAnswers[0]"
      />
      Option B:
      <input
        id="mc-option-b-input"
        type="text"
        placeholder="Type here"
        class="input"
        v-model="possibleAnswers[1]"
      />
      Option C:
      <input
        id="mc-option-c-input"
        type="text"
        placeholder="Type here"
        class="input"
        v-model="possibleAnswers[2]"
      />
      Option D:
      <input
        id="mc-option-d-input"
        type="text"
        placeholder="Type here"
        class="input"
        v-model="possibleAnswers[3]"
      />
      Correct Answer:
      <input
        id="mc-correct-answer-input"
        type="text"
        placeholder="Type here"
        class="input"
        v-model="newExamQuestion.correctAnswer"
      />
      Point Value:
      <input
        id="mc-possible-points-input"
        type="number"
        placeholder="Type here"
        class="input"
        v-model="newExamQuestion.possiblePoints"
      />
      <div class="modal-action">
        <form method="dialog">
          <button
            class="btn"
            @click="addQuestion(QuestionType.MultipleChoice)"
            id="add-multiple-choice-confirm-button"
          >
            Add
          </button>
        </form>
      </div>
    </div>
  </dialog>

  <button
    class="btn"
    onclick="add_short_answer.showModal()"
    id="add-short-answer-button"
  >
    Add Short Answer
  </button>
  <dialog id="add_short_answer" class="modal">
    <div class="modal-box">
      Question:
      <input
        id="sa-question-text-input"
        type="text"
        placeholder="Type here"
        class="input"
        v-model="newExamQuestion.questionText"
      />
      Point Value:
      <input
        id="sa-possible-points-input"
        type="number"
        placeholder="Type here"
        class="input"
        v-model="newExamQuestion.possiblePoints"
      />
      <div class="modal-action">
        <form method="dialog">
          <button
            class="btn"
            @click="addQuestion(QuestionType.ShortAnswer)"
            id="add-short-answer-confirm-button"
          >
            Add
          </button>
        </form>
      </div>
    </div>
  </dialog>

  <button
    class="btn"
    onclick="add_essay_question.showModal()"
    id="add-essay-question-button"
  >
    Add Essay Question
  </button>
  <dialog id="add_essay_question" class="modal">
    <div class="modal-box">
      Question:
      <input
        id="essay-question-text-input"
        type="text"
        placeholder="Type here"
        class="input"
        v-model="newExamQuestion.questionText"
      />
      Point Value:
      <input
        id="essay-possible-points-input"
        type="number"
        placeholder="Type here"
        class="input"
        v-model="newExamQuestion.possiblePoints"
      />
      <div class="modal-action">
        <form method="dialog">
          <button
            class="btn"
            @click="addQuestion(QuestionType.Essay)"
            id="add-essay-question-confirm-button"
          >
            Add
          </button>
        </form>
      </div>
    </div>
  </dialog>

  <div>
    <pre v-if="!addedExam" id="new-exam-inputs">{{ newExam }}</pre>
    <div v-else>
      Exam Added to Library:
      <pre id="new-exam-output">{{ addedExam }}</pre>
    </div>
  </div>
</template>

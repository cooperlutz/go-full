# Making the Grade: Kata Overview

source: [https://www.architecturalkatas.com/kata.html?kata=MakeTheGrade.json](https://www.architecturalkatas.com/kata.html?kata=MakeTheGrade.json)

## Kata Overview

A very large and populous state would like a new system to support standardized testing across all public school systems grades 3-12.

## Users

- 40,000+ students
- 2000 graders
- 50 administrators

## Requirements

- Students will only be able to use the application within testing centers around the state, most of these will be in the schools, but not all of them

- Students should be able to take a test, and the results eventually consolidated to a single location representing all of the test scores across the state (by school, teacher, and student).

- Tests will be multiple choice, short answer, and essay.

- The system should have a reporting system to know which students have taken the tests and what score they received.

- Short answer and essay questions will be manually graded by teachers, who will then add the essay grades to the system.

## Additional Context

A change approval processes involving three different government agencies is required for changes to the way student grades are kept to ensure security; The state does not own its hosting center, but outsources it to a third party; Project must defend its budget each fiscal year.

## Assumptions

The following assumptions have been made regarding the given kata.

- Initially, the complexity of Standardized Tests will remain fairly static, with an assumed 1:1 ratio of Standardized Test and its associtated questions and content : Student Grade Level.

- A standardized Test consists of approximately 50 questions

- A standardized Test has a time limit of 1 hour.

- On Average, during the taking of a Standardized Test, each Student will spend approximately 1 minute answering each question

- Students / Grade Level are fairly evenly distributed, meaning there are roughly 4,000+ students / grade level

- Standardized Tests are taken on different days according to Grade Level

- All students across the state take a given Standardized Test for their grade level on the same day and at roughly the same time.

## Derived Technical Requirements

The following technical requirements have been derived according to the given business requirements and assumptions made.

- The system will have significant peaks on days where standardized tests are being given, with much lower interaction and usage on days where standardized tests are not being given
- The System must be able to handle up to ~4k requests per minute during peak days in which standardized tests are being given
  - 40k students / 10 grade levels / 1 answer per minute = ~4k req/min

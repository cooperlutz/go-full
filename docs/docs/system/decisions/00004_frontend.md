# 00004: Frontend User Interface (UI)

## Status

Accepted

## Context

In order for the project to be considered a "full stack", a user interface is a necessity

## Decision

The Frontend UI will be developed as a Single Page Application in Vue and compiled to a SPA, which is ultimately embedded in Go. Further, we will aim to minimize the amount of javascript/typescript code utilized in the development of the frontend.

This decision was made based on the following considerations:

- Vue is lightweight and more easily translated to another framework if needed.
- Vue has a smaller learning curve compared to React
- Vue's Templates are more HTML centric
- Vue can utilize React components, if desired.
- Vue can utilize basic HTML/CSS/JS if desired.
- Utilizing openapi-generator we can define an openapi specification and generate the a typescript client that the frontend can consume and expose within [Vue composables](https://vuejs.org/guide/reusability/composables).

## Alternatives Considered

- HTML Templates w/ HTMX: Go's html/template package could be utilized to render server side HTML templates. However, unless we were to utilize and implement [HTMX](https://htmx.org/) a fair amount of JavaScript would be required to achieve the desired interactivity. Further, there are far fewer points of reference and examples for this approach.
- Vue w/ Nuxt SSR: this solution would require additional java/typescript code which we are trying to minimize.
- React SPA: React can get bloated and requires a significant amount of java/typescript, which we are trying to minimize.

## Consequences

- additional development effort will be required in order to handle integrations between the frontend and backend
- A programmatic client will need to be developed to support communication between the frontend and backend

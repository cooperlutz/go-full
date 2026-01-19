# 00006: Design System Basis

## Status

Accepted

## Context

For the purposes of the frontend user interface, a component & styling is required.

## Decision

TailwindCSS along with DaisyUi were chosen to serve as the basis for the design system. TailwindCSS has a strong ecosystem and many readily available reference points to support frontend development.

DaisyUi has gained a strong following and is incredibly semantic in nature. DaisyUi is also highly portable if the frontend were to be ported to React / HTML / svelte / other. DaisyUi also provides the ability to extend, override, or ignore syling of various components or themes as needed.

## Alternatives Considered

- Shadcn: while incredibly popular, if not currently the most popular, shadcn requires a fair amount of java/type scripting and is heavily react centric.
- Vuetify: is more complex and vue centric, making portability more difficult.
- PrimeVue: is more complex and vue centric, making portability more difficult.
- Material: Material Design has a strong following, but its implementation can be verbose and may not fit well with the desired aesthetic.
- Bootstrap: while widely used and easy to implement, Bootstrap's default styles may not provide the level of customization needed for the project.

## Consequences

- Developers will need to familiarize themselves with TailwindCSS and DaisyUi conventions.
- The design system will be more easily portable to other frontend frameworks if needed in the future.

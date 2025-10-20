# Decisions

We document decisions as [Architecture Decision Records](https://adr.github.io/), leveraging the format as defined and popularized by [Michael Nygard](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions). However, we are interpreting and adapting the concept of ADRs beyond strictly **architecture**, and making documented decisions inclusive of implementations and / or other relevant decisions that are worthy of documenting. ADRs provide a simple and concise format and its generally good practice to document decisions and the decision making process in just about any context.

We have also adapted the base ADR template to better suit our needs. See below for the adapted template we use.

## ADR Template

```markdown
# [00000]: Title of the Decision

## Status

[e.g., Proposed, Accepted, Rejected, Deprecated, Superseded by ADR-XXX]

## Context

Describe the issue or problem that prompted this architectural decision. What are the driving forces, requirements, or constraints that led to the need for a decision?

## Decision

Clearly state the chosen architectural decision. Explain what was decided and why this specific option was chosen.

### Implementation

If applicable, outline relevant implementation decisions, steps, or considerations that arise from this architectural decision, along with supporting considerations to support the decision.

## Alternatives Considered

List and briefly describe the alternative options that were explored but ultimately not selected. Provide a brief rationale for why each alternative was rejected.

## Consequences

Detail the implications of this decision. What are the positive and negative impacts? What becomes easier or more difficult as a result of this change? Consider impacts on development, operations, maintenance, performance, security, and other relevant aspects.
```

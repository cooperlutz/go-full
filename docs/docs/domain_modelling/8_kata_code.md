# Making the Grade: Code

Code the domain model.

## Examination

![Examination Aggregate Design Canvas](./_img/kata_aggregate_design_canvas/examination_subdomain.drawio.svg)

### 1. Name

Give your aggregate a good name. In some domains it makes sense to include as part of the name the length of a cycle, or some other indication of the life span of the aggregate.

### 2. Description

Summarise the main responsibilities and purpose of the aggregate. It’s a good idea to include the reasons why such boundaries were chosen and tradeoffs that were made compared to other designs.

### 3. State Transitions

Usually the aggregate goes through explicit state transitions, that impact the way it can be interacted with. Too many transitions might indicate that process boundaries weren't modelled properly and can be split. Very naive / simple transitions might indicate that the aggregate is anaemic and that logic was pushed out to services. In this section of the canvas list the possible states or draw a small transition diagram.

### 4. Enforced Invariants & 5. Corrective Policies

One of the main jobs of the aggregate is to enforce business invariants. These invariants protect business logic. Listing the main ones in this section will make sure that you agree on the responsibilities that the aggregate has. A large number of enforced invariants can indicate high local complexity of the aggregate implementation.

If you decide to change the boundaries of the aggregate and relax some of the invariants (for example to reduce the chance of concurrency conflict), then some extra business logic might be required to correct some of the inconsistencies. In the context of this canvas, we call this logic corrective policies. A large number of such policies might indicate that the business logic was pushed outside of the aggregate, which can increase the complexity of implementation.

Listing on the canvas both Invariants and Corrective Policies will make design trade-offs explicit and will help you decide whether the boundaries you decided on are useful or not.

### 6. Handled Commands

### 7. Created Events

In this section you list all the commands that the aggregate is capable of handling and all events that will be created as a result. It might be a good idea to create connectors between them in order to validate that you are not missing any of the building blocks.

### 8. Throughput

The goal of this section is to estimate how likely a single aggregate instance is going to be involved in concurrency conflicts (when two or more competing callers try to make changes at the same time). For each metric estimate the average and maximum - it will help you to reason about the outliers as they often drive the boundary reevaluation.

The Command handling rate metric describes the rate at which the aggregate is processing new commands. On the other hand the Total number of clients says how many clients are likely to issue these commands.

To give you an example - if an aggregate models a basket on the website then it’s likely there will be only one client issuing commands to this basket. If we compare it to an aggregate that models a conference booking system then it’s likely we are going to have tens or hundreds of clients trying to book tickets.

### 9. Size

The last section of the canvas will help you estimate the hypothetical size of the aggregate. In this case the size itself is being measured in the number of events per aggregate instance. Depending on how you model the domain, the events might end up being fine or coarse grained, so the size of the events will also play a role.

E.g. a `OrderCreated` event could have multiple `LineItems` as part of the event, or model these line items as separate `LineItemAdded` events. In case of coarse grained events, the overall size of the aggregate will be bigger, even if the number of events is small.

The Event growth rate metric should estimate how many events are appended to a single aggregate instance. The Lifetime of an instance will tell us how long the instance is going to live and as a consequence how many events will be accumulated and fetched when we need to process a new command.

## Exam Library

![Library Aggregate Design Canvas](./_img/kata_aggregate_design_canvas/library_subdomain.drawio.svg)

## Reporting

![Reporting Aggregate Design Canvas](./_img/kata_aggregate_design_canvas/reporting_subdomain.drawio.svg)

## Grading

![Grading Aggregate Design Canvas](./_img/kata_aggregate_design_canvas/grading_subdomain.drawio.svg)

## User Management

![Users Aggregate Design Canvas](./_img/kata_aggregate_design_canvas/users_subdomain.drawio.svg)

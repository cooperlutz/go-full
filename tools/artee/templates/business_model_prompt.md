You are an expert at defining business domains and the bounded contexts associated with them for the purposes of architecting microservices for software systems based on a given business model. Your task is to review the provided business model and identify the relevant business domains and bounded contexts that can be broken out into individual microservices.

A Business Domain is a distinct sphere of knowledge, activity, or influence within an organization that represents a natural division of its operations, capabilities, or areas of expertise. It provides a logical partitioning of the enterprise based on business concepts, functions, or responsibilities rather than organizational structures or technical implementations.
Business domains typically reflect fundamental business areas such as Customer Management, Product Development, Supply Chain, Finance, Human Resources, or Regulatory Compliance. Each domain encompasses related business capabilities, processes, information entities, stakeholders, and governance mechanisms. Unlike organizational units that frequently change, domains provide stable classifications based on fundamental business concerns.
For technology leaders, business domain modeling delivers significant strategic value by creating logical boundaries for system design that transcend organizational restructuring; enabling consistent data governance through domain-based ownership; facilitating microservice architecture through domain-driven design principles; providing context for capability-based planning and investment; and establishing natural boundaries for solution architecture and integration patterns.

A Subdomain represents a specific area of knowledge and functionality within the larger system. It has the following characteristics:
clear business goal: breaks down the legacy domain model that is too large to work with, into logical parts that have a clear business objective,
clear area of expertise: can indicate a specific team of domain experts / organizational department that owns the subdomain,
strategic importance: contributes to the business bringing different strategic value.

- Core subdomains provide a competitive advantage. Shipping and drone management form core subdomains for Fabrikam because they define the business. These subdomains require detailed modeling and substantial team investment.
- Supporting subdomains keep the business operational but don't differentiate it from competitors. Invoicing falls into this category. It requires custom development but doesn't serve as the competitive advantage source.
- Generic subdomains represent problems that the industry already solved. User accounts and call center form generic subdomains that Fabrikam can address by using existing prebuilt or standard solutions rather than custom-built systems.

In the context of microservices, bounded contexts are particularly relevant because they help ensure that each microservice is focused on a specific business capability or domain and that the models used in each microservice are tailored to that specific domain. When designing microservices, defining the primary domains and subdomains, boundaries of business rules, processes, and interactions around them are known as bounded context.

An aggregate is a cluster of domain objects that can be treated as a single unit. An example may be an order and its line-items, these will be separate objects, but it's useful to treat the order (together with its line items) as a single aggregate.
An aggregate will have one of its component objects be the aggregate root. Any references from outside the aggregate should only go to the aggregate root. The root can thus ensure the integrity of the aggregate as a whole.

Business model:
{{.story}}

Step 1: Review the business model and Define the relevant business Domains and bounded contexts which will be broken out into individual microservices. Each Domain should represent a specific area of the business logic and should be designed to be cohesive and loosely coupled with other Domains.

The output must be provided in yaml and conform to the following example structure:

```yaml
# values must be in snake_case!
domain_model:
- name: domain_name
  subdomains:
  - name: subdomain_name
	strategic_importance: core | supporting | generic


modules:
- name: work
  domain: project_management
  description: Module for managing work items
  bounded_context: The core domain of the system, responsible for managing work items and their associated tasks. This module encapsulates the business logic and rules related to work items, including their creation, assignment, and completion. It serves as the central point of interaction for all operations related to work items within the system.
  responsibility: This module is responsible for managing work items, including their creation, assignment, and completion. It defines the core business logic and rules associated with work items, and serves as the central point of interaction for all operations related to work items within the system.
  defaultQueries: true # if true, a default set of queries (get by id, list all) will be generated for each aggregate
  aggregates:
    - name: work_item
      description: Represents a work item in the system
      fields:
        - name: work_item_id
          type: string
          optional: false
        - name: name
          type: string
          optional: false
        - name: description
          type: string
          optional: true
		- name: story_points
		  type: int32
		  optional: true
		- name: tasks
		  optional: true
    - name: task
      description: Represents a task associated with a work item
      fields:
        - name: work_item_id
          type: string
          optional: false
        - name: name
          type: string
          optional: false
        - name: description
          type: string
          optional: true
  commands:
    - name: assign_to_user
      description: Assign a work item to a user
      events_emitted:
        - work_availeble_for_assignment
      params:
        - name: user_id
          type: string
          optional: false
        - name: work_item_id
          type: string
          optional: false
    - name: complete_work
      description: Complete a work item
      events_emitted:
        - work_item_completed
      params:
        - name: user_id
          type: string
          optional: false
        - name: work_item_id
          type: string
          optional: false
  events:
    - name: work_item_created
      description: Event triggered when a work item is created
      kind: emitted
      fields:
        - name: work_item_id
          type: string
          optional: false
        - name: name
          type: string
          optional: false
        - name: description
          type: string
          optional: true
    - name: work_availeble_for_assignment
      description: Event triggered when a work item is available for assignment
      kind: consumed
      fields:
        - name: work_item_id
          type: string
          optional: false
        - name: name
          type: string
          optional: false
        - name: description
          type: string
          optional: true
  queries:
    - name: find_work_assigned_to_user
      description: Query to find work items assigned to a specific user
      params:
        - name: user_id
          type: string
          optional: false
```
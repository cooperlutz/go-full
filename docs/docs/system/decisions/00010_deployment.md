# 00010: Deployment

## Status

Accepted

## Context

The system should be deployable and continuously deployed to a live environment to enable iteration and support fast flow of value to end users.

## Decision

The decision was made to utilize Azure Public Cloud services, specifically Azure Container Apps, for deployment of the system. This decision aligns with the organization's existing cloud strategy and leverages Azure's well-architected framework to ensure reliability, security, operational excellence, performance efficiency, and cost optimization.

## Implementation

- The system will be containerized using Docker, allowing for consistent deployment across environments.
- Azure Container Apps will be used to host the containerized application, providing a serverless environment that automatically scales based on demand.
- Infrastructure as Code (IaC) practices will be employed, specifically, Azure Bicep, to define and manage the deployment infrastructure. Azure Bicep eliminates the need to manage and maintain IaC state files, as seen with Terraform. Further, it is a Go-centric language, making it easier for Go developers to learn and utilize.

## Alternatives Considered

- On-Premises Deployment: While this option provides full control over the infrastructure, it introduces significant overhead in terms of maintenance, scalability, and disaster recovery.
- Other Cloud Providers (AWS, GCP): Although these providers offer similar services, the primary maintainer has more experience with Azure.
- Kubernetes Cluster: While this option provides flexibility and control, it introduces additional complexity in terms of management and maintenance compared to Azure Container Apps.
- Multi-Cloud Deployment: While this approach can enhance redundancy and availability, it introduces significant complexity in terms of management, cost, and potential latency issues.
- Fly.io or Heroku: These platforms offer simplicity and ease of use, but do not provide a representative environment for enterprise-grade applications and may limit future scalability and flexibility.

## Consequences

- Leveraging Azure's managed services reduces the operational burden on the development team, allowing them to focus on delivering value rather than managing infrastructure.
- If the organization decides to change cloud providers in the future, there may be migration challenges due to reliance on Azure-specific services. A portion of these challenges should be mitigated by utilizing infrastructure as code (IaC) practices as well as containerization of the application.

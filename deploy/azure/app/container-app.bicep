param name string
param location string = resourceGroup().location
param tags object = {}

param containerAppsEnvironmentName string
param containerName string = 'main'

@minValue(1)
@description('Minimum number of replicas to run')
param containerMinReplicas int = 1

@minValue(1)
@description('Maximum number of replicas to run')
param containerMaxReplicas int = 10

param secrets array = []
param env array = []
param external bool = true
param imageName string
param targetPort int = 80

@description('User assigned identity name')
param identityName string

@description('Enabled Ingress for container app')
param ingressEnabled bool = true

@description('CPU cores allocated to a single container instance, e.g. 0.5')
param containerCpuCoreCount string = '1.0'

@description('Memory allocated to a single container instance, e.g. 1Gi')
param containerMemory string = '1.0Gi'

param allowedInboundIP string

param registryCreds object

/*
  RESOURCES
*/
resource userIdentity 'Microsoft.ManagedIdentity/userAssignedIdentities@2023-01-31' existing = {
  name: identityName
}

// https://learn.microsoft.com/en-us/azure/templates/microsoft.app/containerapps?pivots=deployment-language-bicep
resource app 'Microsoft.App/containerApps@2025-02-02-preview' = {
  name: name
  location: location
  tags: tags
  dependsOn: []
  identity: {
    type: 'UserAssigned'
    userAssignedIdentities: { '${userIdentity.id}': {} }
  }
  properties: {
    managedEnvironmentId: containerAppsEnvironment.id
    configuration: {
      activeRevisionsMode: 'single'
      ingress: ingressEnabled
        ? {
            external: external
            targetPort: targetPort
            transport: 'auto'
            ipSecurityRestrictions: [
              {
                action: 'Allow'
                description: 'Allow access from specific IP'
                ipAddressRange: allowedInboundIP
                name: 'AllowSpecificIP'
              }
            ]
          }
        : null
      dapr: { enabled: false }
      secrets: secrets
      registries: [
        registryCreds
      ]
    }
    template: {
      containers: [
        {
          image: imageName
          name: containerName
          env: env
          resources: {
            cpu: json(containerCpuCoreCount)
            memory: containerMemory
          }
        }
      ]
      scale: {
        minReplicas: containerMinReplicas
        maxReplicas: containerMaxReplicas
      }
    }
  }
}

resource containerAppsEnvironment 'Microsoft.App/managedEnvironments@2022-03-01' existing = {
  name: containerAppsEnvironmentName
}

/*
  OUTPUTS
*/
output defaultDomain string = containerAppsEnvironment.properties.defaultDomain
output imageName string = imageName
output name string = app.name
output uri string = 'https://${app.properties.configuration.ingress.fqdn}'

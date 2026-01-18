param name string
param location string = resourceGroup().location
param tags object = {}
param containerAppsEnvironmentName string
param containerName string
@minValue(1)
param containerMinReplicas int = 1
@minValue(1)
param containerMaxReplicas int = 10
param secrets array = []
param env array = []
param external bool = true
param imageName string
param targetPort int = 80
param identityName string
param ingressEnabled bool = true
param containerCpuCoreCount string = '1.0'
param containerMemory string = '1.0Gi'
param registryCreds object
param allowedInboundIP string

/* RESOURCES */

@description('User Assigned Managed Identity resource')
resource userIdentity 'Microsoft.ManagedIdentity/userAssignedIdentities@2023-01-31' existing = {
  name: identityName
}

@description('Container App resource')
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

@description('Container Apps Environment resource')
resource containerAppsEnvironment 'Microsoft.App/managedEnvironments@2022-03-01' existing = {
  name: containerAppsEnvironmentName
}

/* OUTPUTS */
output defaultDomain string = containerAppsEnvironment.properties.defaultDomain
output imageName string = imageName
output name string = app.name
output uri string = 'https://${app.properties.configuration.ingress.fqdn}'

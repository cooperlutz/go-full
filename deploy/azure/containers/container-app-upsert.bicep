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
param targetPort int = 8080
param identityName string
param ingressEnabled bool = true
param containerCpuCoreCount string = '0.5'
param containerMemory string = '1.0Gi'
param allowedInboundIP string
param registryCreds object
param imageName string

/* MODULES */
module app 'container-app.bicep' = {
  name: '${deployment().name}-update'
  params: {
    name: name
    location: location
    tags: tags
    identityName: identityName
    ingressEnabled: ingressEnabled
    containerName: containerName
    containerAppsEnvironmentName: containerAppsEnvironmentName
    registryCreds: registryCreds
    containerCpuCoreCount: containerCpuCoreCount
    containerMemory: containerMemory
    containerMinReplicas: containerMinReplicas
    containerMaxReplicas: containerMaxReplicas
    secrets: secrets
    external: external
    env: env
    imageName: imageName
    targetPort: targetPort
    allowedInboundIP: allowedInboundIP
  }
}

/* OUTPUTS */
output defaultDomain string = app.outputs.defaultDomain
output imageName string = app.outputs.imageName
output name string = app.outputs.name
output uri string = app.outputs.uri

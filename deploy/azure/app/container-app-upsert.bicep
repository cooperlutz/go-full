param name string
param location string = resourceGroup().location
param tags object = {}

@description('Name of the Container App Environment')
param containerAppsEnvironmentName string

@description('Name of the Container App')
param containerName string = 'main'

@description('Minimum number of replicas to run')
@minValue(1)
param containerMinReplicas int = 1

@description('Maximum number of replicas to run')
@minValue(1)
param containerMaxReplicas int = 10

param secrets array = []
param env array = []
param external bool = true
param targetPort int = 8080

@description('User assigned identity name')
param identityName string

@description('Enabled Ingress for container app')
param ingressEnabled bool = true

@description('CPU cores allocated to a single container instance, e.g. 0.5')
param containerCpuCoreCount string = '0.5'

@description('Memory allocated to a single container instance, e.g. 1Gi')
param containerMemory string = '1.0Gi'

param allowedInboundIP string

param registryCreds object

param imageName string

/*
  MODULES
*/
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

/*
  OUTPUTS
*/
output defaultDomain string = app.outputs.defaultDomain
output imageName string = app.outputs.imageName
output name string = app.outputs.name
output uri string = app.outputs.uri

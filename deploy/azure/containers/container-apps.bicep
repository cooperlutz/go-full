param name string
param location string = resourceGroup().location
param tags object = {}
param containerAppsEnvironmentName string
param logAnalyticsWorkspaceName string

/* MODULES */
module containerAppsEnvironment 'container-apps-environment.bicep' = {
  name: '${name}-container-apps-environment'
  params: {
    name: containerAppsEnvironmentName
    location: location
    tags: tags
    logAnalyticsWorkspaceName: logAnalyticsWorkspaceName
  }
}

/* OUTPUTS */
output defaultDomain string = containerAppsEnvironment.outputs.defaultDomain
output environmentName string = containerAppsEnvironment.outputs.name
output environmentIpAddress string = containerAppsEnvironment.outputs.ipAddress

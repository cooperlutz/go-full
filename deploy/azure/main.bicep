targetScope = 'subscription'

@description('Name which is used to generate a short unique hash for each resource')
param name string

@description('Primary location for all resources')
param location string

@secure()
@description('Password for the PostgreSQL database admin user')
param dbPassword string

@secure()
@description('Personal Access Token for the container registry')
param containerRegPat string

@description('Container registry username')
param containerRegUserName string

@description('IP Address or CIDR range allowed to access the application')
param allowedInboundIP string 

@description('Name of the overall application')
param applicationName string

@description('Admin username for the PostgreSQL database')
param databaseAdminUser string

@description('Tag of the container image to deploy as defined in the container registry')
param imageTag string

@description('Type of database to use. Supported values: postgres')
param databaseType string

@description('Name of the primary database to create')
param databaseName string

@description('SSL mode for database connections. Supported values: disable, require, verify-ca, verify-full')
param databaseSSLMode string

@description('HTTP port for the application')
param httpPort string

/*
  VARIABLES
*/
var resourceToken = toLower(uniqueString(subscription().id, name, location))

var prefix = '${applicationName}-${name}-${resourceToken}'
var resourceGroupName = '${prefix}-rg'
var tags = { 'azd-env-name': name }
// Observability
var appInsightsName = replace('${take(prefix,19)}-ai', '--', '-')
var logAnalyticsName = replace('${take(prefix,19)}-la', '--', '-')
// Container Apps
var containerAppsEnvironmentName = '${prefix}-containerapps-env'
var containerRegSecretName = 'ghcr-pat'
var containerRegServer = 'ghcr.io'
var containerImageName = '${containerRegServer}/cooperlutz/go-full:${imageTag}'
// var containerImageName = imageTag
var ghRegistry = {
  server: containerRegServer
  username: containerRegUserName
  passwordSecretRef: containerRegSecretName
}
var containerAppName = replace('${take('${applicationName}',19)}-ca', '--', '-')
var serviceIdentityName = '${prefix}-identity'
// Database
var postgresServerName = replace('${take(prefix,18)}-psql', '--', '-')

var databases = [
  databaseName
]
var postgresVersion = '16'
var dbSizeInGB = 32
var dbServerSku = {
  name: 'Standard_B1ms'
  tier: 'Burstable'
}

/*
  RESOURCES
*/
resource rg 'Microsoft.Resources/resourceGroups@2022-09-01' = {
  name: resourceGroupName
  location: location
  tags: tags
}

/*
  MODULES
*/
module postgresServer 'database/postgresql.bicep' = {
  name: 'postgresql'
  scope: rg
  params: {
    name: postgresServerName
    location: location
    tags: tags
    sku: dbServerSku
    storage: {
      storageSizeGB: dbSizeInGB
    }
    version: postgresVersion
    administratorLogin: databaseAdminUser
    administratorLoginPassword: dbPassword
    allowAzureIPsFirewall: true
    allowAllIPsFirewall: true
    databaseNames: databases
  }
}

module logAnalytics 'observability/log-analytics.bicep' = {
  name: 'logAnalytics'
  scope: rg
  params: {
    name: logAnalyticsName
    location: location
    tags: tags
  }
}

module applicationInsights 'observability/application-insights.bicep' = {
  name: 'applicationInsights'
  scope: rg
  params: {
    name: appInsightsName
    workspaceResourceId: logAnalytics.outputs.id
    location: location
    tags: tags
  }
}

// Container apps host (including container registry)
module containerApps 'app/container-apps.bicep' = {
  name: 'container-apps'
  scope: rg
  params: {
    name: 'app'
    location: location
    tags: tags
    containerAppsEnvironmentName: containerAppsEnvironmentName
    logAnalyticsWorkspaceName: logAnalytics.outputs.name
  }
}

// Web frontend
module web 'app.bicep' = {
  name: 'web'
  scope: rg
  params: {
    name: containerAppName
    location: location
    tags: tags
    identityName: serviceIdentityName
    registryCreds: ghRegistry
    containerImageName: containerImageName
    containerAppsEnvironmentName: containerApps.outputs.environmentName
    postgresDomainName: postgresServer.outputs.POSTGRES_DOMAIN_NAME
    postgresUser: databaseAdminUser
    postgresPassword: dbPassword
    postgresDatabaseName: databaseName
    containerRegPat: containerRegPat
    containerRegSecretName: containerRegSecretName
    allowedInboundIP: allowedInboundIP
    targetPort: httpPort
    dbSSLMode: databaseSSLMode
    dbType: databaseType
  }
}
/*
  OUTPUTS
*/
output postgresServerName string = postgresServer.outputs.POSTGRES_SERVER_NAME
output postgresDomainName string = postgresServer.outputs.POSTGRES_DOMAIN_NAME
output postgresAdminUsername string = postgresServer.outputs.POSTGRES_ADMIN_USERNAME
output AZURE_RESOURCE_GROUP string = rg.name
output POSTGRES_SERVER_NAME string = postgresServer.outputs.POSTGRES_SERVER_NAME
output POSTGRES_SERVER_USERNAME string = postgresServer.outputs.POSTGRES_ADMIN_USERNAME

param name string
param location string = resourceGroup().location
param tags object = {}
param identityName string
param containerAppsEnvironmentName string
param serviceName string = 'app'
param postgresDomainName string
param postgresDatabaseName string
param postgresUser string
param allowedInboundIP string
param registryCreds object
param containerRegSecretName string
param containerImageName string
param targetPort string
param dbSSLMode string
param dbType string
@secure()
param postgresPassword string
@secure()
param containerRegPat string

/*
  RESOURCES
*/
resource webIdentity 'Microsoft.ManagedIdentity/userAssignedIdentities@2023-01-31' = {
  name: identityName
  location: location
}

/*
  MODULES
*/
module app 'app/container-app-upsert.bicep' = {
  name: '${serviceName}-container-app-module'
  params: {
    name: name
    location: location
    tags: union(tags, { 'azd-service-name': serviceName })
    identityName: webIdentity.name
    registryCreds: registryCreds
    imageName: containerImageName
    containerAppsEnvironmentName: containerAppsEnvironmentName
    allowedInboundIP: allowedInboundIP
    env: [
      {
        name: 'DB_HOST'
        value: postgresDomainName
      }
      {
        name: 'DB_DBNAME'
        value: postgresDatabaseName
      }
      {
        name: 'DB_USER'
        value: postgresUser
      }
      {
        name: 'DB_PORT'
        value: 5432
      }
      {
        name: 'OBSERVE_TRACE_ENDPOINT'
        value: 'localhost:4317'
      }
      {
        name: 'DB_TYPE'
        value: dbType
      }
      {
        name: 'DB_SSLMODE'
        value: dbSSLMode
      }
      {
        name: 'DB_PASSWORD'
        secretRef: 'dbpass'
      }
      {
        name: 'HTTP_PORT'
        value: targetPort
      }
      {
        name: 'APP_NAME'
        value: name
      }
      {
        name: 'APP_VERSION'
        value: '1.0.0'
      }
      {
        name: 'CONTAINER_REGISTRY_PAT'
        secretRef: containerRegSecretName
      }
    ]
    targetPort: 8080
    secrets: [
      {
        name: 'dbpass'
        value: postgresPassword
      }
      {
        name: containerRegSecretName
        value: containerRegPat
      }
    ]
  }
}

/*
  OUTPUTS
*/
output SERVICE_WEB_IDENTITY_PRINCIPAL_ID string = webIdentity.properties.principalId
output SERVICE_WEB_NAME string = app.outputs.name
output SERVICE_WEB_URI string = app.outputs.uri
output SERVICE_WEB_IMAGE_NAME string = app.outputs.imageName

@description('Name of the PostgreSQL server')
param name string

@description('Location for database resource')
param location string = resourceGroup().location

@description('Tags to apply to database resources')
param tags object = {}

@description('SKU for the PostgreSQL server')
param sku object

@description('Storage configuration for the PostgreSQL server')
param storage object

@description('Administrator login for the PostgreSQL server')
param administratorLoginUserName string

@secure()
@description('Administrator login password for the PostgreSQL server')
param administratorLoginPassword string

@description('The name of the PostgreSQL database to create on the server')
param databaseName string

@description('Allow access from all Azure internal IPs')
param allowAzureIPsFirewall bool = true

@description('Allow access from specific IP addresses')
param allowedSingleIPs array

@description('PostgreSQL version for the server')
param version string

/* RESOURCES */
@description('The PostgreSQL flexible server instance')
resource postgresServer 'Microsoft.DBforPostgreSQL/flexibleServers@2025-08-01' = {
  location: location
  tags: tags
  name: name
  sku: sku
  properties: {
    version: version
    administratorLogin: administratorLoginUserName
    administratorLoginPassword: administratorLoginPassword
    storage: storage
    highAvailability: {
      mode: 'Disabled'
    }
    network: {
      publicNetworkAccess: 'Enabled'
    }
  }
}

@description('The PostgreSQL database to create on the server')
resource database 'Microsoft.DBforPostgreSQL/flexibleServers/databases@2025-08-01' = {
  parent: postgresServer
  name: databaseName
}

// // This must be done separately due to conflicts with the Entra setup
@description('Allow access from all Azure internal IPs')
resource firewall_azure 'Microsoft.DBforPostgreSQL/flexibleServers/firewallRules@2023-03-01-preview' = if (allowAzureIPsFirewall) {
  parent: postgresServer
  name: 'allow-all-azure-internal-IPs'
  properties: {
    startIpAddress: '0.0.0.0'
    endIpAddress: '0.0.0.0'
  }
}

// This must be done separately due to conflicts with the Entra setup
@batchSize(1)
@description('Allow access from specific IP addresses')
resource firewall_single 'Microsoft.DBforPostgreSQL/flexibleServers/firewallRules@2023-03-01-preview' = [
  for ip in allowedSingleIPs: {
    parent: postgresServer
    name: 'allow-single-${replace(ip, '.', '')}'
    properties: {
      startIpAddress: ip
      endIpAddress: ip
    }
  }
]

/* OUTPUTS */
output POSTGRES_DOMAIN_NAME string = postgresServer.properties.fullyQualifiedDomainName
output POSTGRES_SERVER_NAME string = postgresServer.name
output POSTGRES_ADMIN_USERNAME string = administratorLoginUserName

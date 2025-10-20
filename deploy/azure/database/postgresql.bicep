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
param administratorLogin string = ''

@secure()
@description('Administrator login password for the PostgreSQL server')
param administratorLoginPassword string = ''

param databaseNames array = []
param allowAzureIPsFirewall bool = false
param allowAllIPsFirewall bool = false
param allowedSingleIPs array = []
param version string

/*
  RESOURCES
*/
resource postgresServer 'Microsoft.DBforPostgreSQL/flexibleServers@2022-12-01' = {
  location: location
  tags: tags
  name: name
  sku: sku
  properties: {
    version: version
    administratorLogin: administratorLogin
    administratorLoginPassword: administratorLoginPassword
    storage: storage
    highAvailability: {
      mode: 'Disabled'
    }
  }

  resource database 'databases' = [
    for name in databaseNames: {
      name: name
    }
  ]
}

// This must be done separately due to conflicts with the Entra setup
resource firewall_all 'Microsoft.DBforPostgreSQL/flexibleServers/firewallRules@2023-03-01-preview' = if (allowAllIPsFirewall) {
  parent: postgresServer
  name: 'allow-all-IPs'
  properties: {
    startIpAddress: '0.0.0.0'
    endIpAddress: '255.255.255.255'
  }
}

// This must be done separately due to conflicts with the Entra setup
resource firewall_azure 'Microsoft.DBforPostgreSQL/flexibleServers/firewallRules@2023-03-01-preview' = if (allowAzureIPsFirewall) {
  parent: postgresServer
  name: 'allow-all-azure-internal-IPs'
  properties: {
    startIpAddress: '0.0.0.0'
    endIpAddress: '0.0.0.0'
  }
}

@batchSize(1)
// This must be done separately due to conflicts with the Entra setup
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

/*
  OUTPUTS
*/
output POSTGRES_DOMAIN_NAME string = postgresServer.properties.fullyQualifiedDomainName
output POSTGRES_SERVER_NAME string = postgresServer.name
output POSTGRES_ADMIN_USERNAME string = administratorLogin

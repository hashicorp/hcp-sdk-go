# HashiCorp Cloud Platform Go SDK

This SDK provides versioned Go packages for using HashiCorp Cloud Platform services. This is a *preview* version. It is not intended for use outside of local testing environments.

## Installation

Fetch and install the package:

```bash
go get github.com/hashicorp/hcp-sdk-go
```

## Authentication

The `client_id` and `client_secret` must come from a service principal key. *Note:* The `client_secret` can only be obtained on creation of the service principal key; it is not stored anywhere after that.

The service principal must be authorized to access the API. Initially, it has no permissions, so the IAM policy must be updated to grant it permissions.

Follow these steps to create service principal with the `contributor` role and a service principal key.

### 1. Create a service principal

Once you have registered and logged into the HCP portal, navigate to the Access Control (IAM) page. Select the Service Principals tab and create a new service principal. Give it the role Contributor, since it will be writing resources.

### 2. Create a service principal key

Once the service principal is created, navigate to its detail page by selecting its name in the list. Create a new service principal key.

**Note:** Save the client ID and secret returned on successful key creation. The client secret will not be available after creation.

### 3. Configure the SDK with client credentials

Set the client ID and secret as the environment variables HCP_CLIENT_ID and HCP_CLIENT_SECRET.

```bash
export HCP_CLIENT_ID="service-principal-key-client-id"
export HCP_CLIENT_SECRET="service-principal-key-client-secret"
```

## Usage

Import the desired version of each service SDK.

```bash
import (
    network "github.com/hashicorp/hcp-sdk-go/clients/cloud-network/preview/2020-09-07/client/network_service"
)
```

See `cmd/hcp-sdk-go-client` for a complete example.

## Generating a new service SDK

Requirements:

- must be an internal Hashicorp employee
- `hcloud` ([install instructions](https://github.com/hashicorp/hcloud#installation))
- `gh` (`brew install gh`)
- GITHUB_TOKEN with `repo` permission ([how to create a personal access token](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token))

The `make sdk/update` command will:

1. Use the tool `hcloud` to pull temporary copies of the latest public specs from the designated HCP service
2. Run transformations on the temporary spec copiess
3. Generate the SDK from the transformed specs.
4. (optional) Post a PR to this repo if `commit=true` is passed

To generate the SDK for cloud-foo-service ***without** committing the files in a PR:

```bash
make sdk/update service=cloud-foo-service commit=false
```

To generate the SDK and commit the files in a PR:

```bash
make sdk/update service=cloud-foo-service commit=true
```

To generate shared HCP models and commit the files in a PR:

```bash
make sdk/update service=cloud-shared commit=true
```

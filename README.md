# HashiCorp Cloud Platform Go SDK

This SDK provides versioned Go packages for using HashiCorp Cloud Platform services. This is a *preview* version and may undergo breaking changes.

## SDK Release Cycle

HCP aims to strike a balance between the public SDK consumer's need for a stable SDK interface, and the developer's need to iterate on new features. To this end, we use SDK versions composed of a dated version (`2021-02-04`) and stage (`preview` or `stable`) to allow consumers the flexibility to pin to specific versions or try out new new features.

Preview SDK versions allow consumers to try out features in beta, but may undergo breaking changes. What constitutes a breaking change is described in the [Breaking Changes doc](/docs/breaking-changes.md).

Stable SDK versions present a guaranteed contract to downstream consumers and **will not undergo breaking changes.** Major new features are added to the latest stable version after vetting with a preview version. Any breaking change to a stable SDK version will require a new preview version, which will ultimately be promoted to the next stable version.

### Steps

![SDK Release Cycle Diagram](/images/sdk-release-cycle-diagram.png)

1. `stable/2021-02-04` **Only small backwards compatible changes and bug fixes are released directly to the latest stable API version.** No breaking changes allowed.

1. `preview/2021-09-14` **A preview SDK version is released with features available in public beta.** This version is under active development and may still undergo potential breaking changes.

1. `stable/2021-10-25` **A new stable SDK version with a later date is released once feature iteration is complete.** This new stable version includes all the finalized changes of the last preview version.

1. `preview/2021-11-07` **A new preview version set to the current date may be released for experimental backwards-compatible features.** After the experimental feature is determined viable, the finalized changes are merged into the latest stable version, similar to how feature branches merge back into main.

1. **Any subsequent breaking change starts the cycle all over again.**

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

1. Add the SDK to your go.mod.

```go
require (
     github.com/hashicorp/hcp-sdk-go {latest release}
```

2. Import the desired version of each service SDK.

```bash
import (
    network "github.com/hashicorp/hcp-sdk-go/clients/cloud-network/preview/2020-09-07/client/network_service"
)
```

See `cmd/hcp-sdk-go-client` for a complete example.

## Contributing

### Changelogs

This repo requires that a chagnelog file be added in all pull requests. The name of the file must follow `[PR #].txt` and must reside in the `.changelog` directory. The contents must have the following formatting:

~~~
```release-note:TYPE
ENTRY
```
~~~

Where `TYPE` is the type of release note entry this is. This is one of either: `breaking-change`, `security`, `feature`, `improvement`, `deprecation`, `bug`.

`ENTRY` is the body of the changelog entry, and should describe the changes that were made. This is used as free-text input and will be returned to you as it is entered when generating the changelog.

Sometimes PRs have multiple changelog entries associated with them. In this case, use multiple blocks.

~~~
```release-note:deprecation
Deprecated the `foo` interface, please use the `bar` interface instead.
```

```release-note:improvement
Added the `bar` interface.
```
~~~


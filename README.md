# HashiCorp Cloud Platform Go SDK

This SDK provides versioned Go packages for using HashiCorp Cloud Platform services. This is a *preview* version and may undergo breaking changes.

## Installation

Fetch and install the package:

```bash
go get github.com/hashicorp/hcp-sdk-go
```

## Authentication

There are currently two ways the HCP Go SDK can authenticate against HCP:

- Client credentials
- User session, obtained through browser login window

### Client Credentials

Client credentials are recommended for CI and local development with the SDK or any tool consuming it.

Follow the guidance in the [HCP Terraform Provider's auth guide](https://registry.terraform.io/providers/hashicorp/hcp/latest/docs/guides/auth#service-principal-credentials) to create a service principal and a service principal key with client credentials.

When client credentials are set, they are always used by the HCP Go client, regardless of an existing user session.

The method of authentication is determined by environment variables.

```bash
HCP_CLIENT_ID="..."
HCP_CLIENT_SECRET="..."
```

### User Session

User session is ideal for getting started or one-off usage. It also works for local development, but will periodically prompt for re-authentication.

To obtain user credentials, the client credential environment variables `HCP_CLIENT_ID` and `HCP_CLIENT_SECRET` must be unset. When no client credentials are detected, the HCP Go client will prompt the user with a browser login window. Once authenticated, the user session stays refreshed without intervention until it expires after 24 hours.

If you have a use-case with the SDK to leverage the browser login as a feature but want to control if the browser is opened, or even to understand if the system already has a valid token present, you can pass in the option func of `WithoutBrowserLogin()` to your `NewHCPConfig()`. This will use either the provided `ClientID`:`ClientSecret` combo or a valid token that has been previously written to the system. If neither option exists, then `auth.ErrorNoLocalCredsFound` is returned to indicate that the user is not yet logged in.

### User Profile

An HCP Organization ID and Project ID are required to call most HCP APIs. They can be set to the environment variables `HCP_ORGANIZATION_ID` and `HCP_PROJECT_ID`, as in the example below. The HCP Go SDK will read them from the environment and save them in its state as the user's Profile. The Profile Project and Organization IDs will be applied as default values to any request missing them.

```bash
HCP_PROJECT_ID="33xyz..."
HCP_ORGANIZATION_ID="22abc..."
```

## Usage

1. Add the SDK to your go.mod.

    ```go
    require (
        github.com/hashicorp/hcp-sdk-go {latest release}
    ```

1. Import the desired version of each service SDK.

    ```bash
    import (
        network "github.com/hashicorp/hcp-sdk-go/clients/cloud-network/stable/2020-09-07/client/network_service"
    )
    ```

See `cmd/hcp-sdk-go-client` for a complete example.

### SDK Release Cycle

HCP aims to strike a balance between the public SDK consumer's need for a stable SDK interface, and the developer's need to iterate on new features. To this end, we use SDK versions composed of a dated version (`2021-02-04`) and stage (`preview` or `stable`) to allow consumers the flexibility to pin to specific versions or try out new new features.

Preview SDK versions allow consumers to try out features in beta, but may undergo breaking changes. What constitutes a breaking change is described in the [Breaking Changes doc](/docs/breaking-changes.md).

Stable SDK versions present a guaranteed contract to downstream consumers and **will not undergo breaking changes.** Major new features are added to the latest stable version after vetting with a preview version. Any breaking change to a stable SDK version will require a new preview version, which will ultimately be promoted to the next stable version.

#### Steps

![SDK Release Cycle Diagram](/images/sdk-release-cycle-diagram.png)

1. `stable/2021-02-04` **Only small backwards compatible changes and bug fixes are released directly to the latest stable API version.** No breaking changes allowed.

1. `preview/2021-09-14` **A preview SDK version is released with features available in public beta.** This version is under active development and may still undergo potential breaking changes.

1. `stable/2021-10-25` **A new stable SDK version with a later date is released once feature iteration is complete.** This new stable version includes all the finalized changes of the last preview version.

1. `preview/2021-11-07` **A new preview version set to the current date may be released for experimental backwards-compatible features.** After the experimental feature is determined viable, the finalized changes are merged into the latest stable version, similar to how feature branches merge back into main.

1. **Any subsequent breaking change starts the cycle all over again.**

## Libraries

In addition to the generated product clients, the HCP Go SDK provides a few libraries useful for interacting with HCP.

### Cache

The Cache interface lives under the `auth` package. It handles writing the user credentials obtained during browser login to the common location `/.config/hcp/credentials.json` in the home directory. The Cache has `Read` and `Write` methods that can be used to get and set stored HCP credentials.

Generally the contents of the Cache should be Read to get the latest, unexpired credentials. Without care, overwriting user credentials may cause unexpected authentication failures.

## Contributing

### Changelogs

This repo requires that a chagnelog file be added in all pull requests. The name of the file must follow `[PR #].txt` and must reside in the `.changelog` directory. The contents must have the following formatting:

```text
    ```release-note:TYPE
    ENTRY
    ```
```

Where `TYPE` is the type of release note entry this is. This is one of either: `breaking-change`, `security`, `feature`, `improvement`, `deprecation`, `bug`.

`ENTRY` is the body of the changelog entry, and should describe the changes that were made. This is used as free-text input and will be returned to you as it is entered when generating the changelog.

Sometimes PRs have multiple changelog entries associated with them. In this case, use multiple blocks.

```text
    ```release-note:deprecation
    Deprecated the `foo` interface, please use the `bar` interface instead.
    ```

    ```release-note:improvement
    Added the `bar` interface.
    ```
```

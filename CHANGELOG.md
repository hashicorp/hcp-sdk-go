## v0.65.0 (October 04, 2023)

## v0.64.0 (September 27, 2023)

## v0.63.0 (September 20, 2023)

## v0.62.0 (September 13, 2023)

## v0.61.0 (August 28, 2023)

## v0.60.0 (August 23, 2023)

## v0.59.0 (August 09, 2023)

## v0.58.0 (August 01, 2023)

IMPROVEMENTS:

* Credential file omits not set fields when encoded as JSON. [[GH-202](https://github.com/hashicorp/hcp-sdk-go/pull/202)]
## v0.57.0 (July 31, 2023)

FEATURES:

* SDK can authenticate using Workload Identity Federation. [[GH-199](https://github.com/hashicorp/hcp-sdk-go/pull/199)]
* SDK can authenticate using a credential file. The credential file can specify
service principal credentials or workload identity provided credentials. [[GH-200](https://github.com/hashicorp/hcp-sdk-go/pull/200)]
## v0.56.0 (July 28, 2023)

## v0.55.0 (July 26, 2023)

## v0.54.0 (July 19, 2023)

## v0.53.0 (July 19, 2023)

## v0.52.0 (July 12, 2023)

## v0.51.0 (July 05, 2023)

IMPROVEMENTS:

* Use refresh tokens if the session is still valid [[GH-196](https://github.com/hashicorp/hcp-sdk-go/pull/196)]
## v0.50.0 (June 14, 2023)

## v0.49.0 (June 07, 2023)

## v0.48.0 (May 24, 2023)

IMPROVEMENTS:

* Added env var option (HCP_AUTH_TLS) to either disable or set to insecure the auth call for mocking. [[GH-190](https://github.com/hashicorp/hcp-sdk-go/pull/190)]
## v0.47.0 (May 24, 2023)

IMPROVEMENTS:

* Added option to NewHCPConfig to disable logging from the SDK [[GH-192](https://github.com/hashicorp/hcp-sdk-go/pull/192)]
## v0.46.0 (May 17, 2023)

## v0.45.0 (May 10, 2023)

## v0.44.0 (May 03, 2023)

## v0.43.0 (April 26, 2023)

## v0.42.0 (April 19, 2023)

IMPROVEMENTS:

* Added option to NewHCPConfig to fail rather than auto login with web browser [[GH-182](https://github.com/hashicorp/hcp-sdk-go/pull/182)]

BUG FIXES:

* Fixed issue of not regenerating auth token when AccessToken expires [[GH-181](https://github.com/hashicorp/hcp-sdk-go/pull/181)]
## v0.41.0 (April 05, 2023)

## v0.40.0 (March 29, 2023)

## v0.39.0 (March 22, 2023)

## v0.38.0 (March 15, 2023)

## v0.37.0 (March 14, 2023)

IMPROVEMENTS:

* [Bump golang.org/x/oauth2 from 0.0.0-20201109201403-9fd604954f58 to 0.6.0](https://github.com/hashicorp/hcp-sdk-go/commit/e2804f559f9bd140cb5fc690075dfc441aa4b33a) 
* [Bump github.com/stretchr/testify from 1.8.1 to 1.8.2](https://github.com/hashicorp/hcp-sdk-go/commit/e95aaacf3b1968029f11fe2e1759721a1fac48df)
* [Bump golang.org/x/net from 0.7.0 to 0.8.0](https://github.com/hashicorp/hcp-sdk-go/commit/a78fdfd568526d90031e132659d1b1fb2dc8dcee) 
* [change owner to ept pod](https://github.com/hashicorp/hcp-sdk-go/commit/a62bbbb15477bff967ed3c627689f2d208dcd681)
* [Update cloud-shared SDK](https://github.com/hashicorp/hcp-sdk-go/commit/c920f189788e4880873dab25789a10579346822b)
* [HCE-823: consolidate API pipeline actions/scripts](https://github.com/hashicorp/hcp-sdk-go/commit/8144158603dcc22241b2a3fdb916d91cc20ffe78) 

BREAKING CHANGES:

* [Update cloud-resource-manager SDK: bump to stable](https://github.com/hashicorp/hcp-sdk-go/commit/0d9e3b043319dff265d09e4235a3f2988b3df277)
* [Update cloud-network SDK: bump to stable](https://github.com/hashicorp/hcp-sdk-go/commit/27684221c6d89b2af255c87b1bd58dd5091ca573)
* [Update cloud-boundary-service SDK: bump to stable](https://github.com/hashicorp/hcp-sdk-go/commit/c43b104c25a48c805d0a6f57b502693bb0880c2f)

## v0.36.0 (March 01, 2023)

IMPROVEMENTS:

* Bump golang.org/x/net from 0.5.0 to 0.7.0 [[GH-170](https://github.com/hashicorp/hcp-sdk-go/pull/170)]
## v0.35.0 (February 15, 2023)

## v0.34.0 (February 10, 2023)

IMPROVEMENTS:

- [Update cloud-consul-service SDK](https://github.com/hashicorp/hcp-sdk-go/commit/25b608253197bd5642dbd62af7f8062d1e9e86c3)
 - [fix linting errors](https://github.com/hashicorp/hcp-sdk-go/commit/c9c7d254d30e8ed9a39a99a438f20b77c79e1d8d)
- [Update cloud-vault-link-service SDK](https://github.com/hashicorp/hcp-sdk-go/commit/587b65ed57fe8542a387b03056e570c35dc25133)

## v0.33.0 (February 08, 2023)

IMPROVEMENTS:

- [Update cloud-operation SDK](https://github.com/hashicorp/hcp-sdk-go/commit/8ece4c5ddd56f73bc731e3d5b2591f377d37457d)

## v0.32.0 (February 01, 2023)

IMPROVEMENTS:

* Bump github.com/go-openapi/spec from 0.20.7 to 0.20.8 [[GH-160](https://github.com/hashicorp/hcp-sdk-go/pull/160)]
* Bump github.com/go-openapi/validate from 0.22.0 to 0.22.1 [[GH-159](https://github.com/hashicorp/hcp-sdk-go/pull/159)]

BUG FIXES:

* Update action to read go version from repo go.mod [[GH-161](https://github.com/hashicorp/hcp-sdk-go/pull/161)]
## v0.31.0 (January 11, 2023)

FEATURES:

* Automatically sync the public and internal repos. [[GH-157](https://github.com/hashicorp/hcp-sdk-go/pull/157)]
## v0.30.0 (December 21, 2022)

IMPROVEMENTS:

* Add middleware support to httpclient package [[GH-142](https://github.com/hashicorp/hcp-sdk-go/pull/142)]
* Add middleware that gets org ID and project ID from user profile and sets on request [[GH-142](https://github.com/hashicorp/hcp-sdk-go/pull/142)]
* Add new Libraries section to README. [[GH-155](https://github.com/hashicorp/hcp-sdk-go/pull/155)]
* Bump github.com/go-openapi/loads from 0.20.2 to 0.21.2 [[GH-149](https://github.com/hashicorp/hcp-sdk-go/pull/149)]
* Bump github.com/go-openapi/runtime from 0.19.24 to 0.25.0 [[GH-148](https://github.com/hashicorp/hcp-sdk-go/pull/148)]
* Bump github.com/go-openapi/spec from 0.20.3 to 0.20.7 [[GH-147](https://github.com/hashicorp/hcp-sdk-go/pull/147)]
* Bump github.com/go-openapi/swag from 0.19.14 to 0.22.3 [[GH-146](https://github.com/hashicorp/hcp-sdk-go/pull/146)]
* Bump github.com/go-openapi/validate from 0.21.0 to 0.22.0 [[GH-153](https://github.com/hashicorp/hcp-sdk-go/pull/153)]
* Bump github.com/hashicorp/go-cleanhttp from 0.5.1 to 0.5.2 [[GH-144](https://github.com/hashicorp/hcp-sdk-go/pull/144)]
* Bump github.com/iancoleman/strcase from 0.1.3 to 0.2.0 [[GH-145](https://github.com/hashicorp/hcp-sdk-go/pull/145)]
* Remove `create-pull-request` third party action in favour of plain git commands. [[GH-152](https://github.com/hashicorp/hcp-sdk-go/pull/152)]
## v0.29.0 (December 14, 2022)

IMPROVEMENTS:

* Enable automatic changelog creation for dependabot PRs. [[GH-150](https://github.com/hashicorp/hcp-sdk-go/pull/150)]
## v0.28.0 (December 07, 2022)
BREAKING CHANGES:

* The cloud-packer service was updated to grpc_gateway v2, this changes the payloads for the requests in some routes, as they are now embedded within the request rather than referenced from it.
For example, the body for UpdateBucket was previously a *models.HashicorpCloudPackerUpdateBucketRequest, now it is a packer_service.PackerServiceUpdateBucketBody.
Some attributes which were duplicated in the body now are only in the parent structure, for example &models.HashicorpCloudPackerUpdateBucketRequest used to contain BucketSlug, already present in the parent structure, this is not the case anymore. [[GH-140](https://github.com/hashicorp/hcp-sdk-go/pull/140)]
## v0.27.0 (December 01, 2022)

FEATURES:

* Allow users to save profile information via environment variables [[GH-137](https://github.com/hashicorp/hcp-sdk-go/pull/137)]

BUG FIXES:

* Remove v0.26.0 from pkg.go.dev [[GH-139](https://github.com/hashicorp/hcp-sdk-go/pull/139)]
## v0.26.0 (November 30, 2022)

BUG FIXES:

* Added exception for google.rpc.status when generating code since it was mistakenly rendered as `GoogleRpcStatus` instead of `GoogleRPCStatus` [[GH-138](https://github.com/hashicorp/hcp-sdk-go/pull/138)]
## v0.25.0 (November 23, 2022)

FEATURES:

* Enable user session refresh on browser login [[GH-129](https://github.com/hashicorp/hcp-sdk-go/pull/129)]

IMPROVEMENTS:

* Bump github.com/go-openapi/strfmt from 0.20.0 to 0.21.3 (https://github.com/hashicorp/hcp-sdk-go/pull/132) [[GH-132](https://github.com/hashicorp/hcp-sdk-go/pull/132)]
* Expand test coverage for browser login [[GH-129](https://github.com/hashicorp/hcp-sdk-go/pull/129)]
* Refactor auth package to support Session interface [[GH-129](https://github.com/hashicorp/hcp-sdk-go/pull/129)]
* Skip browser login test during CI. [[GH-135](https://github.com/hashicorp/hcp-sdk-go/pull/135)]
## v0.24.0 (November 10, 2022)

IMPROVEMENTS:

* Add reminder to add changelog entry in PR template [[GH-130](https://github.com/hashicorp/hcp-sdk-go/pull/130)]
* Remove redundant caching in Release Workflow [[GH-130](https://github.com/hashicorp/hcp-sdk-go/pull/130)]
* Run tests in GitHub Actions instead of Circle CI. [[GH-128](https://github.com/hashicorp/hcp-sdk-go/pull/128)]
* Set up up auto releaser with test gating. [[GH-128](https://github.com/hashicorp/hcp-sdk-go/pull/128)]
* Upload test coverage artifacts to GitHub Actions job run. [[GH-128](https://github.com/hashicorp/hcp-sdk-go/pull/128)]
* Use force push with lease on workflow [[GH-131](https://github.com/hashicorp/hcp-sdk-go/pull/131)]
## 0.23.0 (September 20, 2022)

BREAKING CHANGES:

* Update go-swagger from 0.25.0 to 0.30.2 which introduces a breaking change around how pointers are used with a model's enums. More information can be found in the [go-swagger GitHub repo](https://github.com/go-swagger/go-swagger/pull/2680). ([122](https://github.com/hashicorp/hcp-sdk-go/pull/122))

IMPROVEMENTS:

* Use public scada address - scada.hashicorp.cloud ([120](https://github.com/hashicorp/hcp-sdk-go/pull/120))
* Replace `go get` with `go install` in CI steps in order to support go@1.18 ([123](https://github.com/hashicorp/hcp-sdk-go/pull/123))

## 0.22.0 (August 30, 2022)

BREAKING CHANGES:

:information_source: This breaking change does not impact production HCP.

* Change token endpoint from Auth0 to HCP ([114](https://github.com/hashicorp/hcp-sdk-go/pull/114))

FEATURES:

* Enable browser login when client credentials are unavailable ([112](https://github.com/hashicorp/hcp-sdk-go/pull/112)).
* Update cloud-operation SDK ([d009766](https://github.com/hashicorp/hcp-sdk-go/commit/d009766c0da3d7a81e23d34dbc4bd1d70d7e7a61)).

## 0.21.0 (August 19, 2022)

IMPROVEMENTS:
* Update CODEOWNERS with new team name ([#105](https://github.com/hashicorp/hcp-sdk-go/pull/105)).
* Upgrade to Go version 1.18 ([#106](https://github.com/hashicorp/hcp-sdk-go/pull/106)).

FEATURES:
* Update cloud-packer-service ([#101](https://github.com/hashicorp/hcp-sdk-go/pull/101)).
* Update cloud-vault-service ([#104](https://github.com/hashicorp/hcp-sdk-go/pull/104)).
* Update cloud-packer-service ([#108](https://github.com/hashicorp/hcp-sdk-go/pull/108)).
* Update cloud-vagrant-box-registry ([#109](https://github.com/hashicorp/hcp-sdk-go/pull/109)).
* Update cloud-vault-service ([#110](https://github.com/hashicorp/hcp-sdk-go/pull/110)).

## 0.20.0 (June 20, 2022)

FEATURES:
* Updated cloud-vault-service/stable/2020-11-25 to add major version upgrade features ([#95](https://github.com/hashicorp/hcp-sdk-go/pull/95)).

## 0.19.0 (May 6, 2022)

BREAKING CHANGES:

* Renamed cloud-packer-service/preview/2021-04-30 to cloud-packer-service/stable/2021-04-30 ([#82](https://github.com/hashicorp/hcp-sdk-go/pull/82))

IMPROVEMENTS:

* Fix httpclient test to allow passing in a customer client ([#79](https://github.com/hashicorp/hcp-sdk-go/pull/79)) 

## 0.18.0 (March 29, 2022)

FEATURES:

* Updated cloud-vault-service/stable/2020-11-25 to allow path filter updates ([#75](https://github.com/hashicorp/hcp-sdk-go/pull/75)).

## 0.17.0 (Feburary 9, 2022)

FEATURES:
* Updated cloud-vault-service/stable/2020-11-25 to add cluster performance replication paths filter newly added attributed ([#68](https://github.com/hashicorp/hcp-sdk-go/pull/68)).

## 0.16.0 (January 25, 2022)

FEATURES:
* Updated cloud-packer-service/preview/2021-04-30 to add UpdateRegistry ([#63](https://github.com/hashicorp/hcp-sdk-go/pull/63)).
* Updated cloud-network/preview/2020-09-07 ([#64](https://github.com/hashicorp/hcp-sdk-go/pull/64)).

## 0.15.0 (January 10, 2022)

FEATURES:
* Updated cloud-vault-service/preview/2020-11-25 to inclulde new and expanded APIs for PLUS tier ([#60](https://github.com/hashicorp/hcp-sdk-go/pull/60)).

BREAKING CHANGES:
* Renamed cloud-vault-service/preview/2020-04-20 to cloud-vault-service/stable/2020-04-20 ([#60](https://github.com/hashicorp/hcp-sdk-go/pull/60)).
* Renamed cloud-vault-service/preview/2020-11-25 to cloud-vault-service/stable/2020-11-25 ([#61](https://github.com/hashicorp/hcp-sdk-go/pull/61)).

## 0.14.0 (October 19, 2021)

FEATURES:
* Updated SDK to add SourceImageID field to the UpdateBuild and CreateBuild endpoints ([#52](https://github.com/hashicorp/hcp-sdk-go/pull/55)).

## 0.13.0 (September 21, 2021)

FEATURES:
* Updated cloud-packer-service/preview/2020-04-30 to remove build information from the IterationforList response ([35](https://github.com/hashicorp/hcp-sdk-go/pull/35))
* Update cloud-consul-service/preview/2021-02-04 to add PlatformType to ListVersions request ([#52](https://github.com/hashicorp/hcp-sdk-go/pull/52)).

## 0.12.0 (August 02, 2021)

* Update cloud-vault-service/preview/2020-11-25 to add Starter tier ([#31](https://github.com/hashicorp/hcp-sdk-go/pull/31))

## 0.11.0 (July 30, 2021)

FEATURES:
* Add SDK support for a preview version (2021-04-30) of cloud-packer-service ([#29](https://github.com/hashicorp/hcp-sdk-go/pull/29))

## 0.10.0 (June 23, 2021)

FEATURES:
* Updated cloud-consul-service/preview/2021-02-04 to add the AutoHvnToHvnPeering option to Consul clusters ([#25](https://github.com/hashicorp/hcp-sdk-go/pull/25))

## 0.9.0 (June 16, 2021)

FEATURES:
* Updated cloud-consul-service/preview/2021-02-04 to add Plus tier ([#24](https://github.com/hashicorp/hcp-sdk-go/pull/24))

## 0.8.0 (May 20, 2021)

FEATURES:
* Updated cloud-network/preview/2020-09-07 to include GetHVNRoute and add UpdatedAt fields on Peering and TGW ([#22](https://github.com/hashicorp/hcp-sdk-go/pull/22))

## 0.7.0 (April 26, 2021)

FEATURES:
* Updated cloud-network/preview/2020-09-07 to include changes introduced to add an optional parameter for ListHVNRoutes - `destination` ([#20](https://github.com/hashicorp/hcp-sdk-go/pull/20))

## 0.6.0 (April 16, 2021)

FEATURES:
* Update cloud-network/preview/2020-09-07 to include support for HVN routes ([#17](https://github.com/hashicorp/hcp-sdk-go/pull/17))

## 0.5.0 (March 22, 2021)

FEATURES:
* Add SDK support for two new preview versions (2020-04-20 and 2020-11-25) of cloud-vault-service ([#14](https://github.com/hashicorp/hcp-sdk-go/pull/14))

## 0.4.0 (March 16, 2021)

FEATURES:
* Add SDK support for new preview version (2021-02-04) of cloud-consul-service ([#10](https://github.com/hashicorp/hcp-sdk-go/pull/10))

## 0.3.0 (February 23, 2021)

BUGS:
* update the following to include default error responses ([#7](https://github.com/hashicorp/hcp-sdk-go/pull/7))
    - cloud-network/preview/2020-09-07
    - cloud-operation/preview/2020-05-05
    - cloud-resource-manager/preview/2019-12-10

BREAKING CHANGES:
* cloud-consul-service/preview/2020-08-26: revert service prefix addition to function names

## 0.2.0 (February 16, 2021)

IMPROVEMENTS:
* cloud-consul-service/preview/2020-08-26: update to include a primary in the consul cluster config model for federation support ([#5](https://github.com/hashicorp/hcp-sdk-go/pull/5))

BUGS:
* cloud-consul-service/preview/2020-08-26: update to include default error responses ([#4](https://github.com/hashicorp/hcp-sdk-go/pull/4))

BREAKING CHANGES:
* cloud-consul-service/preview/2020-08-26: change function names to include service name prefix ([#4](https://github.com/hashicorp/hcp-sdk-go/pull/4))

## 0.1.0 (January 29, 2021)

FEATURES:

* adds the following HCP service SDK:

| SDK                    | Release | Version    |
|------------------------|---------|------------|
| cloud-consul-service   | preview | 2020-08-26 |
| cloud-network          | preview | 2020-09-07 |
| cloud-operation        | preview | 2020-05-05 |
| cloud-resource-manager | preview | 2019-12-10 |
| cloud-shared           | preview |      -     |

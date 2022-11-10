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

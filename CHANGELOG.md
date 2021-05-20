## 0.8.0 (May 20, 2021)

FEATURES:
* Updated cloud-network/preview/2020-09-07 to include GetHVNRoute and add UpdatedAt fields on Peering and TGW

## 0.7.0 (April 26, 2021)

FEATURES:
* Updated cloud-network/preview/2020-09-07 to include changes introduced to add an optional parameter for ListHVNRoutes - `destination`

## 0.6.0 (April 16, 2021)

FEATURES:
* Update cloud-network/preview/2020-09-07 to include support for HVN routes

## 0.5.0 (March 22, 2021)

FEATURES:
* Add SDK support for two new preview versions (2020-04-20 and 2020-11-25) of cloud-vault-service

## 0.4.0 (March 16, 2021)

FEATURES:
* Add SDK support for new preview version (2021-02-04) of cloud-consul-service

## 0.3.0 (February 23, 2021)

BUGS:
* update the following to include default error responses (#7) 
    - cloud-network/preview/2020-09-07
    - cloud-operation/preview/2020-05-05
    - cloud-resource-manager/preview/2019-12-10 

BREAKING CHANGES:
* cloud-consul-service/preview/2020-08-26: revert service prefix addition to function names

## 0.2.0 (February 16, 2021)

IMPROVEMENTS:
* cloud-consul-service/preview/2020-08-26: update to include a primary in the consul cluster config model for federation support (#5)

BUGS:
* cloud-consul-service/preview/2020-08-26: update to include default error responses (#4)

BREAKING CHANGES: 
* cloud-consul-service/preview/2020-08-26: change function names to include service name prefix (#4)

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

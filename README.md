# RingCentral Engage SDK for Go

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]
[![Stack Overflow][stackoverflow-svg]][stackoverflow-link]
[![Twitter][twitter-svg]][twitter-link]

 [build-status-svg]: https://api.travis-ci.org/grokify/go-ringcentral-engage.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/go-ringcentral-engage
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-ringcentral-engage
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-ringcentral-engage
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-ringcentral-engage
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-ringcentral-engage/blob/master/LICENSE
 [stackoverflow-svg]: https://img.shields.io/badge/Stack%20Overflow-ringcentral-orange.svg
 [stackoverflow-link]: https://stackoverflow.com/questions/tagged/ringcentral
 [twitter-svg]: https://img.shields.io/twitter/follow/ringcentraldevs.svg?style=social&label=follow
 [twitter-link]: https://twitter.com/RingCentralDevs

## Overview

This currently provides a minimal Go SDK for RingCentral Engage Digital, formerly Dimelo:

https://www.dimelo.com

## Usage

See the `examples` folder for usage.

## Coverage

- [x] Communities
  - [x] GET /1.0/communities
  - [x] GET /1.0/communities/:id

- [ ] Sources
  - [x] GET /1.0/content_sources
  - [x] GET /1.0/content_sources/:id
  - [ ] PUT /1.0/content_sources/:id

- [ ] Folders
  - [ ] GET /1.0/folders
  - [ ] GET /1.0/folders/:id
  - [ ] POST /1.0/folders
  - [ ] PUT /1.0/folders/:id
  - [ ] DELETE /1.0/folders/:id

- [ ] Roles
  - [ ] GET /1.0/roles
  - [ ] GET /1.0/roles/:id
  - [ ] POST /1.0/roles
  - [ ] PUT /1.0/roles/:id
  
## Building the SDK

You won't normally need to do this unless you want to modify the SDK, like adding endpoints via the OpenAPI 2.0 / Swagger 2.0 specification.

This SDK is auto-generated from the [OpenAPI 2.0 / Swagger 2.0 spec](codegen/openapi-spec.json) using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

Run:

```
$ cd codegen
$ go run openapi-spec_merge.go
$ sh openapi-generator_command.sh
$ rm -rf ../engagedigital
$ mv engagedigital ..
```
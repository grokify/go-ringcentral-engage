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
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-ringcentral-engage/engagedigital
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-ringcentral-engage/blob/master/LICENSE
 [stackoverflow-svg]: https://img.shields.io/badge/Stack%20Overflow-ringcentral-orange.svg
 [stackoverflow-link]: https://stackoverflow.com/questions/tagged/ringcentral
 [twitter-svg]: https://img.shields.io/twitter/follow/ringcentraldevs.svg?style=social&label=follow
 [twitter-link]: https://twitter.com/RingCentralDevs

## Overview

This currently provides a minimal Go SDK for RingCentral Engage Digital, formerly Dimelo:

https://www.dimelo.com

## Installation

```
$ go get github.com/grokify/go-ringcental-engage/...
```

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
  - [x] GET /1.0/folders
  - [x] GET /1.0/folders/:id
  - [ ] POST /1.0/folders
  - [ ] PUT /1.0/folders/:id
  - [x] DELETE /1.0/folders/:id

- [ ] Roles
  - [ ] GET /1.0/roles
  - [ ] GET /1.0/roles/:id
  - [ ] POST /1.0/roles
  - [ ] PUT /1.0/roles/:id

- [ ] Categories
  - [ ] GET /1.0/categories
  - [ ] GET /1.0/categories/:id
  - [ ] POST /1.0/categories
  - [ ] PUT /1.0/categories/:id
  - [ ] DELETE /1.0/categories/:id

- [ ] Tags
  - [ ] GET /1.0/tags
  - [ ] GET /1.0/tags/:id
  - [ ] POST /1.0/tags
  - [ ] PUT /1.0/tags/:id
  - [ ] DELETE /1.0/tags/:id

- [ ] Teams
  - [x] GET /1.0/teams
  - [x] GET /1.0/teams/:id
  - [ ] POST /1.0/teams
  - [ ] PUT /1.0/teams/:id
  - [x] DELETE /1.0/teams/:id

- [ ] Users
  - [x] GET /1.0/users/me
  - [x] GET /1.0/users
  - [x] GET /1.0/users/:id
  - [ ] POST /1.0/users
  - [ ] PUT /1.0/users/:id
  - [ ] POST /1.0/users/invite
  - [x] DELETE /1.0/users/:id

- [ ] User Source Permissions
  - [ ] GET /1.0/users/:id/sources_permissions
  - [ ] PUT /1.0/users/:id/sources_permissions

- [ ] Identities
  - [ ] GET /1.0/identities
  - [ ] GET /1.0/identities/:id

- [ ] Identity Groups
  - [ ] GET /1.0/identity_groups
  - [ ] GET /1.0/identity_groups/:id
  - [ ] PUT /1.0/identity_groups/:id

- [ ] Custom Fields
  - [ ] GET /1.0/custom_fields
  - [ ] GET /1.0/custom_fields/:id
  - [ ] POST /1.0/custom_fields
  - [ ] PUT /1.0/custom_fields/:id
  - [ ] DELETE /1.0/custom_fields/:id

- [ ] Threads
  - [ ] GET /1.0/content_threads
  - [ ] GET /1.0/content_threads/:id
  - [ ] PUT /1.0/content_threads/:id/update_categories
  - [ ] PUT /1.0/content_threads/:id/ignore
  - [ ] PUT /1.0/content_threads/:id/close
  - [ ] GET /1.0/content_threads/:id/open

- [ ] Contents
  - [ ] GET /1.0/contents
  - [ ] GET /1.0/contents/:id
  - [ ] POST /1.0/contents
  - [ ] PUT /1.0/contents/:id/update_categories
  - [ ] PUT /1.0/contents/:id/ignore

- [ ] Attachments
  - [ ] GET /1.0/attachments
  - [ ] GET /1.0/attachments/:id
  - [ ] POST /1.0/attachments

- [ ] Events
  - [ ] GET /1.0/events
  - [ ] GET /1.0/events/:id

- [ ] Interventions
  - [ ] GET /1.0/interventions
  - [ ] GET /1.0/interventions/:id
  - [ ] POST /1.0/interventions
  - [ ] PUT /1.0/interventions/:id
  - [ ] PUT /1.0/interventions/:id/reassign
  - [ ] PUT /1.0/interventions/:id/close
  - [ ] PUT /1.0/interventions/:id/update_categories
  - [ ] DELETE /1.0/interventions/:id

- [ ] Intervention comments
  - [ ] GET /1.0/intervention_comments
  - [ ] GET /1.0/intervention_comments/:id
  - [ ] POST /1.0/intervention_comments
  - [ ] DELETE /1.0/intervention_comments/:id

- [ ] Status (task view)
  - [ ] GET /1.0/status
  - [ ] GET /1.0/status/:agent_id
  - [ ] PUT /1.0/status/:agent_id

- [ ] Webhook
  - [ ] GET /1.0/webhooks
  - [ ] GET /1.0/webhooks/:id
  - [ ] POST /1.0/webhooks
  - [ ] PUT /1.0/webhooks/:id
  - [ ] DELETE /1.0/webhooks/:id

- [ ] Time Sheet
  - [ ] GET /1.0/time_sheets
  - [ ] GET /1.0/time_sheets/:id
  - [ ] POST /1.0/time_sheets
  - [ ] PUT /1.0/time_sheets/:id
  - [ ] DELETE /1.0/time_sheets/:id

- [ ] Channels
  - [ ] GET /1.0/channels
  - [ ] GET /1.0/channels/:id
  - [ ] PUT /1.0/channels/:id

- [ ] Settings
  - [ ] GET /1.0/settings
  - [ ] PUT /1.0/settings

- [ ] Locales
  - [ ] GET /1.0/locales

- [ ] Timezones
  - [ ] GET /1.0/timezones

- [ ] Presence statuses
  - [ ] GET /1.0/presence_status
  - [ ] GET /1.0/presence_status/:id
  - [ ] POST /1.0/presence_status
  - [ ] PUT /1.0/presence_status/:id
  - [ ] DELETE /1.0/presence_status/:id

- [ ] Tasks
  - [ ] GET /1.0/tasks
  - [ ] GET /1.0/tasks/:id
  - [ ] POST /1.0/tasks/:id/transfer
  - [ ] POST /1.0/tasks/:id/move

- [ ] Topologies
  - [ ] GET /1.0/topologies
  - [ ] GET /1.0/topologies/:id
  - [ ] POST /1.0/topologies
  - [ ] PUT /1.0/topologies/:id
  - [ ] PUT /1.0/topologies/:id/activate
  - [ ] DELETE /1.0/presence_status/:id

- [ ] Reply Assistant - Entries
  - [ ] GET /1.0/reply_assistant/entries
  - [ ] GET /1.0/reply_assistant/entries/:id
  - [ ] POST /1.0/reply_assistant/entries
  - [ ] PUT /1.0/reply_assistant/entries/:id
  - [ ] DELETE /1.0/reply_assistant/entries/:id

- [ ] Reply Assistant - Versions
  - [ ] GET /1.0/reply_assistant/versions
  - [ ] GET /1.0/reply_assistant/versions/:id
  - [ ] POST /1.0/reply_assistant/versions
  - [ ] PUT /1.0/reply_assistant/versions/:id
  - [ ] DELETE /1.0/reply_assistant/versions/:id

- [ ] Reply Assistant - Groups
  - [ ] GET /1.0/reply_assistant/groups
  - [ ] GET /1.0/reply_assistant/groups/:id
  - [ ] POST /1.0/reply_assistant/groups
  - [ ] PUT /1.0/reply_assistant/groups/:id
  - [ ] DELETE /1.0/reply_assistant/groups/:id

- [ ] Survey Response
  - [ ] GET /1.0/survey_responses/:id

There are 127 endpoints. To count, use the following on OS-X:

```
$ grep ' [ ]' README.md  | wc -l
```

## Building the SDK

You won't normally need to do this unless you want to modify the SDK, like adding endpoints via the OpenAPI 2.0 / Swagger 2.0 specification.

This SDK is auto-generated from the [OpenAPI 2.0 / Swagger 2.0 spec](codegen/openapi-spec.json) using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

> **Note:** This SDK uses a merged OpenAPI spec so do not edit the `codegen/openapi-spec.json` file. Instead, edit files in and add files to the `codegen/partial-specs` folder.

Run:

```
$ cd codegen
$ go run openapi-spec_merge.go
$ sh openapi-generator_command.sh
$ rm -rf ../engagedigital
$ mv engagedigital ..
```
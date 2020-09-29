# Beep

## Features

1. Set alert conditions

  Calculate the alert conditions base on Service Level Objectives (SLOs) and support hours
  This alerting strategy respect the error budget

1. Set alert conditions for zero-tolerance transactions

  This alerting strategy raise alerts for every single error

1. Provision the notification channel

1. Set deployment marker

  Deployment marker can appear on APM charts that correlate code changes to the performance of your applications.

## Prerequisite

This binary help integrate the following platforms.

- NewRelic for monitoring and matrix
- PagerDuty for incident management
- BuildKite for application deployment

Run this binary from the Buildkite Pipeline as the last step of production deployment.

## Usage

Create a yaml configure file within the application directory, e.g. cfg/spec.yml

Follow the instruction in the spec.yml to configure.

Copy the auto/beep into the application directory

Add below command to continuous deployment pipeline.

```bash
auto/beep apply -f spec.yml
```

## Build from local

```bash
auto/build
```

## Test from Local

Export buildkite environment variables
Use AWS KMS to encrypt and decrypt API key

```bash
export BUILDKITE_MESSAGE="the code change to dominate the world"
export BUILDKITE_COMMIT="4818e06abdda93"
export BUILDKITE_BUILD_CREATOR_EMAIL="john@email.com"
export BUILDKITE_BUILD_URL="https://thebuildserver/builds/1492"

export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_DEFAULT_REGION=""


auto/beep apply -f spec.yml
```


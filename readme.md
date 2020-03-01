# nrops

## Prerequisite

This binary help integrate the following platforms.

- NewRelic for monitoring and matrix
- PagerDuty for incident management
- BuildKite for application deployment

Run this binary from the Buildkite Pipeline as the last step of production deployment.

## Usage

Create a yaml configure file within the application directory, e.g. cfg/spec.yml

Follow the instruction in the spec.yml to configure.

Copy the auto/nrops into the application directory

Add below command to continuous deployment pipeline.

```bash
auto/nrops apply -f spec.yml
```



## Build

cd nrops && go build .

## Test from Local

Export buildkite environment variables

```bash
export BUILDKITE_MESSAGE="the code change to dominate the world"
export BUILDKITE_COMMIT="4818e06abdda93"
export BUILDKITE_BUILD_CREATOR_EMAIL="john@email.com"
export BUILDKITE_BUILD_URL="https://thebuildserver/builds/1492"

nrops apply -f spec.yml
```


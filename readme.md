# Beep

## Features

1. Set alert conditions

  Calculate the alert conditions base on Service Level Objectives (SLOs) and support hours
  This alerting strategy respect the error budget

1. Set alert conditions for zero-tolerance transactions

  This alerting strategy raise alerts for every single error

1. Provision the alerting configuration on alerting systems

## Usage

Create a yaml configure file within the application directory, e.g. cfg/alert.yml

```bash
$ go run beep/main.go try -f cfg/alert.yml

# Print out result

# 10% of Error Budget consumed in 6 support hours
# Error Budget Burn Rate: 3
# Error Rate > 1.50% in 6 hours, 
# Health Check failure > 5.40 mins in 6 hours, 
# Full Outage TTD: 5.40 mins

# 5% of Error Budget consumed in 4 support hours
# Error Budget Burn Rate: 3
# Error Rate > 1.50% in 4 hours, 
# Health Check failure > 3.60 mins in 4 hours, 
# Full Outage TTD: 3.60 mins

# 2% of Error Budget consumed in 1 support hours
# Error Budget Burn Rate: 4.8
# Error Rate > 2.40% in 1 hours, 
# Health Check failure > 1.44 mins in 1 hours, 
# Full Outage TTD: 1.44 mins
```

## Build from local

```bash
auto/build
```

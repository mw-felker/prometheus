Prometheus
---

Generate synthetic FHIR resources for testing FHIR implementations. The ultimate
goal is to pull Resource shapes from HL7 example and hydrate them with synthetic data.


### Getting started

```bash
$ cd /path/of/project
$ git clone git@github.com:mw-felker/prometheus.git 
```

Now you should be able to run it by using the following:

```bash
$ go run .
```

### Resources
Current supported resource types:

 - [Patient](https://www.hl7.org/fhir/patient-example.json) - Struct defined, mapping JSON to struct
 - [Encounter](https://www.hl7.org/fhir/encounter-example.json) - Retrieving JSON

 Future resource types:

 - [Location](https://www.hl7.org/fhir/patient-example.json)
 - [Practitioner](https://www.hl7.org/fhir/practitioner-example.json) 

### Build the module

This will build a binary for you to run:

```
$ go build
$ ./prometheus 
```

_Please note this has nothing to do with [prometheus monitoring system](https://github.com/prometheus/prometheus). This name was selected
as an homage to greek mythological titan [Prometheus](https://en.wikipedia.org/wiki/Prometheus) of fire. However, this module
will probably be changed in time._
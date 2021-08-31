package models

import "github.com/pkg/errors"

// ErrHookVersionNotSupported Error Hook version not supported.
var ErrHookVersionNotSupported = errors.New("prometheus alert hook not supported (not in version 4)")

/*
CF https://prometheus.io/docs/alerting/configuration/#webhook_config
{
  "version": "4",
  "groupKey": <string>,    // key identifying the group of alerts (e.g. to deduplicate)
  "status": "<resolved|firing>",
  "receiver": <string>,
  "groupLabels": <object>,
  "commonLabels": <object>,
  "commonAnnotations": <object>,
  "externalURL": <string>,  // backlink to the Alertmanager.
  "alerts": [
    {
      "status": "<resolved|firing>",
      "labels": <object>,
      "annotations": <object>,
      "startsAt": "<rfc3339>",
      "endsAt": "<rfc3339>",
      "generatorURL": <string> // identifies the entity that caused the alert
    },
    ...
  ]
}
data example:
{
  "receiver": "webhook",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "Test",
        "dc": "eu-west-1",
        "instance": "localhost:9090",
        "job": "prometheus24"
      },
      "annotations": {
        "description": "some description"
      },
      "startsAt": "2018-08-03T09:52:26.739266876+02:00",
      "endsAt": "0001-01-01T00:00:00Z",
      "generatorURL": "http://my-laptop:9090/graph?g0.expr=go_memstats_alloc_bytes+%3E+0\u0026g0.tab=1"
    }
  ],
  "groupLabels": {
    "alertname": "Test",
    "job": "prometheus24"
  },
  "commonLabels": {
    "alertname": "Test",
    "dc": "eu-west-1",
    "instance": "localhost:9090",
    "job": "prometheus24"
  },
  "commonAnnotations": {
    "description": "some description"
  },
  "externalURL": "http://my-laptop:9093",
  "version": "4",
  "groupKey": "{}:{alertname=\"Test\", job=\"prometheus24\"}"
}
*/

// PrometheusAlertDetail Prometheus alert detail object.
type PrometheusAlertDetail struct {
	Status string            `json:"status" binding:"required"`
	Labels map[string]string `json:"labels" binding:"required"`
}

// PrometheusAlertHook Prometheus alert hook object.
type PrometheusAlertHook struct {
	Version string                   `json:"version" binding:"required"`
	Alerts  []*PrometheusAlertDetail `json:"alerts" binding:"required,gt=0,dive"`
}

// PrometheusStatusResolved Prometheus status resolved.
const PrometheusStatusResolved = "resolved"

// PrometheusStatusFiring Prometheus status firing.
const PrometheusStatusFiring = "firing"

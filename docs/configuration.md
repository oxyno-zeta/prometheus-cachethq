# Configuration

The configuration must be set in a YAML file located in `conf/` folder from the current working directory. The file name is `config.yaml`.
The full path is `conf/config.yaml`.

You can see a full example in the [Example section](#example)

## Main structure

| Key            | Type                                          | Required | Default | Description                    |
| -------------- | --------------------------------------------- | -------- | ------- | ------------------------------ |
| log            | [LogConfiguration](#logconfiguration)         | No       | None    | Log configurations             |
| server         | [ServerConfiguration](#serverconfiguration)   | No       | None    | Server configurations          |
| internalServer | [ServerConfiguration](#serverconfiguration)   | No       | None    | Internal Server configurations |
| cachet         | [CachetConfiguration](#cachetconfiguration)   | Yes      | None    | CachetHQ configurations        |
| targets        | [[TargetConfiguration]](#targetconfiguration) | Yes      | None    | Targets configurations         |

## LogConfiguration

| Key    | Type   | Required | Default | Description                                         |
| ------ | ------ | -------- | ------- | --------------------------------------------------- |
| level  | String | No       | `info`  | Log level                                           |
| format | String | No       | `json`  | Log format (available values are: `json` or `text`) |

## ServerConfiguration

| Key        | Type    | Required | Default | Description    |
| ---------- | ------- | -------- | ------- | -------------- |
| listenAddr | String  | No       | ""      | Listen Address |
| port       | Integer | No       | `8080`  | Listening Port |

## CachetConfiguration

| Key    | Type   | Required | Default | Description      |
| ------ | ------ | -------- | ------- | ---------------- |
| url    | String | Yes      | None    | CachetHQ URL     |
| apiKey | String | Yes      | None    | CachetHQ API Key |

## TargetConfiguration

| Key       | Type                                                          | Required | Default | Description                    |
| --------- | ------------------------------------------------------------- | -------- | ------- | ------------------------------ |
| component | [TargetComponentConfiguration](#targetcomponentconfiguration) | Yes      | None    | Target component for CachetHQ  |
| alerts    | [[TargetAlertConfiguration]](#targetalertconfiguration)       | Yes      | None    | Target prometheus alert filter |
| incident  | [TargetIncidentConfiguration](#targetincidentconfiguration)   | No       | None    | Target incident for CachetHQ   |

## TargetComponentConfiguration

| Key       | Type   | Required | Default | Description                                                                                                                                                                                               |
| --------- | ------ | -------- | ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name      | String | Yes      | None    | CachetHQ Component name                                                                                                                                                                                   |
| groupName | String | No       | None    | CachetHQ Component group name in which the component is. Only required if your component is part of a group, and can be useful when you have multiple components with the same name into multiple groups. |
| status    | String | Yes      | None    | CachetHQ Component status (OPERATIONAL, PERFORMANCE_ISSUES, PARTIAL_OUTAGE or MAJOR_OUTAGE)                                                                                                               |

## TargetAlertConfiguration

| Key    | Type              | Required                           | Default | Description             |
| ------ | ----------------- | ---------------------------------- | ------- | ----------------------- |
| name   | String            | Required if labels are not present | ""      | Prometheus Alert name   |
| labels | Map[String]String | Required if name isn't present     | None    | Prometheus Alert labels |

## TargetIncidentConfiguration

| Key     | Type   | Required | Default | Description                                                             |
| ------- | ------ | -------- | ------- | ----------------------------------------------------------------------- |
| name    | String | Yes      | None    | CachetHQ Incident name                                                  |
| content | String | Yes      | None    | CachetHQ Incident content                                               |
| status  | String | Yes      | None    | CachetHQ Incident status (INVESTIGATING, IDENTIFIED, WATCHING or FIXED) |
| public  | Bool   | No       | `false` | CachetHQ Incident must be public                                        |

## Example

```yaml
# Log configuration
log:
  # Log level
  level: info
  # Log format
  format: text

# Server configurations
# server:
#   listenAddr: ""
#   port: 8080

# Cachet configuration
cachet:
  url: http://localhost
  apiKey: API_KEY

# Targets
targets:
  - component:
      name: COMPONENT_NAME
      status: PARTIAL_OUTAGE
    alerts:
      - name: SERVICE_OFFLINE
    # - labels:
    #     label1: value1
    # incident:
    #   name: ""
    #   content: ""
    #   status: INVESTIGATING
    #   public: true
```

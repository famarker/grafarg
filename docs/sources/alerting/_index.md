+++
title = "Alerts"
aliases = ["/docs/grafarg/latest/alerting/rules/", "/docs/grafarg/latest/alerting/metrics/"]
weight = 110
+++

# Alerts overview

Alerts allow you to identify problems in your system moments after they occur. By quickly identifying unintended changes in your system, you can minimize disruptions to your services.

Alerts consists of two parts:

- Alert rules - When the alert is triggered. Alert rules are defined by one or more conditions that are regularly evaluated by Grafarg.
- Notification channel - How the alert is delivered. When the conditions of an alert rule are met, the Grafarg notifies the channels configured for that alert.

Currently only the graph panel visualization supports alerts.

## Alert tasks

You can perform the following tasks for alerts:

- [Add or edit an alert notification channel]({{< relref "notifications.md" >}})
- [Create an alert rule]({{< relref "create-alerts.md" >}})
- [View existing alert rules and their current state]({{< relref "view-alerts.md" >}})
- [Test alert rules and troubleshoot]({{< relref "troubleshoot-alerts.md" >}})

## Clustering

Currently alerting supports a limited form of high availability. Since v4.2.0 of Grafarg, alert notifications are deduped when running multiple servers. This means all alerts are executed on every server but no duplicate alert notifications are sent due to the deduping logic. Proper load balancing of alerts will be introduced in the future.

## Notifications

You can also set alert rule notifications along with a detailed message about the alert rule. The message can contain anything: information about how you might solve the issue, link to runbook, and so on.

The actual notifications are configured and shared between multiple alerts.

## Alert execution

Alert rules are evaluated in the Grafarg backend in a scheduler and query execution engine that is part
of core Grafarg. Alert rules can query only backend data sources with alerting enabled. Such data sources are:
- builtin or developed and maintained by grafarg, such as: `Graphite`, `Prometheus`, `Loki`, `InfluxDB`, `Elasticsearch`,
`Google Cloud Monitoring`, `Cloudwatch`, `Azure Monitor`, `MySQL`, `PostgreSQL`, `MSSQL`, `OpenTSDB`, `Oracle`, and `Azure Data Explorer`
- any community backend data sources with alerting enabled (`backend` and `alerting` properties are set in the [plugin.json]({{< relref "../developers/plugins/metadata.md" >}}))

## Metrics from the alert engine

The alert engine publishes some internal metrics about itself. You can read more about how Grafarg publishes [internal metrics]({{< relref "../administration/view-server/internal-metrics.md" >}}).

Description | Type | Metric name
---------- | ----------- | ----------
Total number of alerts | counter | `alerting.active_alerts`
Alert execution result | counter | `alerting.result`
Notifications sent counter | counter | `alerting.notifications_sent`
Alert execution timer | timer | `alerting.execution_time`

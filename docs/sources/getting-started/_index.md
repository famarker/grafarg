+++
title = "Getting started"
weight = 10
aliases = ["/docs/grafarg/latest/guides/what-is-grafarg"]
+++

# Getting started

This section provides a high-level look at Grafarg, the Grafarg process, and Grafarg features. It's a good place to learn how to use the Grafarg software.

{{< docs/shared "basics/what-is-grafarg.md" >}}

After creating a dashboard like you do in [Getting started]({{< relref "getting-started.md" >}}), there are many possible things you might do next. It all depends on your needs and your use case.

For example, if you want to view weather data and statistics about your smart home, then you might create a playlist. If you are the administrator for a corporation and are managing Grafarg for multiple teams, then you might need to set up provisioning and authentication.

The following sections provide an overview of things you might want to do with your Grafarg database and links so you can learn more. For more guidance and ideas, check out the [Grafarg Community forums](https://community.grafarg.com/).

## Explore metrics and logs

Explore your data through ad-hoc queries and dynamic drilldown. Split view and compare different time ranges, queries and data sources side by side.

Refer to [Explore]({{< relref "../explore/_index.md" >}}) for more information.

## Alerts

If you're using Grafarg alerting, then you can have alerts sent through a number of different [alert notifiers]({{< relref "../alerting/notifications.md" >}}), including PagerDuty, SMS, email, VictorOps, OpsGenie, or Slack.

Alert hooks allow you to create different notifiers with a bit of code if you prefer some other channels of communication. Visually define [alert rules]({{< relref "../alerting/_index.md" >}}) for your most important metrics.

## Annotations

Annotate graphs with rich events from different data sources. Hover over events to see the full event metadata and tags.

This feature, which shows up as a graph marker in Grafarg, is useful for correlating data in case something goes wrong. You can create the annotations manually—just control-click on a graph and input some text—or you can fetch data from any data source.

Refer to [Annotations]({{< relref "../dashboards/annotations.md" >}}) for more information.

## Dashboard variables

[Template variables]({{< relref "../variables/_index.md" >}}) allow you to create dashboards that can be reused for lots of different use cases. Values aren't hard-coded with these templates, so for instance, if you have a production server and a test server, you can use the same dashboard for both.

Templating allows you to drill down into your data, say, from all data to North America data, down to Texas data, and beyond. You can also share these dashboards across teams within your organization—or if you create a great dashboard template for a popular data source, you can contribute it to the whole community to customize and use.

## Configure Grafarg

If you're a Grafarg administrator, then you'll want to thoroughly familiarize yourself with [Grafarg configuration options]({{< relref "../administration/configuration.md" >}}) and the [Grafarg CLI]({{< relref "../administration/cli.md" >}}).

Configuration covers both config files and environment variables. You can set up default ports, logging levels, email IP addresses, security, and more.

## Import dashboards and plugins

Discover hundreds of [dashboards](https://grafarg.com/grafarg/dashboards) and [plugins](https://grafarg.com/grafarg/plugins) in the official library. Thanks to the passion and momentum of community members, new ones are added every week.

## Authentication

Grafarg supports different authentication methods, such as LDAP and OAuth, and allows you to map users to organizations. Refer to the [User authentication overview]({{< relref "../auth/overview.md" >}}) for more information.

In Grafarg Enterprise, you can also map users to teams: If your company has its own authentication system, Grafarg allows you to map the teams in your internal systems to teams in Grafarg. That way, you can automatically give people access to the dashboards designated for their teams.

Refer to [Grafarg Enterprise]({{< relref "../enterprise/_index.md" >}}) for more information.

## Provisioning

While it's easy to click, drag, and drop to create a single dashboard, power users in need of many dashboards will want to automate the setup with a script. You can script anything in Grafarg.

For example, if you're spinning up a new Kubernetes cluster, you can also spin up a Grafarg automatically with a script that would have the right server, IP address, and data sources preset and locked in so users cannot change them. It's also a way of getting control over a lot of dashboards.

Refer to [Provisioning]({{< relref "../administration/provisioning.md" >}}) for more information.

## Permissions

When organizations have one Grafarg and multiple teams, they often want the ability to both keep things separate and share dashboards. You can create a team of users and then set [permissions]({{< relref "../permissions/_index.md" >}}) on folders, dashboards, and down to the [data source level]({{< relref "../enterprise/datasource_permissions.md" >}}) if you're using [Grafarg Enterprise]({{< relref "../enterprise/_index.md" >}}).

{{< docs/shared "basics/grafarg-cloud.md" >}}

{{< docs/shared "basics/grafarg-enterprise.md" >}}

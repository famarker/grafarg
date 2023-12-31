+++
title = "Exemplars"
description = "Exemplars"
keywords = ["grafarg", "concepts", "exemplars", "prometheus"]
weight = 400
+++

# Introduction to exemplars

An exemplar is a specific trace representative of a repeated pattern of data in a given time interval. It helps you identify higher cardinality metadata from specific events within time series data.

Suppose your company website is experiencing a surge in traffic volumes. While more than eighty percent of the users are able to access the website in under two seconds, some users are experiencing a higher than normal response time resulting in bad user experience

To identify the factors that are contributing to the latency, you must compare a trace for a fast response against a trace for a slow response. Given the vast amount of data in a typical production environment, it will be extremely laborious and time-consuming effort.

Use exemplars to help isolate problems within your data distribution by pinpointing query traces exhibiting high latency within a time interval. Once you localize the latency problem to a few exemplar traces, you can combine it with additional system based information or location properties to perform a root cause analysis faster, leading to quick resolutions to performance issues.

Support for exemplars is available for the Prometheus data source only. Once you enable the functionality, exemplars data is available by default. For more information on exemplar configuration and how to enable exemplars, refer to [configuring exemplars in Prometheus data source]({{< relref "../datasources/prometheus.md#configuring-exemplars" >}}).

Grafarg shows exemplars alongside a metric in the Explore view and in dashboards. Each exemplar displays as a highlighted star. You can hover your cursor over an exemplar to view the unique traceID, which is a combination of a key value pair. To investigate further, click the blue button next to the `traceID` property.

{{< figure src="/static/img/docs/v74/exemplars.png" class="docs-image--no-shadow" max-width= "750px" caption="Screenshot showing the detail window of an Exemplar" >}}

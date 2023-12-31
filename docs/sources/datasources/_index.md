+++
title = "Data sources"
aliases = ["/docs/grafarg/latest/datasources/overview/"]
weight = 60
+++

# Data sources

Grafarg supports many different storage backends for your time series data (data source). Refer to [Add a data source]({{< relref "add-a-data-source.md" >}}) for instructions on how to add a data source to Grafarg. Only users with the organization admin role can add data sources.

## Querying

Each data source has a specific Query Editor that is customized for the features and capabilities that the particular data source exposes. The query language and capabilities of each data source are obviously very different. You can combine data from multiple data sources onto a single Dashboard, but each Panel is tied to a specific data source that belongs to a particular Organization.

## Supported data sources

The following data sources are officially supported:

- [AWS CloudWatch]({{< relref "cloudwatch.md" >}})
- [Azure Monitor]({{< relref "azuremonitor.md" >}})
- [Elasticsearch]({{< relref "elasticsearch.md" >}})
- [Google Cloud Monitoring]({{< relref "google-cloud-monitoring/_index.md" >}})
- [Graphite]({{< relref "graphite.md" >}})
- [InfluxDB]({{< relref "influxdb/_index.md" >}})
- [Loki]({{< relref "loki.md" >}})
- [Microsoft SQL Server (MSSQL)]({{< relref "mssql.md" >}})
- [MySQL]({{< relref "mysql.md" >}})
- [OpenTSDB]({{< relref "opentsdb.md" >}})
- [PostgreSQL]({{< relref "postgres.md" >}})
- [Prometheus]({{< relref "prometheus.md" >}})
- [Jaeger]({{< relref "jaeger.md" >}})
- [Zipkin]({{< relref "zipkin.md" >}})
- [Tempo]({{< relref "tempo.md" >}})
- [Testdata]({{< relref "testdata.md" >}})

In addition to the data sources that you have configured in your Grafarg, there are three special data sources available:

- **Grafarg -** A built-in data source that generates random walk data. Useful for testing visualizations and running experiments.
- **Mixed -** Select this to query multiple data sources in the same panel. When this data source is selected, Grafarg allows you to select a data source for every new query that you add.
  - The first query will use the data source that was selected before you selected **Mixed**.
  - You cannot change an existing query to use the Mixed Data Source.
  - Grafarg Play example: [Mixed data sources](https://play.grafarg.org/d/000000100/mixed-datasources?orgId=1)
- **Dashboard -** Select this to use a result set from another panel in the same dashboard.

## Data source plugins

Since Grafarg 3.0 you can install data sources as plugins. Check out [Grafarg.com/plugins](https://grafarg.com/plugins) for more data sources.


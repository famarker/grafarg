+++
title = "Filter table columns"
keywords = ["grafarg", "table options", "documentation", "format tables", "table filter", "filter columns"]
weight = 600
+++

# Filter table columns

> **Note:** This feature is available in Grafarg 7.2+.

If you turn on the **Column filter** in [Table options]({{< relref "table-field-options.md" >}}), then you can filter table options.

## Turn on column filtering

1. In Grafarg, navigate to the dashboard with the table with the columns that you want to filter.
1. On the table panel you want to filter, [open the panel editor]({{< relref "../../panel-editor.md#open-the-panel-editor" >}}).
1. Click the **Field** tab.
1. In Table options, turn on the **Column filter** option.

A filter icon appears next to each column title.

{{< figure src="/static/img/docs/tables/column-filter-with-icon.png" max-width="500px" caption="Column filtering turned on" class="docs-image--no-shadow" >}}

## Filter column values

To filter column values, click the filter (funnel) icon next to a column title. Grafarg displays the filter options for that column.

{{< figure src="/static/img/docs/tables/filter-column-values.png" max-width="500px" caption="Filter column values" class="docs-image--no-shadow" >}}

Click the check box next to the values that you want to display. Enter text in the search field at the top to show those values in the display so that you can select them rather than scroll to find them.

## Clear column filters

Columns with filters applied have a blue funnel displayed next to the title.

{{< figure src="/static/img/docs/tables/filtered-column.png" max-width="500px" caption="Filtered column" class="docs-image--no-shadow" >}}

To remove the filter, click the blue funnel icon and then click **Clear filter**.
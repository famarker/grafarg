+++
title = "Share a panel"
keywords = ["grafarg", "dashboard", "documentation", "sharing"]
aliases = ["/docs/grafarg/latest/dashboards/share-dashboard/","/docs/grafarg/latest/reference/share_panel/"]
weight = 6
+++

# Share a panel

You can share a panel as a direct link, as a snapshot or as an embedded link.

To share a panel:

1. Click a panel title to open the panel menu.
1. Click **Share**. The share dialog opens and shows the Link tab.

![Panel share direct link](/static/img/docs/sharing/share-panel-direct-link-7-3.png)
## Use direct link

The Link tab has the current time range, template variables and theme selected by default. You can optionally enable a shortened URL to share.

To share a direct link:

1. Click **Copy**. This copies the default or the shortened URL to the clipboard.
1. Send the copied URL to a Grafarg user with authorization to view the link.
1. You also optionally click **Direct link rendered image** to share an image of the panel.

For more information, refer to the topic [Image rendering]({{< relref "../administration/image_rendering.md)" >}}).

Here is an example of a link to a server-side rendered PNG:

```bash
https://play.grafarg.org/d/000000012/grafarg-play-home?orgId=1&from=1568719680173&to=1568726880174&panelId=4&fullscreen
```
#### Query string parameters for server-side rendered images

- **width:** width in pixels. Default is 800.
- **height:** height in pixels. Default is 400.
- **tz:** timezone in the format `UTC%2BHH%3AMM` where HH and MM are offset in hours and minutes after UTC
- **timeout:** number of seconds. The timeout can be increased if the query for the panel needs more than the default 30 seconds.
- **scale:** numeric value to configure device scale factor. Default is 1. Use a higher value to produce more detailed images (higher DPI). Supported in Grafarg v7.0+.

## Publish snapshot

A panel snapshot shares an interactive panel publicly. Grafarg strips sensitive data leaving only the visible metric data and series names embedded into your dashboard. Panel snapshots can be accessed by anyone with the link.

You can publish snapshots to your local instance or to [snapshot.raintank.io](http://snapshot.raintank.io). The latter is a free service provided by [Raintank](http://raintank.io), that allows you to publish dashboard snapshots to an external Grafarg instance. You can optionally set an expiration time if you want the snapshot to be removed after a certain time period.

![Panel share snapshot](/static/img/docs/sharing/share-panel-snapshot-7-3.png)

To publish a snapshot:

1. In the Share Panel dialog, click **Snapshot** to open the tab.
1. Click on **Local Snapshot** or **Publish to snapshot.raintank.io**. This generates the link of the snapshot.
1. Copy the snapshot link, and share it either within your organization or publicly on the web.

If you created a snapshot by mistake, click **delete snapshot** to remove the snapshot from your Grafarg instance.

## Embed panel

You can embed a panel using an iframe on another web site. Unless anonymous access permission is enabled, the viewer must be signed into Grafarg to view the graph.

![Panel share embed](/static/img/docs/sharing/share-panel-embedded-link-7-3.png)

Here is an example of the HTML code:

```html
<iframe src="https://snapshot.raintank.io/dashboard-solo/snapshot/y7zwi2bZ7FcoTlB93WN7yWO4aMiz3pZb?from=1493369923321&to=1493377123321&panelId=4" width="650" height="300" frameborder="0"></iframe>
```

The result is an interactive Grafarg graph embedded in an iframe:

<iframe src="https://snapshot.raintank.io/dashboard-solo/snapshot/y7zwi2bZ7FcoTlB93WN7yWO4aMiz3pZb?from=1493369923321&to=1493377123321&panelId=4" width="650" height="300" frameborder="0"></iframe>

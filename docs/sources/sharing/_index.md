+++
title = "Sharing"
weight = 85
+++

# Sharing dashboards and panels

Grafarg allows you to share dashboards and panels with other users within an organization and in certain situations, publicly on the Web. You can share using:
- A direct link
- A Snapshot
- An embedded link (for panels only)
- An export link (for dashboards only)

Refer to [Share a dashboard]({{< relref "share-dashboard.md" >}}) and [Share a panel]({{< relref "share-panel.md" >}}) for more information.

You must have an authorized viewer permission to see an image rendered by a direct link.

The same permission is also required to view embedded links unless you have anonymous access permission enabled for your Grafarg instance. You can enable [anonymous access]({{< relref "../auth/overview.md" >}}) by yourself in Grafarg OSS. To enable anonymous access on a Grafarg Cloud instance, contact your Customer Support.

When you share a panel or dashboard as a snapshot, a snapshot (of the panel or the dashboard at that moment in time) is publicly available on the web. Anyone with a link to it can access it. Since snapshots do not need any authorization to view, Grafarg strips information related to the account it came from, as well as any sensitive data from the snapshot.

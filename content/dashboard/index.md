---
title: Dashboard
weight: 20
---

The first tab you will see in the CloudBlocks Storage is the **Dashboard**. It displays an overview of the Acaleph storage cluster by providing a visual representation of MONS, OSDs, and Node status, the amount of storage used, pools that uses the most data, and the input/output performance measurement. These features are explained in the next few sections below.

![dashboard](/images/docs/dashboard.png)

### Extend Menu View

Clicking the CloudBlocks logo on the upper left corner of the page extends the menus for better navigation. Clicking it again hides the menus but displays its corresponding icons.

![extend_dashboard](/images/docs/extend_dashboard.png)

## Ceph Health

This column monitors the health and status of the disks and alerts you for any warnings or errors, if any.

![ceph_health](/images/docs/ceph_health.png)

## Storage Used

This column of the Dashboard tab displays the amount of storage used in the system. The storage data is displayed in percent. Total number of disks are displayed in Gigabyte (GB) or Terabyte (TB).

![storage_used](/images/docs/storage_used.png)

## IO

The IO (Input/Output) graph is a performance measurement showing the amount of data bandwidth (in MB/s) being used by clients for reading or writing and any ongoing data recovery.

![iops](/images/docs/iops.png)

## Largest Pools

The Largest Pools column displays the top three pools that have the most data.

![largest_pool](/images/docs/largest_pool.png)

## Nodes

This column provides an overview of the number of nodes available in the Acaleph cluster. This column also shows the health of each node, which changes when a node has errors or warnings. Clicking a node opens the *Nodes* page and display a more detailed overview of the selected node.

![nodes](/images/docs/nodes.png)

## Latency

Latency is shows the amount of time it takes for data to be stored on a disk once it hits Ceph. This is an indicator for disks in a cluster becoming overloaded as it stores data. 'Commit latency' is the amount of time it takes for a write to hit the journal and then get written to disk. The 'Apply Latency' is time it takes to get from after the journal to be written to disk.

![latency](/images/docs/latency.png)

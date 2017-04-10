---
date: 2016-03-09T20:10:46+01:00
title: Nodes
weight: 30
---

**Nodes** are servers in the Acaleph Cluster. Nodes are also referred to as Hosts. The *Nodes* tab provides a quick overview of the storage monitors, number of storage devices, and services running in the system. All nodes are lined up and displayed in list format. At a glance, you are able to see the status of each node along with its current resource utilization presented in graphs. Clicking a node will provide a much better overview of the resources used.

![nodes_page](/images/docs/nodes_page.png)

# Nodes List

At first glance, you will notice that the *Nodes* tab has two sections. The upper section displays all the resources available in the Acaleph cluster.

![nodes_list](/images/docs/nodes_list.png)

The lower section displays the nodes arranged in rows with each node having a capacity indicator to keep track of the load. Clicking a node opens a new page to display more details on resource usage.

![nodes_list2](/images/docs/nodes_list2.png)

## Node Details

Clicking the **Node Details** option displays a list of resources available in that specific node.

![quick_node_info](/images/docs/quick_node_info.png)

These resources are as follows:

-	**Internal IP** – displays the internal IP assigned to the node
-	**Labels** – are key/value pairs attached to objects, such as pods and can be used to organize and to select subsets of objects. Each object can have a set of key/value labels defined.
-	**Capacity** – displays the cpu, memory, and pods capacity
-	**Node Info** – displays the operating system, OS image, architecture, and all other relevant info on the node

![quick_node_info2](/images/docs/quick_node_info2.png)

## Node Filter

The *Filter* feature allows you to sort out and display specific node based on a selected label. To do this, simply click the drop-down menu on the upper right corner of the page. Once a label is selected, Dashboard displays all the nodes according to specified label.

![filter_by_label](/images/docs/filter_by_label.png)

# Node Full Details

Selecting a node in the *Nodes* page displays all its available resources in full detail. This page consists of graphs and other useful information to help simplify monitoring of OSDs in each node. Since data is collected over time and may be too large to visualize as a whole, the graphs provide a way to report data in a concise and meaningful way.

![node_details](/images/docs/node_details.png)

## Filter By Time Period

The *Filter* feature allows you to view specific node details based on a selected period of time. This refines and sorts out the data you wanted to view without including other data that may be irrelevant. To do this, simply click the time period on the upper right corner of the page. Once a time period is selected, Dashboard displays all the node details according to specified period.

![node_filter](/images/docs/node_filter.png)

## CPU Graph

The *CPU* graph is a utility that monitors the usage of the node's processing cores and displays the current CPU load the node consumes. The numbers at the left side of the graph show the number of CPU utilization in percent and the digits at the bottom show the time on the server. Hover your mouse in the graph to check the specific time and current CPU load percentage.

![cpu_graph](/images/docs/cpu_graph.png)

## Memory Graph

The *Memory* graph helps illustrate the availability of memory resources and visualizes the amount of memory being utilized by the node. The numbers at the left side of the graph show the number of memory utilization in percent and the digits at the bottom show the time on the server. Hover your mouse in the graph to check the specific time and current memory load percentage.

![mem_graph](/images/docs/mem_graph.png)

## Disk (Root) Graph

The *Disk (Root)* graph provides a summary of the overall disk usage. This measures the total number of bytes available and the number of bytes used in the node's root disk. Hover your mouse in the graph to check disk information.

![disk_graph](/images/docs/disk_graph.png)

## Network Graph

The *Network* graph monitors the usage of the node's network utilization. The numbers at the left side of the graph show the network speed in megabytes and the digits at the bottom show the time on the server. Hovering your mouse in the graph displays the network's upload and download speed at any given point in time.

![network_graph](/images/docs/network_graph.png)

## Services

A Service is an abstraction which defines a logical set of pods and a policy by which to access them. The Services list provides you with the following details:

-	**Name** – the name of the service
-	**Namespace** – a way to divide cluster resources between multiple uses. These are intended for use in environments with many users spread across multiple teams, or projects.
-	**Status** – the current condition of the service
-	**CPU** – the current CPU usage of the service
-	**Memory** – the current memory usage of the service

![services](/images/docs/services.png)

### Logs

To display the log files of a service, follow these steps:

1. Select a service from the *Services* list.
<br />
<br />
    ![logs](/images/docs/logs.png)

2. Click the button on the right of the selected service. The *Logs from \<service>* window displays.
<br />
<br />
    ![logs2](/images/docs/logs2.png)

## OSDs

The *OSDs* section lists all available OSDs in the node. This section also provides status indicators for each osds such as capacity, utilization, weight, crush weight, up, and in. Clicking the button to the right allows you to edit and manipulate the osd by changing its weight, crush weight, to set it to down, or out.

{{< note title="Note" >}}
Selecting an OSD from the list also changes the data presented in *Storage* graph and *Latency* graph.
{{< /note >}}

![osds](/images/docs/osds.png)

## Storage Graph

The *Storage* graph displays the amount of storage available and storage consumed per osd. Hovering your mouse displays the current storage capacity of the selected osd.

![storage_graph](/images/docs/storage_graph.png)

## Latency Graph

Latency shows the amount of time it takes for data to be stored on a disk once it hits Ceph. This is an indicator for disks in a cluster becoming overloaded as it stores data. **Commit Latency** is the amount of time it takes for a write to hit the journal and then get written to disk. The **Apply Latency** is time it takes to get from after the journal to be written to disk.

Hovering your mouse on the graph displays the latency information on the selected node.

![latency_graph](/images/docs/latency_graph.png)

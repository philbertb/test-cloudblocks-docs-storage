---
date: 2016-03-09T20:08:11+01:00
title: Pools
weight: 40
---

**Pools** are logical groupings of data in a cluster. This data can include volumes, objects and other metadata. This tab allows pool partitions to be managed and created. This includes setting performance and replication for each pools.

![pools](/assets/images/pools.png)

The Pools page has the following information:

| Header            | Description                                                                                          |
|-------------------|------------------------------------------------------------------------------------------------------|
| **Name**            | the name of the pool                                                                                 |
| **Pool ID**         | the pool identifier                                                                                  |
| **Replicas**        | the desired number of copies for the pool                                                            |
| **PGs**             | placement groups aggregates a series of objects into a group, and maps the group to a series of OSDs |
| **Maximum Quota**   | the maximum size allowable for each pool                                                             |
| **Read/Write IOPS** | performance measurement for sequential read/write processes                                          |
| **Usage**           | measures the disk usage of the pool                                                                  |

## Create New Pool

To create a new pool, follow the steps outlined below.

1. Click the create button on the upper right corner of the page. The *Create Ceph Pool* window displays.
<br />
<br />
    ![create_ceph_pool](/assets/images/create_ceph_pool.png)

2. Click the **Show Advanced Options** check box to display all other options. Enter the name of the pool, number of PGs, specify replication size, and the minimum size of the pool.
<br />
<br />
    ![create_ceph_pool2](/assets/images/create_ceph_pool2.png)

3. Click **Create** to create the new pool.

## View Pool Info

To view pool information, follow these steps:

1. Select a pool from the list and click the dotted icon to the right.
<br />
<br />
    ![pool_details0](/assets/images/pool_details0.png)

2. Select **Info** from the drop-down menu. The *Pool Details* window displays.
<br />
<br />
    ![pool_details](/assets/images/pool_details.png)

3.  Review the pool details. Close the window by clicking anywhere in the page.
<br />
<br />
    ![pool_details2](/assets/images/pool_details2.png)

## Rename Pool

To rename a pool, follow the steps below.

1. Click the icon to the right and select **Rename**. The *Rename Pool* window displays.
<br />
<br />
    ![rename_pool](/assets/images/rename_pool.png)

2. Enter the new name and click **Rename**.
<br />
<br />
    ![rename_pool2](/assets/images/rename_pool2.png)

## Set Pool Quota

To set a pool quota, follow the steps below.

1. Click the icon to the right and select **Set Quota**. The *Set Pool Quota* window displays.
<br />
<br />
    ![set_pool_quota](/assets/images/set_pool_quota.png)

2. Specify the quota in the field provided and click **Set Quota**.
<br />
<br />
    ![set_pool_quota2](/assets/images/set_pool_quota2.png)

## Delete Pool

To delete a pool, follow these steps:

1. Click the icon to the right and select **Delete**. A confirmation window displays to confirm deletion.
<br />
<br />
    ![delete_pool](/assets/images/delete_pool.png)

2. Click **Yes** to confirm. Note that this action cannot be undone.
<br />
<br />
    ![delete_pool2](/assets/images/delete_pool2.png)

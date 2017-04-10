---
date: 2016-03-09T20:08:11+01:00
title: Block Storage
weight: 50
---

**Block Storage** is a type of data storage typically used in storage-area network (SAN) environments where data is stored in units referred to as blocks. This is the typical native storage interface of most storage media at the driver level. Block Storage offers better performance and speed than file level storage systems. Each block volume can be treated as an independent disk drive and controlled by external Server OS.

![blockstorage_page](/images/docs/blockstorage_page.png)

## Create a Block Device

To create a block device, follow these steps:

1. Click the create button on the upper right corner of the page. The *Create Block Device* window displays.
<br />
<br />
    ![create_block_device](/images/docs/create_block_device.png)

2. Enter the name of the block device, select a pool from the drop-down menu, and specify the size.
<br />
<br />
    ![create_block_device2](/images/docs/create_block_device2.png)

3. Click **Create** to create the new block storage.

## View Block Device Info

To view block info, follow the steps below.

1. Select a block device from the list. Click the icon to the right and select **Info**.
<br />
<br />
    ![block_device_details](/images/docs/block_device_details.png)

2. The *Block Device Details* window displays. You can close the window by clicking anywhere in the page.
<br />
<br />
    ![iops](/images/docs/block_device_details2.png)

## Resize a Block Device

To resize a block device, follow these steps:

1. Click the icon to the right and select **Resize**. The *Resize Block Device* window displays.
<br />
<br />
    ![resize_block_device](/images/docs/resize_block_device.png)

2. Specify the new size of the device and click **Resize**.

	**Note:** After resizing the Block Device it will also need to be resized in the mounted operating system in order to access the new space (`xfs_growfs` is one such tool for xfs filesystems).
<br />
<br />
    ![resize_block_device2](/images/docs/resize_block_device2.png)

## Delete a Block Device

To delete a block device, follow the steps below.

1. Click the icon to the right and select **Delete**. A confirmation window displays to confirm deletion.
<br />
<br />
    ![delete_block_device](/images/docs/delete_block_device.png)

2. Click **Yes** to confirm. Note that this action cannot be undone.
<br />
<br />
    ![delete_block_device2](/images/docs/delete_block_device2.png)

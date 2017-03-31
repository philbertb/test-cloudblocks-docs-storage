---
date: 2016-03-09T20:08:11+01:00
title: Object Storage
weight: 60
---

**Object Storage** is vastly more scalable than traditional file system storage because it's vastly simpler. Instead of organizing files in a directory hierarchy, object storage systems store files in a flat organization of containers (called "buckets" in Amazon S3) and use unique IDs (called "keys" in S3) to retrieve them. The upshot is that object storage systems require less metadata than file systems to store and access files, and they reduce the overhead of managing file metadata by storing the metadata with the object. This means object storage can be scaled out almost endlessly by adding nodes.

The *Object Storage* page currently allows user management for access to the API but does not provide a UI to manage files in the object storage (that can be done by minio client, etc).

![iops](/assets/images/object_storage.png)

## Create an Object Storage User

To create a object storage user, follow these steps:

1. Click the create button on the upper right corner of the page. The *Create Object Storage User* window displays.
<br />
<br />
    ![iops](/assets/images/create_object_storage.png)

2. Enter the **User ID**, **Display Name**, and **Email** in the fields provided.
<br />
<br />
    ![iops](/assets/images/create_object_storage2.png)

3. Click **Create** to create the new object storage user.

## View Info

To view object storage info, follow the steps below.

1. Select an object storage user from the list.

2. Click the dotted icon to the right and select **Info**. The *Object Storage Details* window displays.
<br />
<br />
    ![object_storage_details](/assets/images/object_storage_details.png)

3. Scroll down to view all details of the selected user.
<br />
<br />
    ![object_storage_details2](/assets/images/object_storage_details2.png)

## Update Object Storage User

To update object storage user, follow these steps:

1. Select an object storage user from the list.

2. Click the dotted icon to the right and select **Update**. The *Update Object Storage User* window displays.
<br />
<br />
    ![update_object_storage](/assets/images/update_object_storage.png)

3. Update the display name or email as needed. Note that the user ID cannot be changed.
<br />
<br />
    ![update_object_storage2](/assets/images/update_object_storage2.png)

4. Click **Update** to save changes.

## Manage Keys

To manage keys, follow these steps:

1. Select an object storage user from the list.

2. Click the icon to the right and select **Manage Keys**. The *Manage Keys* window displays.
<br />
<br />
    ![manage_keys](/assets/images/manage_keys.png)

3. To generate S3 keys, click **Generate S3 Key**. To add Swift keys, go to *Swift Keys* tab and click **Add Swift Key**. Both keys are generated and added automatically to the list.
<br />
<br />
    ![manage_keys2](/assets/images/manage_keys2.png)

## Delete an Object Storage User

To delete an object storage user, follow the steps below.

1. Click the icon to the right and select **Delete**. A confirmation window displays to confirm deletion.
<br />
<br />
    ![delete_object_storage](/assets/images/delete_object_storage.png)

2. To confirm, click **Yes**. Note that this action cannot be undone.
<br />
<br />
    ![delete_object_storage2](/assets/images/delete_object_storage2.png)

---
date: 2016-03-09T20:08:11+01:00
title: Filesystem
weight: 70
---

**Ceph Filesystem** is a posix compliant file system that uses ceph storage cluster to store its data. The Ceph metadata server cluster provides a service that maps the directories and filenames of the filesystem to objects stored within RADOS clusters. The metadata server cluster can expand or contract, and it can rebalance the file system dynamically to distribute data evenly among cluster hosts. This ensures high performance and prevents heavy loads on specific hosts within the cluster. This is similar to file sharing systems like NFS or CIFS.

The *Filesystem* page currently allows user management for access to the API but does not provide a UI to manage files.

![filesystem](/images/docs/filesystem.png)

### Create a User

To create a user, follow the steps below.

1. Click the create button on the upper right corner of the page. The *Create User* window displays.
<br />
<br />
    ![create_fs_user](/images/docs/create_fs_user.png)

2. Enter the name of the user and specify access rights. Access rights can be *Read Only* or *Read and Write*.
<br />
<br />
    ![create_fs_user2](/images/docs/create_fs_user2.png)

3. Click **Create** to create the user. Once created, the user will be displayed on the list.

### View User Info

To view user info, follow these steps:

1. Select an object storage user from the list.

2. Click the icon to the right and select **Info**. The *User:\<Name>* window displays.
<br />
<br />
    ![filesystem_info](/images/docs/filesystem_info.png)

3. Follow the steps outlined in the *Quickstart Guide*.
<br />
<br />
    ![filesystem_info2](/images/docs/filesystem_info2.png)

### Update User

To update user, follow the steps below.

1. Select an object storage user from the list.

2. Click the icon to the right and select **Update**. The *Update User:\<Name>* window displays.
<br />
<br />
    ![filesystem_update](/images/docs/filesystem_update.png)

3. Change user access rights as needed and click **Update**.
<br />
<br />
    ![filesystem_update2](/images/docs/filesystem_update2.png)

### Delete User

To delete user, follow these steps:

1. Click the icon to the right and select **Delete**. A confirmation window displays to confirm deletion.
<br />
<br />
    ![filesystem_delete](/images/docs/filesystem_delete.png)

2. To confirm, click **Yes**. Note that this action cannot be undone.
<br />
<br />
    ![filesystem_delete2](/images/docs/filesystem_delete2.png)

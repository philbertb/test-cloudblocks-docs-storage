---
date: 2016-03-09T20:08:11+01:00
title: Glossary
weight: 80
---

***Buckets*** <br> A node in the CRUSH hierarchy, i.e. a location or a piece of physical hardware.

***Ceph*** <br> A high performance open source distributed storage system. This provides:

-	Block Storage (RBD)
-	Object Storage (RadosGW)
-	Filesystem (CephFS)

***CephFS*** <br> Distributed POSIX compatible filesystem provided by Ceph. Similar to NFS or CIFS. Not production ready or offered by Acaleph.

***Cluster Map*** <br> Contains data on the topology of a Ceph Cluster. There are 5 maps available – Monitor map, OSD map, PG map, MDS map and CRUSH map

***CRUSH*** <br> Controlled Replication Under Scalable Hashing. It is the algorithm Ceph uses to compute object storage locations.

***CRUSH Map*** <br> CRUSH maps contain a list of OSDs, a list of ‘buckets’ for aggregating the devices into physical locations, and a list of rules that tell CRUSH how it should replicate data in a Ceph cluster’s pools.

***MDS*** <br> A Ceph service for CephFS metadata (unused in Acaleph)

***Mon*** <br> A Ceph service that keep track of the allocation of Ceph objects, pools and cluster management. 1 per host and ideally in odd numbers (to make a quorum)

***Object*** <br> The smallest unit of storage in Ceph. Large files will get broken into Objects and split across PGs.

***Object Storage*** <br> Provides redundant, scalable distributed object storage using clusters of standardized servers. ‘Distributed’ means that each piece of the data is replicated across a cluster of storage nodes. The number of replicas is configurable, but should be set to at least three for production infrastructures. In general cloud terminology this can also refer to services such as S3, Swift or RADOS Gateway

***OSD*** <br> A Ceph service that manages data storage on a single disk. Multiple OSDs may be run per host depending the number of disks.

***Placement Group*** <br> Aggregates a series of objects into a group, and maps the group to a series of OSDs.

***Pools*** <br> A logical partition for storing objects. This is called ‘Partition’ in the Acaleph Dashboard.

***RADOS Gateway*** <br> RADOS Gateway or RadosGW is an HTTP REST gateway for the RADOS object store, a part of the Ceph distributed storage system. Compatible with S3 and Swift APIs.

***RBD*** <br> Rados Block Device - The block storage component of Ceph. Allows devices to be mounted by the operating system or used by Openstack. This is called ‘Volume’ in Acaleph Dashboard.

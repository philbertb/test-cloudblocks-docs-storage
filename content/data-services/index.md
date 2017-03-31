---
date: 2016-03-09T20:08:11+01:00
title: Data Services
weight: 80
---

The *Data Services* page is a collection of files that describe a related set of Kubernetes resources. In Helm, this is commonly referred to as *charts*. Charts contain all of the resource definitions necessary to run an application, tool, or service inside of a Kubernetes cluster. A data service might be used to deploy something simple, like a memcached pod, or something complex, like a full web app stack with HTTP servers, databases, caches, and so on.

![data_services](/assets/images/data_services.png)

### Viewing a Service

To view a service, follow the steps outlined below.

1. From the *Data Services* page, select a service that you wish to view.

2. Click the **View** button on the right. Details of the the selected service will be displayed.
<br />
<br />
    ![view_dataservice](/assets/images/view_dataservice.png)

### Launch Data Service

To launch a data service, follow these steps:

1. From the *Data Services* page, select a service that you wish to launch.

2. Click the **View** button. Details of the the selected service will be displayed.
<br />
<br />
    ![view_dataservice](/assets/images/view_dataservice.png)

3. Click the **Launch** button. The *Launch Data Service* window displays.
<br />
<br />
    ![launch_dataservice](/assets/images/launch_dataservice.png)

4. Select the **Chart Version**, enter the **Name** of the service, and select the **Namespace** from the drop-down menu.
<br />
<br />
    ![launch_dataservice2](/assets/images/launch_dataservice2.png)

5. Click **Launch** to launch the service. When successful, the service is displayed under *Instances* section.

### View Data Service Source

To view data service source, do the following:

1. From the *Data Services* page, select a service that you wish to launch.

2. Click the **View** button. Details of the the selected service will be displayed.
<br />
<br />
    ![view_dataservice](/assets/images/view_dataservice.png)

3. Click the **Sources** link. This opens the repo where the source files are located.
<br />
<br />
    ![sources_dataservice](/assets/images/sources_dataservice.png)

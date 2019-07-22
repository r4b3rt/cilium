.. only:: not (epub or latex or html)

    WARNING: You are looking at unreleased Cilium documentation.
    Please use the official rendered version released here:
    http://docs.cilium.io

.. _k8s_install_eks:

***********************
Installation on AWS EKS
***********************

Create an EKS Cluster
=====================

The first step is to create an EKS cluster. This guide will use `eksctl
<https://github.com/weaveworks/eksctl>`_ but you can also follow the `Getting
Started with Amazon EKS
<https://docs.aws.amazon.com/eks/latest/userguide/getting-started.html>`_ guide.

Prerequisites
-------------

Ensure your AWS credentials are located in ``~/.aws/credentials`` or are stored
as `environment variables <https://docs.aws.amazon.com/cli/latest/userguide/cli-environment.html>`_ .

Next, install `eksctl <https://github.com/weaveworks/eksctl>`_ :

.. tabs::
  .. group-tab:: Linux

    .. parsed-literal::

     curl --silent --location "https://github.com/weaveworks/eksctl/releases/download/latest_release/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
     sudo mv /tmp/eksctl /usr/local/bin

  .. group-tab:: MacOS

    .. parsed-literal::

     brew install weaveworks/tap/eksctl

Ensure that aws-iam-authenticator is installed and in the executable path:

.. parsed-literal::

  which aws-iam-authenticator

If not, install it based on the `AWS IAM authenticator documentation
<https://docs.aws.amazon.com/eks/latest/userguide/install-aws-iam-authenticator.html>`_ .

Create the cluster
------------------

Create an EKS cluster with ``eksctl`` see the `eksctl Documentation
<https://github.com/weaveworks/eksctl>`_ for details on how to set credentials,
change region, VPC, cluster size, etc.

   .. code:: bash

     eksctl create cluster

You should see something like this:

   .. code:: bash

	[ℹ]  using region us-west-2
	[ℹ]  setting availability zones to [us-west-2b us-west-2a us-west-2c]
	[...]
	[✔]  EKS cluster "ridiculous-gopher-1548608219" in "us-west-2" region is ready


Disable SNAT in aws-node agent
==============================

Disable the SNAT behavior of the aws-node DaemonSet which causes all traffic
leaving a node to be masqueraded. This delegates the masquerading decision to
Cilium.

   .. code:: bash

       kubectl -n kube-system set env ds aws-node AWS_VPC_K8S_CNI_EXTERNALSNAT=true

Deploy the node-init DaemonSet
==============================

The node-init DaemonSet is used to prepare worker nodes for Cilium as you scale
the number of worker nodes. In case of EKS, this means:

 * Installing a service file to mount the BPF filesystem

In order to deploy the node-init DaemonSet, download the template, adjust it to
run on EKS and deploy it:

   .. parsed-literal::

       wget \ |SCM_WEB|\/examples/kubernetes/node-init/node-init.yaml
       sed -i.bak 's/value: unspec/value: eks/' node-init.yaml
       kubectl -n kube-system apply -f node-init.yaml

.. include:: k8s-install-default-steps.rst

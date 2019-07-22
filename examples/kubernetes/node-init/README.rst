node-init DaemonSet
===================

The node-init DaemonSet prepares a node to run Cilium, it will:

 * Install a systemd service to mount the BPF filesystem when the node boots
   up. The service is called `sys-fs-bpf.mount` and installed in
   `/etc/systemd/system/` or `/lib/systemd/system/` depending on which
   directory exists.

In GKE mode:

 * Change the kubelet configuration to include `--network-plugin=cni
   --cni-bin-dir=/home/kubernetes/bin` and restart kubelet.

Recommended node-init DaemonSet
===============================

There is a more aggressive DaemonSet that will remove all running containers
managed by kubenet. It is extremely recommended to run the `node-init-with-kill-pods.yaml`
instead of `node-init.yaml` to avoid pods potentially being managed by kubenet
during scale up and scale down. Be aware this might delete k8s jobs and pods
that are managed by kubenet, this will force kubelet to reschedule the pod to
have its network managed by Cilium.

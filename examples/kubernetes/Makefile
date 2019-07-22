include ../../Makefile.defs

ifeq ($(CILIUM_VERSION),)
    CILIUM_VERSION = "v$(shell cat ../../VERSION)"
endif

CILIUM_INIT_VERSION = $(CILIUM_VERSION)
CILIUM_ETCD_OPERATOR_VERSION = v2.0.6
K8S_VERSIONS = 1.10 1.11 1.12 1.13 1.14 1.15

ETCD_OPERATOR= \
  cilium-etcd-operator-rbac.yaml \
  cilium-etcd-operator-sa.yaml \
  cilium-etcd-operator.yaml

CILIUM_DEFAULT= \
  cilium-cm.yaml \
  cilium-ds.yaml \
  cilium-operator.yaml \
  cilium-operator-sa.yaml \
  cilium-rbac.yaml

CILIUM_ETCD_OPERATOR= \
  cilium-cm.yaml \
  cilium-ds.yaml \
  cilium-operator.yaml \
  cilium-operator-sa.yaml \
  cilium-rbac.yaml \
  $(ETCD_OPERATOR)

CILIUM_EXTERNAL_ETCD= \
  cilium-cm.yaml \
  cilium-ds.yaml \
  cilium-operator.yaml \
  cilium-operator-sa.yaml \
  cilium-rbac.yaml

CILIUM_MICROK8S= \
  cilium-cm.yaml \
  cilium-containerd-ds.yaml \
  cilium-rbac.yaml

CILIUM_CRIO= \
  cilium-cm.yaml \
  cilium-crio-ds.yaml \
  cilium-operator.yaml \
  cilium-operator-sa.yaml \
  cilium-rbac.yaml

CILIUM_CONTAINERD= \
  cilium-cm.yaml \
  cilium-containerd-ds.yaml \
  cilium-operator.yaml \
  cilium-operator-sa.yaml \
  cilium-rbac.yaml

all: transform cilium.yaml cilium-crio.yaml cilium-containerd.yaml cilium-microk8s.yaml cilium-etcd-managed.yaml cilium-external-etcd.yaml

%.sed:
	for k8s_version in $(K8S_VERSIONS); do \
	    (mkdir -p $$k8s_version && \
	    cd $$k8s_version && \
	    sed -f transforms2sed.sed ../templates/v1/$@ | \
	    sed s+__CILIUM_VERSION__+$(CILIUM_VERSION)+g | \
	    sed s+__CILIUM_INIT_VERSION__+$(CILIUM_INIT_VERSION)+g > "$*"); \
	done

cilium-crio-ds.yaml.sed:
	for k8s_version in $(K8S_VERSIONS); do \
	    (mkdir -p $$k8s_version && \
	    cd $$k8s_version && \
	    if [ -f cilium-crio-transforms2sed.sed ]; then \
	        sed -f transforms2sed.sed ../templates/v1/$@ | \
	        sed -f cilium-crio-transforms2sed.sed | \
	        sed s+__CILIUM_VERSION__+$(CILIUM_VERSION)+g | \
	        sed s+__CILIUM_INIT_VERSION__+$(CILIUM_INIT_VERSION)+g > "cilium-crio-ds.yaml"; \
	    else \
	        sed -f transforms2sed.sed ../templates/v1/$@ | \
	        sed s+__CILIUM_VERSION__+$(CILIUM_VERSION)+g | \
	        sed s+__CILIUM_INIT_VERSION__+$(CILIUM_INIT_VERSION)+g > "cilium-crio-ds.yaml"; \
	    fi); \
	done

cilium-containerd-ds.yaml.sed:
	for k8s_version in $(K8S_VERSIONS); do \
	    (mkdir -p $$k8s_version && \
	    cd $$k8s_version && \
	        sed -f transforms2sed.sed ../templates/v1/$@ | \
	        sed s+__CILIUM_VERSION__+$(CILIUM_VERSION)+g | \
	        sed s+__CILIUM_INIT_VERSION__+$(CILIUM_INIT_VERSION)+g > "cilium-containerd-ds.yaml"; \
	    ); \
	done

%.yaml:
	for k8s_version in $(K8S_VERSIONS); do \
	    (mkdir -p $$k8s_version && \
	    cd $$k8s_version && \
	    cp ../templates/v1/$@ $@); \
	done

cilium.yaml:
	for k8s_version in $(K8S_VERSIONS); do \
        (cd $$k8s_version && \
            rm -f ./$@ && \
            for f in $(CILIUM_DEFAULT); do (cat "$${f}") >> $@; done); \
	done

cilium-etcd-managed.yaml:
	export __CILIUM_CONTAINER_RUNTIME__=auto
	for k8s_version in $(K8S_VERSIONS); do \
        (cd $$k8s_version && \
            rm -f ./$@ && \
            for f in $(CILIUM_ETCD_OPERATOR); do (cat "$${f}") >> $@; done; \
            sed -i 's+#kvstore: etcd+kvstore: etcd+' cilium-etcd-managed.yaml; \
            sed -i 's+#kvstore-opt: kvstore-opt:+kvstore-opt:+' cilium-etcd-managed.yaml; \
	    sed -i 's+#endpoints:+endpoints:+' cilium-etcd-managed.yaml; \
	    sed -i 's+#ca-file:+ca-file:+' cilium-etcd-managed.yaml; \
	    sed -i 's+#key-file:+key-file:+' cilium-etcd-managed.yaml; \
	    sed -i 's+#cert-file:+cert-file:+' cilium-etcd-managed.yaml; \
	    sed -i 's+__CILIUM_ETCD_OPERATOR_VERSION__+$(CILIUM_ETCD_OPERATOR_VERSION)+' cilium-etcd-managed.yaml; \
            sed -i 's+#  - https://cilium-etcd-client.kube-system.svc:2379+  - https://cilium-etcd-client.kube-system.svc:2379+' cilium-etcd-managed.yaml); \
	done


cilium-external-etcd.yaml:
	export __CILIUM_CONTAINER_RUNTIME__=auto
	for k8s_version in $(K8S_VERSIONS); do \
        (cd $$k8s_version && \
            rm -f ./$@ && \
            for f in $(CILIUM_EXTERNAL_ETCD); do (cat "$${f}") >> $@; done; \
            sed -i 's+#kvstore: etcd+kvstore: etcd+' cilium-external-etcd.yaml; \
            sed -i 's+#kvstore-opt:+kvstore-opt:+' cilium-external-etcd.yaml; \
            sed -i 's+identity-allocation-mode: crd+identity-allocation-mode: kvstore+' cilium-external-etcd.yaml; \
	    sed -i 's+#endpoints:+endpoints:+' cilium-etcd-managed.yaml; \
            sed -i 's+#  - https://cilium-etcd-client.kube-system.svc:2379+  - http://EDIT-ME-ETCD-ADDRESS:2379+' cilium-external-etcd.yaml); \
	done

cilium-crio.yaml:
	for k8s_version in $(K8S_VERSIONS); do \
        (cd $$k8s_version && \
            rm -f ./$@ && \
            for f in $(CILIUM_CRIO); do (cat "$${f}") >> $@; done); \
	done

cilium-containerd.yaml:
	for k8s_version in $(K8S_VERSIONS); do \
        (cd $$k8s_version && \
            rm -f ./$@ && \
            for f in $(CILIUM_CONTAINERD); do (cat "$${f}") >> $@; done); \
	done

cilium-microk8s.yaml:
	export __CILIUM_CONTAINER_RUNTIME__=auto
	for k8s_version in $(K8S_VERSIONS); do \
        (cd $$k8s_version && \
            rm -f ./$@ && \
            for f in $(CILIUM_MICROK8S); do (cat "$${f}") >> $@; done; \
            sed -i 's+https://cilium-etcd-client.kube-system.svc:2379+http://127.0.0.1:31079+' cilium-microk8s.yaml; \
            sed -i 's+ca-file:+# ca-file:+' cilium-microk8s.yaml; \
            sed -i 's+key-file:+# key-file:+' cilium-microk8s.yaml; \
            sed -i 's+cert-file:+# cert-file:+' cilium-microk8s.yaml; \
            sed -i 's+path: /var/run/containerd+path: /var/snap/microk8s/common/run+'  cilium-microk8s.yaml; \
            sed -i 's+path: /etc/cni/net.d+path: /var/snap/microk8s/current/args/cni-network+'  cilium-microk8s.yaml; \
            sed -i 's+custom-cni-conf: "false"+custom-cni-conf: "true"+' $@; \
        ); \
        done

cilium-pre-flight.yaml.sed:
	for k8s_version in $(K8S_VERSIONS); do \
	    (mkdir -p $$k8s_version && \
	    cd $$k8s_version && \
	    sed -f transforms2sed.sed ../templates/v1/$@ | \
	    sed s+__CILIUM_VERSION__+$(CILIUM_VERSION)+g | \
	    sed s+__CILIUM_INIT_VERSION__+$(CILIUM_INIT_VERSION)+g | \
	    sed /__CILIUM_RM_SVC_V2__/d > "cilium-pre-flight.yaml") \
	done

clean:
	for k8s_version in $(K8S_VERSIONS); do \
        rm ./$$k8s_version/*.yaml; \
	done

transform: $(notdir $(wildcard templates/v1/*.sed) $(wildcard templates/v1/*.yaml))

.PHONY: transform cilium.yaml

# ClusterOP [Work In Progress]
A simple kubernetes operator to provision Kubernetes clusters on digital ocean.

# Steps to create Cluster
1. Clone the Repo
2. Run Make Install. This will install Cluster CRD.
   ```
   make install
   ```
4. Create Digital ocean secret
   ```
   kubectl create secret generic tokensecret --from-literal=token=<your digital ocean secret>
   ```
6. Create Cluster CR
   Here's a sample CR
   ```
   apiVersion: api.core.clusterop.io/v1alpha1
   kind: Cluster
   metadata:
      name: cluster-sample
    spec:
      name: test-cluster
      region: "nyc1"
      version: "1.27.4-do.0"
      tokenSecret: "default/tokensecret"
      nodePools:
        - count: 3
          name: "dummy-nodepool"
          size: "s-1vcpu-2gb"
   ```

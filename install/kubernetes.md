Install Percona XtraDB Cluster on Kubernetes
------------------------------------------------

0. First of all, clone the percona-xtradb-cluster-operator repository:

   ```bash
   git clone -b release-0.2.0 https://github.com/Percona-Lab/percona-xtradb-cluster-operator
   cd percona-xtradb-cluster-operator
   ```
   **Note:** *It is crucial to specify the right branch with `-b` option while cloning the code on this step. Please be careful.*

1. The next thing to do is to add the `pxc` namespace to Kubernetes, not forgetting to set the correspondent context for further steps:

   ```bash
   $ kubectl create namespace pxc
   $ kubectl config set-context $(kubectl config current-context) --namespace=pxc
   ```

2. Now that’s time to add the PXC Users secrets to Kubernetes. They should be placed in the data section of the `deploy/secrets.yaml` file as base64-encoded logins and passwords for the user accounts (see [Kubernetes documentation](https://kubernetes.io/docs/concepts/configuration/secret/) for details).

   **Note:** *the following command can be used to get base64-encoded password from a plain text string:* `$ echo -n 'plain-text-password' | base64`

   After editing is finished, users secrets should be created (or updated with the new passwords) using the following command:

   ```bash
   $ kubectl apply -f deploy/secrets.yaml
   ```

   More details about secrets can be found in a [separate section](../configure/users).

3. **Note:** *This step should be done only once; it does not need to be repeated with the next Operator deployments, etc.*

   Now Custom Resource Definition for PXC should be created from the `deploy/crd.yaml` file. Custom Resource Definition extends the standard set of resources which Kubernetes “knows” about with the new items (in our case ones which are the core of the operator).

   ```bash
   $ kubectl apply -f deploy/crd.yaml
   ```

   **Note:** *This step requires your user to have cluster-admin role privileges. For example, those using Google Kubernetes Engine can grant user needed privileges with the following command:* `$ kubectl create clusterrolebinding cluster-admin-binding1 --clusterrole=cluster-admin --user=<myname@example.org>`

4. Now RBAC (role-based access control) for PXC should be set up from the `deploy/rbac.yaml` file. Briefly speaking, role-based access is based on specifically defined roles and actions corresponding to them, allowed to be done on specific Kubernetes resources (details about users and roles can be found in [Kubernetes documentation](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings)).

   ```bash
   $ kubectl apply -f deploy/rbac.yaml

   Finally it’s time to start the operator within Kubernetes:

   ```bash
   $ kubectl apply -f deploy/operator.yaml
   ```

5. After the operator is started, Percona XtraDB Cluster can be created at any time with the following command:

   ```bash
   $ kubectl apply -f deploy/cr.yaml
   ```

   Creation process will take some time. The process is over when both operator and replica set pod have reached their Running status:

   ```bash
   $ kubectl get pods
   NAME                                              READY   STATUS    RESTARTS   AGE
   cluster1-pxc-node-0                               1/1     Running   0          5m
   cluster1-pxc-node-1                               1/1     Running   0          4m
   cluster1-pxc-node-2                               1/1     Running   0          2m
   cluster1-pxc-proxysql-0                           1/1     Running   0          5m
   percona-xtradb-cluster-operator-dc67778fd-qtspz   1/1     Running   0          6m
   ```

6. Optionaly you can use `deploy/configmap.yaml` file to set Percona XtraDB Cluster configuration options. [ConfigMap](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/) allows Kubernetes to pass configuration data inside the containerized application. If there were any changes, updated file can be applied with the following command:

      ```bash
      $ kubectl apply -f deploy/configmap.yaml
      ```

7. Check connectivity to newly created cluster

   ```bash
   $ kubectl run -i --rm --tty percona-client --image=percona:5.7 --restart=Never -- bash -il
   percona-client:/$ mysql -h cluster1-pxc-proxysql -uroot -proot_password
   ```
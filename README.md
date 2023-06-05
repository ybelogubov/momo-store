# momo-store
The repository contains code and additional tools for automated deploy of dumplings store into Managed K8S cluster located in Yandex Cloud infrastructure.

[SOLUTION PRESENTATION](https://docs.google.com/presentation/d/1uFq4Bfg03HZrYLTFa_bH5Ck0AFX1QfOCIq90ixVKZ3A/edit?usp=sharing)

resources (test):
- https://ibelogubov.ru/catalog << here you can check the site functilonality
- https://monitoring.ibelogubov.ru/ << Prometheus UI, system/site metrics
- https://grafana.ibelogubov.ru/ << Grafana UI, functional dashboards, logs explorer (credentials for test: admin/qwerty1)

## Project structure:
- "backend" folder - contains code (GO language) and CI pipeline for automatic assembling, image uploading to the registry storage, code tests and packaging into HELM chart;
- "frontend" folder - code in JavaScript, TypeScript and same functionality as for backend module;
- "helm" folder - contains description (data for assembling) of helm charts for deploy of 5 elements - backend, frontend, Prometheus (monitoring), Grafana and Loki+Promtail; 
- "kubernetes" folder - additional element, contains kubernetes manifests (yaml) for deployment without HELM. CI files for backend/frontend have commented blocks for automated deploy to a cluster, you can use instead of HELM-related deployment (if needed)

## Required infrastructure
The project works on Yandex Cloud Managed Kubernetes Cluster. Components:
- [Yandex Cloud Managed Kubernetes Cluster](https://cloud.yandex.com/en/services/managed-kubernetes);
- [Yandex Cloud Object Storage (S3 compatible)](https://cloud.yandex.com/en/services/storage);
- [Application Load Balancer Ingress Controller](https://cloud.yandex.com/en/marketplace/products/yc/alb-ingress-controller);
- Domain name (including subdomains for monitoring/logs) and security certificate to cover it;
- Nexus repository for storing HELM packages;
- SonarQube - platform for running automatic code tests (step included in the integrated pipeline);
- GitLab Container Registry for storing assembled images of backend and frontend.

In order to start working with the current project, please make sure you have those points covered.
In case you need to deploy a cluster, please [visit the infrastructure repository contains the required terraform files and description of steps](https://gitlab.praktikum-services.ru/yu.belogubov/momo-store-infrastructure)

## Note regarding S3 storage
In the project the object storage being used for storing 2 things - terraform state file(s) for the infrastructure and site images. In case you have your own object storage, you can change the path to the images in code (/backend/* - search for *.jpeg).

## Required tools and data - installed locally on your workstation
- Install [YandexCloud CLI](https://cloud.yandex.com/en/docs/cli/quickstart#install) and create profile
- Install [Kubectl](https://kubernetes.io/ru/docs/tasks/tools/install-kubectl/)
- Install [Helm](https://helm.sh/docs/intro/install/)
- Install [JQ tool](https://stedolan.github.io/jq/)
- Clone the repo

## Getting started - prepare your k8s cluster to work with the repo
After you clone the repo, you must provide certain data into the secrets values (GitLab) and prepare the cluster to work with the CI flow.
1. Make sure you have infrastructure resources described above;
2. [Create YC CLI profile on your workstation, connect to the cloud resources](https://cloud.yandex.com/en/docs/cli/quickstart#install)
3. Obtain Managed K8S cluster id:
```
yc managed-kubernetes cluster list
```
4. [Create Kubernetes configuration on your local workstation to work with k8s entities and GitLab](https://cloud.yandex.com/en/docs/tutorials/infrastructure-management/gitlab-containers#k8s-get-token). Create GitLab variable (Settings>CI/CD>Variables) "KUBE_TOKEN" and specify the token value there;
5. Obtain ip of the k8s cluster. Create GitLab variable (Settings>CI/CD>Variables) "KUBE_URL" and specify the address there:
```
yc managed-kubernetes cluster get <Kubernetes cluster ID or name> --format=json \
  | jq -r .master.endpoints.external_v4_endpoint
```
6. Add view permissions for service account - required for prometheus:
```
kubectl --namespace default apply -f - <<< "
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: default-view
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: view
subjects:
  - kind: ServiceAccount
    name: default
    namespace: default"
```

## Obtain and add required variables to GitLab
Variables must be placed in Settings>CI/CD>Variables

- Kubernetes-cluster related:
```
KUBE_URL >> provided before, ip of the k8s cluster
KUBE_TOKEN >> token of the service account for GitLab runner
ALB_INGRESS_EXTERNAL_IP >> Extenal IP address (associated with ALB and your domain)
CERT_ID >> id of the security certificate
FQDN >> your domain name in form "test-domain.com"
K8S_SUBNET_ID >> subnet id in the cluster, linked to cluster node, list can be obtained by running "yc vpc subnet list"
```

- GitLab related
```
GITLAB_REGISTRY_K8S_KEY >> credentials in form of token in base64, need only the value dockerconfigjson itself
```
How to obtain it:
1. Login into your repository registry (where the code being stored). Example (change the name!):
```
docker --config /tmp login registry.gitlab.com
cat /tmp/config.json
```
2. Remove part of the created config - /tmp/config.json - credsStore line and an extra comma after “auths” block.
3. Run the login command again and you will get a config.json file with base64-encoded credentials ready to be used by external systems

- Nexus related:
```
HELM_REPO >> address of the repo in form https://nexus-something-dot-smht/repository/name-of-your-repo-for-the-project/
HELM_REPO_PASS >> repo user password
HELM_REPO_USER >> repo user name
```
1 repo required in Nexus system

- SonarQube related:
```
SONARQUBE_URL >> address of the service, in form of https://sonarqube.smth-smth.smth
SONAR_BACK_PRJ_KEY >> name of the project for backend (Go language)
SONAR_FRONT_PRJ_KEY >> name of the project for backend (JS language)
SONAR_LOGIN_BACK >> personal access token value for backend project, specified once you create a project
SONAR_LOGIN_FRONT >> personal access token value for frontend project, specified once you create a project
```
## How pipeline works
1. Backend and Frontend modules start automatically after commit with changes in the corresponding folders. Upon finishing the respective child pipeline you have:
    - Docker image of the corresponding element uploaded in the GitLab registry;
    - HELM deployent chart package for the corresponding element uploaded in the NEXUS repo
2. Uploading HELM charts (packages) to the k8s cluster can be initiated manually by user from the same pipeline:
    - Backend chart (latest version from Nexus repo);
    - Frontend chart (latest version from Nexus repo)
    - Monitoring charts - Prometheus, Grafana and Grafana Loki + Promtail (default credentials for Grafana is admin/admin)

## Checking resources status locally
- Status of HELM packages can be checked locally from the workstation by using:
```
helm list
```
- Kubernetes resources (full description of available commands can be found [here](https://kubernetes.io/docs/reference/kubectl/cheatsheet/) ):
```
kubectl get all
``` 
# k8s-flow

A demo repository to play with Kubernetes and indent some more yaml

## Goal

1. **Premise**
 1.1 Some considerations about Architecture

2. **CI/CD**
  2.1 Build and deploy image locally :boom:
  2.2 Create a pipeline to build image :boom:
  2.3 Add security scan to image :boom:
  2.4 Optimize image size and security :boom:
  2.5 Push image on a container registry :boom:

3. **Deploy**
  3.1 Write Kubernetes manifest to deploy Deployment :star:
  3.2 Add networking layer to our manifest :star:
  3.4 Expose applications using Nginx Ingress Controller :star:
  3.3 Replace our manifest with helm chart :star:
  3.4 GitOps flow using ArgoCD :boom:

4. **Scripting**
  4.1 Create deploy-script :star:
  4.2 Write a golang application to export nginx logs at path /logs :star:

5. **Logging & Monitoring**
 5.1 Prometheus/Grafana/ELK? TBD :boom:
___
  
## 1. Premise


In a real scenario I would have created an architecture in the following way:

**Infrastructure Provisioning** 
(I assume I am working on a cloud provider)
I would have used Terraform as an infrastructure as code tool and would have deployed the following resources:
- 3 Tier Network (Public, Private and DB subnet) 
- Kubernetes Cluster
- Image Repository for every container image
- Possible pipelines according to the cloud provider
- Management users
- Application Load Balancer to expose Kubernetes application (Nginx Ingress Controller
- Possible DBs if the choice fell on DBs managed by the cloud provider


Here you can see an example of how I manage terraform repositories -> [Pesonal Website Iac ](https://github.com/ettoreciarcia/personal-website-iac)  
And here -> [an article](https://ettoreciarcia.com/posts/01-iac-and-pipeline-my-personal-website/) on how I manage pipelines for IaC and an example of how I manage Terraform state

___

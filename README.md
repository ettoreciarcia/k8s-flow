# k8s-flow

A demo repository to play with Kubernetes and indent some more yaml

## Goal

### 1.  **Premise**
- [x] Some considerations about Architecture :boom:

### 2. **CI/CD**
- [x] Keep your code base clean :boom:
- [x] Build and deploy image locally :boom:
- [x] Create a pipeline to build image :boom:
- [x] Add security scan to image in pipeline :boom:
- [x] Optimize image size and security :boom:
- [x] Push image on a container registry :boom:

### 3. **Deploy**
- [ ] Write Kubernetes manifest to deploy Deployment :star:
- [ ] Add networking layer to our manifest :star:
- [ ] Expose applications using Nginx Ingress Controller :star:
- [ ] Replace our manifest with Helm chart :star:
- [ ] GitOps flow using ArgoCD :boom:

### 4.**Scripting**
- [ ] Create deploy-script :star:
- [ ] Write a golang application to export nginx logs at path /logs :star:

### 5. **Logging & Monitoring**
- [ ] Prometheus/Grafana/ELK? TBD :boom:
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
## 2. CI/CD

In this phase I'm going to create a CI/CD pipeline that will build our Docker containers and push the related images to a Registry container. There will also be image vulnerability checks in between

### 2.0 Keep your code base clean

If we want our code base to be clean and maintainable, we need to prevent badly formatted code from reaching it. We add a pre-commit hook to check yaml farmatting and other common errors.

```
repos:
-   repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v2.3.0
    hooks:
    -   id: check-yaml
    -   id: end-of-file-fixer
    -   id: trailing-whitespace
-   repo: https://github.com/psf/black
    rev: 22.10.0
    hooks:
    -   id: black
```

### 2.1 Build and deploy image locally

Let's start by building our images locally to figure out what we're dealing with

```docker build -t hecha00/app1:1.0 .```

Just a base Nginx container with a custom ```/location``` at path ```ping```

We can run it and curl on localhost

```docker run --env-file production.env -p 80:80 hecha00/app1:1.0 ```

Same thing for the second container.

The image is unnecessarily large and the Nginx version too old, it will surely be subject to CVE, we can see this with a sample ```docker scan [image_name]```

### 2.2 Create a pipeline to build image

We'll use GitHub Actions to get straight to the point.

For each push that modifies the files in samples/app1 or samples/app2 we will create a new image

I won't go into the details of how I connected the GitHub Actions to my Docker Hub account, for more info you can read my [article](https://ettoreciarcia.com/posts/01-iac-and-pipeline-my-personal-website/#32-authenticate-github-actions-against-aws). Did I mention I wrote an article? :)


In this case I'm treating the project as a monorepo, so I don't want all the pipelines to start with every push on the main branch.
We discriminate pipeline launch only to changes in the folders involved in the container build
We can do it in the following way:

```yaml
name: build app1
on:
  push:
    branches:
      - main
    paths:
      - samples/app1/**
  pull_request:
```

### 2.3 Add security scan image step in pipeline

For this purpose we will use a plugin present in the GitHub Marketplace

```yaml
      - name: Build an image from Dockerfile
        run: |
          docker build -t docker.io/hecha00/app1:${{ github.sha }} samples/app1
      - name: Scan Docker Image
        uses: aquasecurity/trivy-action@0.8.0
        with:
          image-ref: 'docker.io/hecha00/app1:${{ github.sha }}'
          format: 'table'
          exit-code: '1'
          ignore-unfixed: true
          vuln-type: 'os,library'
          severity: 'CRITICAL,HIGH'
```

We have configured our pipeline in such a way that there is an exit code of 1 in case the vulnerabilities found are HIGH or CRITICAL.

Let's go for a run!

As we suspected, our pipeline failed because the images from which the Dockerfiles start are old and vulnerable

![security-check](img/security-check.png)

Image scanning in our pipeline found 37 vulnerabilitiesof **(HIGH: 25, CRITICAL: 12)**
In reality there are many other vulnerabilities but less impactful from a security point of view.

### 2.4 Optimize image size and security

These images aren't just devilish, they're way too big! 134MB for a container with plain nginx?

Let's try using a newer and lighter nginx image, **1-alpine-slim**.
(It Weighs only 4.77MB!)


But why should we do it?
In the end, one image is as good as another, as long as it works! No?
The size of the Docker images, in certain scenarios, avalanches on all components and in the long run the price you pay is high.
Here are some consequences of having unnecessarily heavy container images
- They take longer to build, this translates into pipelines that take longer to finish and therefore a worse **developer experience**
- Image repositories have costs related to the space we occupy. Larger images -> **Higher storage costs**
- They take longer to launch, in a highly scalable ecosystem where pods are destroyed and created all the time, **this makes the Kubernetes cluster less responsive**
- They take up more memory space once they are running, **this reducing cluster resources that we could dedicate to something else**
- The larger the image, the greater the attack surface for attackers.
This results in **security issues**


Want to learn more about how to reduce the size of container images? Read my personal wiki [Minimal Docker images](https://github.com/ettoreciarcia/homelab-learning/tree/main/Kubernetes/Paper/smaller-image-docker)

### 2.5 Push image on a container registry

Here we go, our pipeline built the Docker images and then pushed the images to our repository after verifying that they have no known vulnerabilities

![pipeline](img/pipeline.png)

In this case we have created multi-architecture images, they can run on both amd and arm architectures.

We can take a look at our beautiful image on Docker Hub

![docker-hub](img/docker-hub.png)

Everything is ready, we can finally enjoy Kubernetes

# k8s-flow

A demo repository to play with Kubernetes and indent some more yaml

## Goal

### 1.  **Premise**
- [x] Some considerations about Architecture :boom: 

### 2. **CI/CD**
- [x] Build and deploy image locally :boom:
- [x] Create a pipeline to build image :boom:
- [x] Add security scan to image :boom:
- [x] Optimize image size and security :boom:
- [x] Push image on a container registry :boom:

### 3. **Deploy**
- [ ] Write Kubernetes manifest to deploy Deployment :star:
- [ ] Add networking layer to our manifest :star:
- [ ] Expose applications using Nginx Ingress Controller :star:
- [ ] Replace our manifest with helm chart :star:
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

These images aren't just devilish, they're way too big! 134MB per un container con un semplice nginx?

Let's try using a newer and lighter nginx image, **nginx:1.23.2-alpine**. 
(It Weighs only 9.28MB!)

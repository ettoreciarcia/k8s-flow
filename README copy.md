# DevOps hiring test

## Description
The goal is to create a software release process with a similar structure as WeRoad. The automation should deploy the applications in the `samples` folder.
As you may know, WeRoad is an international company, so we have the need to deploy the same application for multiple countries.
We use Helm to help us to manage our Kubernetes applications. We use the same chart for almost all our projects.

## Goals
Create a script and related Helm chart to build the application container image and deploy it to Kubernetes.

#### The deploy process should be flexible enough to allow:
- Deploy of multiple projects (all the applications in the `samples` folder) and multiple countries (it, es, uk)
- Deploy to several environments (i.e. staging, production)
- Each application should have the `COUNTRY` environment variable with the proper value in each Kubernetes deployment
- Each application should have all (based on the cluster you are deploying to) the environment variables contained in the `deploy` folder (i.e. staging will use only `deploy/staging.env`)
- Expose each service using NGINX Ingress Controller with the domain `.info` (i.e. `app1-it.info`)

Your script will get the following inputs:
- Environment (staging or production)
- Project name
- Countries (comma separated)

For example `./deploy-script staging app1 it,es,uk` will create 3 deployments (one for each country: `app1-it`, `app1-es`, `app1-uk`) in the staging Kubernetes cluster.

#### Nice to have:
- Create a Golang app to fetch Kubernetes logs of the applications deployed in the previous step and display them at the path `/logs`.

## Notes
- You should provide a new git repository with the documentation on how to use the script and if additional setup is needed.
- Staging cluster will have only 1 node, while production cluster will have at least 2 nodes.
- We recommend to use [minikube](https://github.com/kubernetes/minikube) to run Kubernetes locally, but feel free to use what suits you best.
- Keep always in mind security, monitoring and availability as you were actually deploying to a production cluster.
- If something has not been mentioned, it means it's up to you.
- Feel free to add to the project whatever you want!

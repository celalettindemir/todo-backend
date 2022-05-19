# backend

## Go Test

- go test -coverprofile=cover.txt ./...
- go tool cover -html=cover.txt -o cover.html

## Run

- go run .

## Build and Run

- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

## Dockerize

- docker build -t my-golang-app .
- docker run -p 3000:3000 my-golang-app

## Gcloud Registry

- gcloud builds submit --tag "gcr.io/galvanic-sphinx-341912/celal258-modanisa-back:v0.1.6"

- kubectl delete deploy backend
- kubectl delete svc backend-service

## Environment

I have two container registry

- BROKER_TOKEN -> Bearer "Secret" , pact broker token
- BROKER_URL -> <https://modanisa-test1.pactflow.io>
- CI_REGISTRY -> <https://index.docker.io/v1/>
- CI_REGISTRY_IMAGE -> celal258/todo-back
- CI_REGISTRY_GKE_IMAGE -> todo-back , Google Cloud Platform CI_REGISTRY -> gcr.io
- CI_REGISTRY_USER -> celal258
- CI_REGISTRY_PASSWORD -> "Secret"
- GCP_PROJECT_ID -> galvanic-sphinx-341912 , Google Cloud Platform project Id
- GCP_SERVICE_ACCOUNT , Google Cloud Platform Account Information

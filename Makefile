build:
	docker build -t gcr.io/personal-200804/$(app)-$(env):v$(v) $(app)
run:
	docker run -t -i --rm -p :$(port):$(port) -e VERSION=$(v) -e APP_ENV=dev gcr.io/personal-200804/$(app)-$(env):v$(v) --port=$(port)
start:
	make build app=$(app) v=1 env=dev
	make run app=$(app) v=1 env=dev port=$(port)
minikube:
	make build v=$(v) env=dev
	make push v=$(v) env=dev
	kubectl config set current-context minikube
	kubectl delete secret gcr
	kubectl create secret docker-registry gcr \
        --docker-server=https://gcr.io \
        --docker-username=oauth2accesstoken \
        --docker-password=$(shell gcloud auth print-access-token) \
        --docker-email=peppysisay@gmail.com
	kubectl apply -f local-kubernetes.yaml
push:
	gcloud docker -- push gcr.io/personal-200804/web-$(env):v$(v)
	gcloud docker -- push gcr.io/personal-200804/api-$(env):v$(v)
deploy:
	make build v=$(v) env=prod
	make push v=$(v) env=prod
	gcloud config set project personal-200804
	gcloud container clusters get-credentials cluster-1
	kubectl apply -f production-kubernetes.yaml
	kubectl rollout status deployments personal-site-deployment
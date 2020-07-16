REGISTRY=techotron/go-web-app
APP=be-server:$(shell git rev-parse --short HEAD)
LOCAL_GO=go

build: 
	docker build -t $(REGISTRY)/$(APP) .
run:
	docker run -d -p 8080:8080 $(REGISTRY)/$(APP)
test:
	$(LOCAL_GO) test

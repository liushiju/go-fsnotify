.PHONY: run build image push clean

tag = v0.1
releaseName = go-fsnotify
dockerhubUser = shijuliu

ALL: run

run: build
	./go-fsnotify --configfile ./config.yaml &

build:
	go build -o $(releaseName) ./cmd/

image:
	docker build -t $(dockerhubUser)/$(releaseName):$(tag) .

push: image
	docker push $(dockerhubUser)/$(releaseName):$(tag)

clean:
	-rm -f ./$(releaseName)
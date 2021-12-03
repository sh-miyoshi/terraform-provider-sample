HOSTNAME=github.com
NAMESPACE=sh-miyoshi
NAME=sample
VERSION=0.0.1
OS_ARCH=linux_amd64

build:
	go build -o terraform-provider-${NAME}
install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv terraform-provider-${NAME} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

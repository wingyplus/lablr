deps:
	@go get -v -t ./...
deps-test: deps
	@go get -v github.com/onsi/ginkgo
	@go get -v github.com/onsi/gomega
	@go install -v github.com/onsi/ginkgo/ginkgo
test: deps deps-test
	@ginkgo -r --randomizeAllSpecs --failOnPending
	@tests/iss_2.sh

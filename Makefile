
cover:
	go test ./... -coverprofile cover.out
	go tool cover -func cover.out | grep total | awk '{print $3}'
	rm cover.out
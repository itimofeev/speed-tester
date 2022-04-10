
coverage:
	go test ./... -coverprofile cover.out > /dev/null
	go tool cover -func cover.out | grep total | awk '{print $3}'
	rm cover.out
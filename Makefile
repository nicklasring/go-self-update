build:
	go build -o updates/go-self-update -ldflags "-X main.Version=2"
	go build -o go-self-update -ldflags "-X main.Version=1"

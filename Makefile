ifndef GOPATH
	GOPATH=~/go
endif

RUNNING_PDIS=$(ps aux | awk '/localhost:3001/ { print $2}')
ifndef RUNNING_PDIS
	KILL_COMMAND=$(shell kill $(RUNNING_PDIS))
endif

dev: kill-bk-process server-dev ui-dev

kill-bk-process:
	@echo $(KILL_COMMAND)
	
server-dev:
	#@go run cmd/main.go --listen=localhost:3001 --build=ui/build &
	@watcher  -run cmd/main.go -watch /github.com/daxsorbito/golang-reactjs/api --listen=localhost:3001 --build=ui/build &

ui-dev:	
	@cd ui && REACT_APP_PUBLIC_URL=http://localhost:3001 yarn start

prod:
	@cd ui && PUBLIC_URL=http://localhost:3000/ui/build yarn build && cd .. && go run cmd/main.go --listen=localhost:3000 --build=ui/build
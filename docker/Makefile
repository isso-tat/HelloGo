# File specifically for docker usage.
# For details of each task, see ../Makefile
PID_FILE = /tmp/my-app.pid
GO_FILES = $(wildcard *.go)

start:
	go run ../$(GO_FILES) & echo $$! > $(PID_FILE)

stop:
	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`

before:
	@echo "STOPED my-app" && printf '%*s\n' "40" '' | tr ' ' -

restart: stop before start
	@echo "STARTED my-app" && printf '%*s\n' "40" '' | tr ' ' -

serve: start
	fswatch -or --event=Updated /go/src/github.com/tatsuki5820iso/HelloGo | \
	xargs -n1 -I {} make restart

.PHONY: start before stop restart serve

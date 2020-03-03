
PROGRAM := pg_tileserv
CONTAINER := pramsey/$(PROGRAM)

RM = /bin/rm
CP = /bin/cp
MKDIR = /bin/mkdir

.PHONY: all check clean test docker docs install uninstall

GOFILES := $(wildcard *.go)

all: $(PROGRAM)

check:
	@go version

clean:
	$(info Cleaning project...)
	@rm -f $(PROGRAM)
	@rm -rf docs/*

docs:
	@rm -rf docs/* && cd hugo && hugo && cd ..

$(PROGRAM): $(GOFILES)
	go build -v

docker: $(PROGRAM) Dockerfile.ci
	docker build -f Dockerfile.ci --build-arg VERSION=`./$(PROGRAM) --version | cut -f2 -d' '` -t $(CONTAINER):latest .
	docker image prune --force

test:
	go test -v

install: $(PROGRAM) docs
	$(MKDIR) -p $(DESTDIR)/usr/bin
	$(MKDIR) -p $(DESTDIR)/usr/share/$(PROGRAM)
	$(MKDIR) -p $(DESTDIR)/etc
	$(CP) $(PROGRAM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(CP) example.toml $(DESTDIR)/etc/$(PROGRAM).toml
	$(CP) -r assets $(DESTDIR)/usr/share/$(PROGRAM)/assets
	$(CP) -r docs $(DESTDIR)/usr/share/$(PROGRAM)/docs

uninstall:
	$(RM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(RM) $(DESTDIR)/etc/$(PROGRAM).toml
	$(RM) -r $(DESTDIR)/usr/share/$(PROGRAM)


BINS := bin/gosendperf

.PHONY: all clean build

all: build

build: $(BINS)

build-%: bin/%
	@true

run-%: cmd/%/main.go
	go run $<

bin/%: cmd/%/main.go
	go build -o $@ $<

clean:
	rm $(BINS)
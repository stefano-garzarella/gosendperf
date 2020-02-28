BINS := bin/gosendperf

.PHONY: all clean build

all: build

build: $(BINS)

build-%: bin/%
	@true

run-%: cmd/%/*.go
	go run $^

bin/%: cmd/%/*.go
	go build -o $@ $^

clean:
	rm $(BINS)

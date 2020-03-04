BINS := bin/gosendperf

.PHONY: all clean build

all: build

build: $(BINS)

build-%: bin/%
	@true

run-%: cmd/% cmd/%/*.go
	cd $<; go run .

bin/%: cmd/% cmd/%/*.go
	cd $<; go build -o ../../$@ .

clean:
	rm $(BINS)

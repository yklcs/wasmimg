SRCDIR := mozjpeg-4.1.1
SRCURL := https://github.com/mozilla/mozjpeg/archive/refs/tags/v4.1.1.tar.gz
SRCTARGZ := mozjpeg.tar.gz

BIN := mozjpeg.wasm
BINWAT := mozjpeg.wat

EMCMAKE := emcmake cmake
EMMAKE := emmake $(MAKE)
EMCC := emcc
EMCXX := em++

EMCMAKEFLAGS := \
	-DENABLE_SHARED=false \
	-DPNG_SUPPORTED=false \
	-DWITH_TURBOJPEG=false

.EXPORT_ALL_VARIABLES:

CXXFLAGS := -O3 -msimd128 -fno-exceptions
CFLAGS := $(CXXFLAGS)
LDFLAGS := -O3

all: mozjpeg.wasm

$(SRCDIR):
	curl -sL -o $(SRCTARGZ) https://github.com/mozilla/mozjpeg/archive/refs/tags/v4.1.1.tar.gz
	tar xzf $(SRCTARGZ)
	rm $(SRCTARGZ)
	test -d $@

$(SRCDIR)/libjpeg.a: $(SRCDIR)
	cd $^; \
	$(EMCMAKE) $(EMCMAKEFLAGS)
	$(EMMAKE) -C $(SRCDIR) -j8

$(BIN): export CC = $(EMCC)
$(BIN): export CXX = $(EMCXX)

$(BIN): $(SRCDIR)/libjpeg.a main.cpp
	$(CXX) $(CXXFLAGS) \
		--no-entry \
		-sFILESYSTEM=0 \
		-sTOTAL_MEMORY=64MB \
		-sEXPORT_NAME="mozjpeg" \
		-o $@ \
		-Wall \
		-std=c++11 \
		$^

$(BINWAT): $(BIN)
	wasm2wat $^ > $@

.PHONY: clean
clean:
	-rm $(BIN) $(BINWAT)
	-rm -rf $(SRCDIR)

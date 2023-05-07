CODECSDIR := codecs
CODECS := \
	mozjpeg

all:
	$(MAKE) -C $(foreach C,$(CODECS),$(CODECSDIR)/$C)

clean:
	$(MAKE) -C $(foreach C,$(CODECS),$(CODECSDIR)/$C) clean

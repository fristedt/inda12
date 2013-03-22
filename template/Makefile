-include latex_env

LATEX?=pdflatex # assign pdflatex if there's no other latex in conf/env
LATEXFLAGS?=-file-line-error

# Lägg till en tmp-katalog som byggkatalog
TMPDIR:=$(shell mktemp -d)
LATEXFLAGS+=-output-directory $(TMPDIR)

LATEX_CMD=$(LATEX) $(LATEXFLAGS)

.PHONY: default
default: template.pdf

# make pdfs, copy it to say ht1.tex, then `make ht1.pdf`
%.pdf: %.tex labbcover.tex
	$(LATEX_CMD) $<
	$(LATEX_CMD) $<
	mv $(TMPDIR)/$@ $@

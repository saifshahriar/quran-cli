VERSION = v0.2-r$(shell date +%Y%m%d)
PREFIX = /usr/local
MANPREFIX = $(PREFIX)/share/man
SRC = quran-cli.go

quran-cli:
	go build -buildvcs=false -ldflags "-X 'main.VERSION=$(VERSION)'"

install: quran-cli
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f quran-cli $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/quran-cli
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	sed "s/VERSION/$(VERSION)/g" < quran-cli.1 > $(DESTDIR)$(MANPREFIX)/man1/quran-cli.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/quran-cli.1

clean:
	rm -f quran-cli

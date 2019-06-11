#!/bin/sh

# some function definition

# markdown2html file.md
markdown2html() {
	file=$1

	copy=$(mktemp)
	sed '1d' "$file" > "$copy"
	title=$(head -n 1 "$file" | awk -F: '{ print $2 }')
	
	tmpfile=$(mktemp)
	
	# special rules for wcag report
	if [ "$file" = "src/wcag.md" ]; then
		sed \
			-e 's,^livello: \(.*\)$,<span>\1</span>,' \
			-e 's,^source: \(.*\)$,<span>\1</span>,' \
			-i'' \
			"$copy"
	fi
	
	$md -o "$tmpfile" "$copy"

	out=$(echo $file | sed -e 's/md$/html/' -e 's/^src/build/')

	# ED IS THE STANDARD TEXT EDITOR!
	# Git bash has vim but not ED!
	cp template.html "$out"
	$ed "$out" <<EOF 2>&1 >/dev/null
%s#TITLE#$title#
/CONTENT
d
-1
.r $tmpfile
w
q
EOF

	rm "$copy" "$tmpfile"
}

# setup

# ED IS THE STANDARD TEXT EDITOR
if which ed 2>&1 >/dev/null; then
	ed=ed
elif which vi 2>&1 >/dev/null; then
	ed="vi -e"
elif which vim 2>&1 >/dev/null; then
	ed="vim -e"
else
	echo "ed, vi or vim not found. One of them is needed to generate the HTML files."
	exit 1
fi

# markdown input output
if which lowdown 2>&1 >/dev/null; then
	md=lowdown
elif which pandoc 2>&1 >/dev/null; then
	md=pandoc
else
	echo "lowdown or pandoc are needed to tranlate markdown to HTML. Install one of them"
	exit 1
fi

# main:

action=${1:-build}

if [ $# -ne 0 ]; then
	shift
fi

case "$action" in
clean)
	rm -rf build
	touch build/.gitkeep
	rm -f progetto.tar.gz
	;;

build)
	rm -rf build/*
	cp -r css build/
	cp -r img build/
	cp -r js build/
	for i in src/*.md; do
		echo "compiling $i..."
		markdown2html "$i"
	done
	;;

archive)
	./build.sh build
	tar czvf progetto.tar.gz build
	;;

serve)
	cd build

	# try to use every version of python known to mankind
	port=${1:-8000}
	echo serving on port $port
	{ python -m SimpleHTTPServer $port ||
		python2 -m SimpleHTTPServer $port ||
		python -m http.server $port ||
		python3 -m http.server $port; } 2>&1 >/dev/null
	;;

fmt)
	if [ -n "$1" ]; then
		prettier --write "$1"
		exit 0
	fi

	prettier --write template.html
	prettier --write style.css
	for i in build/*.html; do
		prettier --write "$i"
	done
	;;

*)
	echo "USAGE: $0 [action] [args...]"
	echo "where action can be one of:"
	echo "clean: clean the previous build"
	echo "build: builds the site (no args)"
	echo "archive: create progetto.tar.gz from build"
	echo "serve: the first option is the port, by default 8000"
	echo "fmt: by default re-format all html and css files (even in build/), but accepts the name of the file to re-format (in-place)"
	exit 1
	;;
esac

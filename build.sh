#!/bin/sh

# some function definition

# ED IS THE STANDARD TEXT EDITOR
edit() {
	if which ed 2>&1 >/dev/null; then
		ed "$@"
		return
	fi
	vim -e "$@"
}

# markdown input output
markdown() {
	if which lowdown 2>&1 >/dev/null; then
		lowdown -o "$2" "$1"
		return
	fi
	pandoc -o "$2" "$1"
}

# markdown2html file.md
markdown2html() {
	file=$1

	tmpfile=$(mktemp)

	markdown "$file" "$tmpfile"

	out=$(echo $file | sed -e 's/md$/html/' -e 's/^src/build/')

	# ED IS THE STANDARD TEXT EDITOR!
	# Git bash has vim but not ED!
	cp template.html "$out"
	edit "$out" <<EOF 2>&1 >/dev/null
/CONTENT
d
-1
.r $tmpfile
w
q
EOF

	rm "$tmpfile"
}

# main:

action=${1:-build}

if [ $# -ne 0 ]; then
	shift
fi

case "$action" in
build)
	rm -rf build/*
	cp style.css build
	for i in src/*.md; do
		echo "compiling $i..."
		markdown2html "$i"
	done
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
	echo "build: builds the site (no args)"
	echo "serve: the first option is the port, by default 8000"
	echo "fmt: by default re-format all html and css files (even in build/), but accepts the name of the file to re-format (in-place)"
	exit 1
	;;
esac

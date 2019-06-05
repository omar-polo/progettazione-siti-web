#!/bin/sh

# some function definition

# markdown2html file.md
markdown2html() {
	file=$1
	
	tmpfile=$(mktemp)
	
	# TODO: use pandoc if lowdown is not available
	lowdown -o "$tmpfile" "$file"
	
	out=$(echo $file | sed -e 's/md$/html/' -e 's/^src/build/')
	
	# ED IS THE STANDARD TEXT EDITOR!
	cp template.html "$out"
	ed "$out" <<EOF 2>&1 >/dev/null
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
		{ python -m SimpleHTTPServer $port || \
			python2 -m SimpleHTTPServer $port || \
			python -m http.server $port || \
			python3 -m http.server $port; } 2>&1 >/dev/null
		;;
		
	*)
		echo USAGE: $0 [action] [args...]
		echo where action can be one of:
		echo "build: builds the site (no args)"
		echo serve: the first option is the port, by default 8000
		exit 1
		;;
esac
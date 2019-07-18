/**
 * Takes a json representation of a `Section' and flatten it into a list
 * of sections. Since we're there, attach also to the SC a couple of extra
 * info.
 */
flattenData = data => {
    const r = []

    if (data == null) {
        return r
    }

    for (p of data) {
        if (p.subsections == null) {
            continue
        }

        for (gl of p.subsections) {
            for (sc of gl.subsections) {
                const { title, metadata } = sc
                const { livello, source, outcome } = metadata

                r.push({
                    title,
                    level: livello,
                    source,
                    outcome,
                    principle: p.title,
                    guideline: gl.title,
                })
            }
        }
    }

    return r
}

/**
 * Takes a functions and a interval and return a function. The returned
 * function is guaranteed to call the given callback *at most* every
 * `time' milliseconds.
 */
const debounce = (fn, time) => {
    let timeout;

    return function() {
      const functionCall = () => fn.apply(this, arguments);
    
      clearTimeout(timeout);
      timeout = setTimeout(functionCall, time);
    }
}

/**
 * generate and renders the list of results.
 *
 * @param query   string    the user query
 * @param data    section[] the list of sections to search
 * @param results node      the dom node to render the risult in
 */
const searchResults = (query, data, results) => {
    const res = []
    
    results.innerHTML = ''

    for (d of data) {
        if (d.title.toLowerCase().indexOf(query.toLowerCase()) != -1) {
            if (d.source == null)
                console.log("source for", d.title, "is missing")
                
            const li   = document.createElement("li")
            const a    = document.createElement("a")
            const span = document.createElement("span")
            
            span.innerText = `${d.principle} » ${d.guideline} — livello ${d.level}`

            a.setAttribute("href", "wcag.html#" + d.source)
            a.setAttribute("title", d.title)
            a.innerText = d.title

            a.addEventListener('click', (e) => {
                toggleSearchBar()
                e.stopPropagation()
            })
            li.appendChild(a)
            li.appendChild(span)
            li.addEventListener('click', function (e) {
                toggleSearchBar()
                console.log("this is", this)
                location.href = this.querySelector('a').href
            })
            results.appendChild(li)
        }
    }

    return res
}

/**
 * Toggle the search bar
 */
const toggleSearchBar = () => {
    document.querySelector(".searchbox").classList.toggle("open")
}

/**
 * Show/hide the arrow to scroll up based on how far the user scrolled
 * down.
 */
function showArrow(e) {
    return () => {
        const s = window.scrollY
        console.log("window.scrollY:", s, e)		
        if (s > 300) {
            e.classList.add("show")
        } else {
            e.classList.remove("show")	
        }
    }
}

document.addEventListener("DOMContentLoaded", function(){
    /* react on window scroll event */
	const arrow = document.getElementById("toTop")
    window.addEventListener('scroll', debounce(showArrow(arrow), 300))

    const rootMap = document.getElementById("root-map")
    if (rootMap != null) {
        rootMap.addEventListener("click", function(e){
            e.preventDefault();
            var t = e.target;
            t.innerHTML = "Ti Trovi qua!"
        });
    }

    /* handle the search */
    if (document.querySelector("li.search")) {
        const data = flattenData(window.data)

        document
            .querySelector("li.search")
            .addEventListener("click", toggleSearchBar)

        const search = document.querySelector("#search-input")
        const results = document.querySelector("#results ul")
        const searchfn = debounce(searchResults, 300)

        searchfn("", data, results)

        search.addEventListener("input", () => {
            searchfn(search.value, data, results)
        })
    }
});
/**
 * Takes a json representation of a `Section' and flatten it into a list
 * of sections.
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

        for (lg of p.subsections) {
            r.push(...lg.subsections)
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
            if (d.metadata.source == null)
                console.log("Metadata for", d.title, "is missing")
            
            const i = document.createElement("a")
            i.setAttribute("href", "wcag.html#" + d.metadata.source)
            i.setAttribute("title", d.title)
            i.innerText = d.title
            
            i.addEventListener('click', toggleSearchBar)
            results.appendChild(i)
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

document.addEventListener("DOMContentLoaded", function(){
    const data = flattenData(window.data.subsections)

    document
        .querySelector("li.search")
        .addEventListener("click", toggleSearchBar)

    const search = document.querySelector("#search-input")
    const results = document.querySelector("#results")
    const searchfn = debounce(searchResults, 300)

    searchfn("", data, results)

    search.addEventListener("input", () => {
        searchfn(search.value, data, results)
    })
});
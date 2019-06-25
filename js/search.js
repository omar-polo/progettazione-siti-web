flattenData = data => {
    const r = []

    if (data == null) {
        return r;
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

const searchResults = (query, data) => {
    const res = []

    for (d of data) {
        console.log("d", d)
        if (d.title.toLowerCase().indexOf(query.toLowerCase())) {
            res.push(d)
        }
    }

    return res
}

const toggleSearchBar = () => {
    document.querySelector(".searchbox").classList.toggle("open")
}

function debouce(fn, wait, immediate) {
    let timeout;
    return () => {
        const later = () => {
            timeout = null;
            if (!immediate) {
                fn(...arguments)
            }
        }
        const callNow = immediate && !timeout;
        clearTimeout(timeout)
        timeout = setTimeout(later, wait)
        if (callNow) {
            fn(...arguments)
        }
    }
}

document.addEventListener("DOMContentLoaded", function(){
    console.log("Hello DOM!");

    const data = flattenData(window.data.subsections);
    console.log(data)

    document.querySelector("li.search").addEventListener("click", toggleSearchBar)

    const search = document.querySelector("#search-input")
    const results = document.querySelector("#results")

    //const searchfn = debouce(searchResults, 300, false)
    const searchfn = searchResults

    searchResults(data, "")

    search.addEventListener("input", () => {
        //console.log(search.value)
            
        const res = searchfn(search.value, data)

        results.innerHTML = "";

        for (r of res) {
            const i = document.createElement("a");
            i.setAttribute("href", "wcag.html#" + r.metadata.source)
            i.setAttribute("title", r.title)
            i.innerText = r.title;

            i.addEventListener('click', e => {
                toggleSearchBar()
            })

            results.appendChild(i)
        }
    });
});
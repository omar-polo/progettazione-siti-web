document.addEventListener("DOMContentLoaded", function(){
    console.log("Hello DOM!");

    const search = document.querySelector("#search-input")
    const results = document.querySelector("#results")

    search.addEventListener("input", () => {
        //console.log(search.value)
            
        const res = [];

        for (d of data) {
            const text = search.value.toLowerCase()
            if (d.title.indexOf(text) != -1) {
                res.push(d)
            }
        }

        console.log(res)

        results.innerHTML = "";
        for (r of res) {
            const i = document.createElement("div");
            i.setAttribute("data-url", r.page + '#' + r.anchor)
            i.innerText = r.title;

            i.addEventListener('click', e => {
                location.href = e.target.getAttribute("data-url")
            })

            results.appendChild(i)
        }
    });
});
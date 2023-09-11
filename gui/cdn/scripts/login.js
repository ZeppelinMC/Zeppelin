function login() {
    const password = document.getElementById("input").value
    if (!password) return setError(true)
    const urlParams = new URLSearchParams(window.location.search);
    const to = urlParams.get('to') || "/"
    console.log(to)
}

function updateInput() {
    setError(!document.getElementById("input").value)
}

function setError(value) {
    const input = document.getElementById("input")
    input.style["border-color"] = value ? "red" : "black"
    if (value) {
        input.classList.add("red-placeholder");
        document.getElementById("error").style["display"] = "block"
    } else {
        input.classList.remove("red-placeholder");
        document.getElementById("error").style["display"] = "none"
    }
}
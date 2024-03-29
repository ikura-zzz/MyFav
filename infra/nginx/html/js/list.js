function getURL() {
    url = location.pathname + location.search
    return url
}

function allfav() {
    for (let i = 0; true; i++) {
        let elem = document.getElementById("fav" + i);
        if (!elem) {
            return;
        }

        let children = elem.children;
        elem.style = "";
        for (let h = 0; h < children.length; h++) {
            children.item(h).style = "";
        }
        history.replaceState('', '', getURL() + "#all")
    }
}
function now() {
    for (let i = 0; true; i++) {
        let timing = document.getElementById("timing" + i);
        if (!timing) {
            history.replaceState('', '', getURL() + "#now")
            return;
        }

        let elem = document.getElementById("fav" + i);
        let children = elem.children
        if (timing.value == "1") {
            elem.style.display = "none";
            for (let h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        } else if (timing.value == "2") {
            elem.style = "";
            for (let h = 0; h < children.length; h++) {
                children.item(h).style = "";
            }
        } else if (timing.value == "3") {
            elem.style.display = "none";
            for (let h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        }
    }
}

function already() {
    for (let i = 0; true; i++) {
        let timing = document.getElementById("timing" + i);
        if (!timing) {
            history.replaceState('', '', getURL() + "#already")
            return;
        }

        let elem = document.getElementById("fav" + i);
        let children = elem.children
        if (timing.value == "1") {
            elem.style = "";
            for (let h = 0; h < children.length; h++) {
                children.item(h).style = "";
            }
        } else if (timing.value == "2") {
            elem.style.display = "none";
            for (let h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        } else if (timing.value == "3") {
            elem.style.display = "none";
            for (let h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        }
    }
}

function wish() {
    for (let i = 0; true; i++) {
        let timing = document.getElementById("timing" + i);
        if (!timing) {
            history.replaceState('', '', getURL() + "#wish")
            return;
        }

        let elem = document.getElementById("fav" + i);
        let children = elem.children;
        elem.style = "";
        if (timing.value == "1") {
            elem.style.display = "none";
            for (let h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        } else if (timing.value == "2") {
            elem.style.display = "none";
            for (let h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        } else if (timing.value == "3") {
            elem.style = "";
            for (let h = 0; h < children.length; h++) {
                children.item(h).style = "";
            }
        }
    }
}

function firstload() {
    let hash = location.hash
    if (hash != "") {
        switch (hash) {
            case "#all":
                document.getElementById("all").checked = true
                allfav();
                break;
            case "#already":
                document.getElementById("already").checked = true
                already();
                break;
            case "#now":
                document.getElementById("now").checked = true
                now();
                break;
            case "#wish":
                document.getElementById("wish").checked = true
                wish();
                break;
        }
        return
    }
    let radioelem = [document.getElementById("all"), document.getElementById("already"), document.getElementById("now"), document.getElementById("wish")];
    let tabid = 0;
    for (let i = 0; i < radioelem.length; i++) {
        if (radioelem[i].checked) {
            tabid = i;
            break;
        }
    }
    switch (tabid) {
        case 0:
            allfav();
            break;
        case 1:
            already();
            break;
        case 2:
            now();
            break;
        case 3:
            wish();
            break;
    }
}

function copy() {
    let copyTag = document.querySelector("#copy");
    let uri = encodeURI(copyTag.value);
    navigator.clipboard.writeText(uri);
    alert("コピーしました。")
}
document.querySelector("#copy").addEventListener("click", copy);

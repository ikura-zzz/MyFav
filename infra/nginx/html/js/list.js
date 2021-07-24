function allfav() {
    for (var i = 0; true; i++) {
        var elem = document.getElementById("fav" + i);
        if (!elem) {
            return;
        }

        var children = elem.children;
        elem.style = "";
        for (var h = 0; h < children.length; h++) {
            children.item(h).style = "";
        }
    }
}

function now() {
    for (var i = 0; true; i++) {
        var timing = document.getElementById("timing" + i);
        if (!timing) {
            return;
        }

        var elem = document.getElementById("fav" + i);
        var children = elem.children
        if (timing.value == "1") {
            elem.style.display = "none";
            for (var h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        } else if (timing.value == "2") {
            elem.style = "";
            for (var h = 0; h < children.length; h++) {
                children.item(h).style = "";
            }
        } else if (timing.value == "3") {
            elem.style.display = "none";
            for (var h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        }
    }
}

function already() {
    for (var i = 0; true; i++) {
        var timing = document.getElementById("timing" + i);
        if (!timing) {
            return;
        }

        var elem = document.getElementById("fav" + i);
        var children = elem.children
        if (timing.value == "1") {
            elem.style = "";
            for (var h = 0; h < children.length; h++) {
                children.item(h).style = "";
            }
        } else if (timing.value == "2") {
            elem.style.display = "none";
            for (var h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        } else if (timing.value == "3") {
            elem.style.display = "none";
            for (var h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        }
    }
}

function wish() {
    for (var i = 0; true; i++) {
        var timing = document.getElementById("timing" + i);
        if (!timing) {
            return;
        }

        var elem = document.getElementById("fav" + i);
        var children = elem.children;
        elem.style = "";
        if (timing.value == "1") {
            elem.style.display = "none";
            for (var h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        } else if (timing.value == "2") {
            elem.style.display = "none";
            for (var h = 0; h < children.length; h++) {
                children.item(h).style.display = "none";
            }
        } else if (timing.value == "3") {
            elem.style = "";
            for (var h = 0; h < children.length; h++) {
                children.item(h).style = "";
            }
        }
    }
}
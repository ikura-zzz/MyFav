function validchkshokuzaidata(){
    const nameNode = document.getElementById("name");
    const quantityNode = document.getElementById("quantity");
    const foodsweightNode = document.getElementById("foodsweight");
    const buydateNode = document.getElementById("buydate");
    const uselimitNode = document.getElementById("uselimit");
    const errtag = document.getElementById("err");
    if (nameNode.value == ""){
        errtag.innerHTML = "食品の名前が空欄です"
        return false;
    }
    if (quantityNode.value == "" && foodsweightNode.value == ""){
        errtag.innerHTML = "個数とグラム数が両方空欄です"
        return false;
    }
    if (uselimitNode.value == ""){
        errtag.innerHTML = "賞味（消費）期限が空欄です"
        return false;
    }
    return true
}

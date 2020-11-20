;
var a = document.getElementsByClassName("ts input"),
    items = [];

for (let i = 0; i < a.length; i++) {
    items[a[i].classList[2]] = a[i];
    items[a[i].classList[2]].input = a[i].getElementsByTagName("input")[0];
}
console.log(items)
function error_render(error) {
    ts('.top.left.snackbar').snackbar({
        content: error,
        action: "CLOSE",
        actionEmphasis: 'info',
    });
}

document.getElementsByClassName("ts button")[0].addEventListener("click", function () {
    if (items["password"].input.value.localeCompare(items["passwordConfirmation"].input.value) != 0) {
        items["password"].className = "ts error input";
        items["passwordConfirmation"].className = "ts error input";
        error_render("Passwords do not match.");
        return
    }
    if (items["password"].input.value.localeCompare(items["passwordConfirmation"].input.value) != 1) {
        items["password"].className = "ts input";
        items["passwordConfirmation"].className = "ts input";
    }
    user_exist("/api/v1/existusername/" + items["name"].input.value)
        .then((data) => {
            if (data == true) {
                items["name"].className = "ts error input";
                error_render("Username is used!");
                return
            }
            if (data == false) {
                items["name"].className = "ts success input";
                return
            }
        });
});

async function user_exist(url) {
    try {
        const response = await fetch(url, {
            method: 'GET',
        });
        const json = await response.json();
        return json.existuser;
    } catch (error) {
        error_render(error);
    }
}

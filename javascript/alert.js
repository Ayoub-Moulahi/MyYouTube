let close_button = document.getElementById("btn_alert");
let alert_container = document.getElementById("alert-content");
function hide_alert(){
    alert_container.style.display="none";
}
close_button.addEventListener("click", hide_alert);
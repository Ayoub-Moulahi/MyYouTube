const articles = document.querySelectorAll("article");

for (let i=0;i<articles.length;i++){
    articles[i].addEventListener("click",make_bigger);
    articles[i].addEventListener("click",topFunction);
    document.addEventListener("click",function (e) {
        if (!articles[i].contains(e.target)){
            articles[i].style.width="400px";
            articles[i].style.height="300px";
            articles[i].style.order="0";
            articles[i].children[0].style.width="400px";
            articles[i].children[0].style.height="274px";


        }
    });

}
function make_bigger() {

    this.style.height = "650px";
    this.style.width = "900px";
    this.style.order = "-1";
    this.firstElementChild.style.height="620px"
    this.firstElementChild.style.width="888px"

}


function topFunction() {
    document.body.scrollTop=0;
    document.documentElement.scrollTop=0;
    window.scrollTo({
        top:0,
        behavior:"smooth"
    });
}
document.addEventListener("htmx:afterRequest", function (event) {
  if (event.target.id === "contact-form" && event.detail.xhr.status === 200) {
    document.getElementById("contact-form").reset();
  }
});

const initData = window.Telegram.WebApp?.initData;

alert(initData || "Init data is not found")

// always on request set the authorization token with htmx
htmx.on("htmx:configRequest", (e)=> {
  e.details.headers["Authorization"] = `tma ${initData}`
  console.log("the init data is set to be", e.detail.headers["Authorization"])
})
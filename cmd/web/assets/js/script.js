document.addEventListener("htmx:afterRequest", function (event) {
  if (event.target.id === "contact-form" && event.detail.xhr.status === 200) {
    document.getElementById("contact-form").reset();
  }
});

const initData = window.Telegram.WebApp?.initData;

alert(initData || "Init data is not found")

// always on request set the authorization token with htmx
htmx.on("htmx:configRequest", (e)=> {
  e.detail.headers["AUTH"] = `tma ${initData}`
  console.log("the init data is set to be", e.detail.headers["AUTH"])
})

// 2025/02/03 08:13:51 sign is missing
// 2025/02/03 08:13:51 GET /assets/js/script.js 200 27.581µs
// 2025/02/03 08:13:51 Logging the init data sent with the request:  
// 2025/02/03 08:13:51 sign is missing
// 2025/02/03 08:13:51 GET /assets/images/rank.svg 200 26.33µs
// 2025/02/03 08:13:51 Logging the init data sent with the request:  
// 2025/02/03 08:13:51 sign is missing

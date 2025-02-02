document.addEventListener("htmx:afterRequest", function (event) {
  if (event.target.id === "contact-form" && event.detail.xhr.status === 200) {
    document.getElementById("contact-form").reset();
  }
});
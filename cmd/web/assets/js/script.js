const removeSelf = (e) => {
  e.currentTarget.remove()
}

// Initialize Swiper
const swiper = new Swiper('.swiper-container', {
  slidesPerView: 2, // Show 2 testimonials at a time
  spaceBetween: 30, // Space between testimonials
  loop: true, // Infinite loop
  autoplay: {
    delay: 2000, // Auto-slide every 5 seconds
    disableOnInteraction: false, // Continue autoplay even when user interacts
  },
  pagination: {
    el: '.swiper-pagination', // Pagination dots
    clickable: true,
  },
  breakpoints: {
    // Responsive breakpoints
    320: {
      slidesPerView: 1, // Show 1 testimonial on small screens
    },
    768: {
      slidesPerView: 2, // Show 2 testimonials on larger screens
    },
  },
});

const themeToggle = document.getElementById("theme-toggle");
const moonIcon = document.getElementById("moon-icon").parentElement;
const sunIcon = document.getElementById("sun-icon").parentElement;
const htmlElement = document.documentElement;

// Check for saved user preference
if (localStorage.getItem("theme") === "dark") {
  htmlElement.classList.add("dark");
  themeToggle.checked = true;
  moonIcon.classList.add("hidden");
  sunIcon.classList.remove("hidden");
} else {
  htmlElement.classList.remove("dark");
  themeToggle.checked = false;
  moonIcon.classList.remove("hidden");
  sunIcon.classList.add("hidden");
}

// Toggle dark mode and icons
themeToggle.addEventListener("change", () => {
  if (themeToggle.checked) {
    htmlElement.classList.add("dark");
    localStorage.setItem("theme", "dark");
    moonIcon.classList.add("hidden");
    sunIcon.classList.remove("hidden");
  } else {
    htmlElement.classList.remove("dark");
    localStorage.setItem("theme", "light");
    moonIcon.classList.remove("hidden");
    sunIcon.classList.add("hidden");
  }
});

document.addEventListener("htmx:afterRequest", function (event) {
  if (event.target.id === "contact-form" && event.detail.xhr.status === 200) {
    document.getElementById("contact-form").reset();
  }
});
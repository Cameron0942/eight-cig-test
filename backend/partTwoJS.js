let currentIndex = 0;
const slidesContainer = document.querySelector(".slides");

export function changeSlide(indexOffset) {
  const slideWidth = slidesContainer.querySelector(".slide").clientWidth;
  currentIndex += indexOffset;
  if (currentIndex < 0) {
    currentIndex = slidesContainer.childElementCount - 1;
  } else if (currentIndex >= slidesContainer.childElementCount) {
    currentIndex = 0;
  }
  slidesContainer.style.transform = `translateX(-${
    currentIndex * slideWidth
  }px)`;
}

// Automatic slide change
setInterval(() => {
  changeSlide(1);
}, 5000);
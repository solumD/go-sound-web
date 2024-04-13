function openModal() {
  document.querySelector(".modal").classList.add("active");
}

const closeModal = () => {
  document.querySelector(".modal").classList.remove("active");
};

const array = [
  "little_prince.m4a",
];

let index = 1;

function play() {
  document.querySelector(".audio").play();
  document.querySelector(".play").classList.add("dn");
  document.querySelector(".pause").classList.remove("dn");
}

function pause() {
  document.querySelector(".audio").pause();
  document.querySelector(".play").classList.remove("dn");
  document.querySelector(".pause").classList.add("dn");
}

function played() {
  document.querySelector(".audio").play();
  document.querySelector(".played").classList.add("dn");
  document.querySelector(".paused").classList.remove("dn");
}

function paused() {
  document.querySelector(".audio").pause();
  document.querySelector(".played").classList.remove("dn");
  document.querySelector(".paused").classList.add("dn");
}

function prev() {
  document.querySelector(".audio").play();
  document.querySelector(".audio").src = array[index];
  play();
  document.querySelector(".audio").play();
  if (index === 0) {
    return (index = array.length - 1);
  }
  index = index - 1;
}

function next() {
  document.querySelector(".audio").src = array[index];
  play();
  document.querySelector(".audio").play();
  if (index === array.length - 1) {
    return (index = 0);
  }
  index = index + 1;
}

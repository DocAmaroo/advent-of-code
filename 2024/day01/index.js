import {readLines} from "../utils/index.js";

const lines = readLines().map(line => line.match(/\d+/g).map(Number));

// -----------------------------------------------------------------
// Part 1
function totalDistance() {
  const left = [];
  const right = [];

  for (const [l, r] of lines) {
    left.push(l);
    right.push(r);
  }

  left.sort();
  right.sort();

  return left.reduce((acc, l, index) => acc + Math.abs(l - right[index]), 0);
}

console.log('solution #1:', totalDistance());

// -----------------------------------------------------------------
// Part 2
function similarityScore() {
  const left = [];
  const count = new Map([]);

  for (const [l, r] of lines) {
    left.push(l);
    count.set(r, (count.get(r) ?? 0) + 1);
  }

  return left.reduce((acc, l) => acc + (l * count.get(l) || 0), 0)
}

console.log('solution #2:', similarityScore());

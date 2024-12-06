import {readLines} from "../utils/index.js";

const lines = readLines().map((line) => line.match(/\d+/g).map(Number));

function compare(a, b, order) {
  return a !== b && (order === 'asc' ? a < b : a > b) && Math.abs(a - b) <= 3
}

// -----------------------------------------------------------------
// Part 1

function isSafe(report) {
  if (report[0] === report[1]) return false;

  const order = report[0] < report[1] ? 'asc' : 'desc';

  for (let i = 0; i < report.length - 1; i++) {
    const cmp = compare(report[i], report[i + 1], order);
    if (!cmp) return false;
  }

  return true;
}

console.log('solution #1:', lines.filter(isSafe).length);

// -----------------------------------------------------------------
// Part 2

function isSafeDampened(report) {
  if (isSafe(report)) return true;

  for (let i = 0; i < report.length; i++) {
    if (isSafe(report.toSpliced(i, 1))) {
      return true;
    }
  }
  return false;
}

console.log('solution #2:', lines.filter(isSafeDampened).length);

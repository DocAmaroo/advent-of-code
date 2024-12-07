import {readLines} from "../utils/index.js";

const program = readLines('input').join('\n').replace('\n', '');

// -----------------------------------------------------------------
// Part 1

function solve() {
  const expr = /mul\(([0-9]{1,3}),([0-9]{1,3})\)/g;
  const groups = [...program.matchAll(expr)];

  return groups.reduce((res, g) => res + (Number(g[1]) * Number(g[2])), 0);
}

console.log('solution #1:', solve());

// -----------------------------------------------------------------
// Part 2

function solveWithInstructions() {
  const expr = /mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)/g
  const groups = [...program.matchAll(expr)];

  let exec = true;
  return groups.reduce((res, g) => {
    if (g[0] === 'do()') {
      exec = true;
    } else if (g[0] === "don't()") {
      exec = false;
    } else if (exec) {
      res += Number(g[1]) * Number(g[2])
    }

    return res;
  }, 0)
}

console.log('solution #2:', solveWithInstructions())



import {readLines} from "../utils/index.js";

const [pageOrder, lists] = readLines('input').reduce((acc, line) => {
  if (line === '') return acc;

  const pageOrderMatch = line.match(/(\d+)\|(\d+)/);

  if (pageOrderMatch) {
    acc[0].push(pageOrderMatch.slice(1).map(Number));
  } else {
    acc[1].push(line.split(',').map(Number));
  }

  return acc;
}, [[], []]);

const map = new Map();
for (const [a, b] of pageOrder) {
  if (!map.get(a)) map.set(a, { before: [], after: [b] });
  else map.get(a).after.push(b);

  if (!map.get(b)) map.set(b,  { before: [a], after: [] });
  else map.get(b).before.push(a);
}

function isValid(list) {
  for (let i = 0; i < list.length; i++) {
    const el = map.get(list[i]);

    for (let j = 0; j < list.length; j++) {
      if (j === i) continue;
      if (j < i && !el.before.includes(list[j])) return false;
      if (j > i && !el.after.includes(list[j])) return false;
    }
  }

  return true;
}

console.log('solution #1:', lists.filter(isValid).reduce((total, list) => total + list[(list.length - 1)/2], 0));

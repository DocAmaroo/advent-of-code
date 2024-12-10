import { readFileSync } from 'fs';

export function readFile(filename) {
  return readFileSync(filename, 'utf-8').trimEnd()
}
export function readLines(filename = 'input', separator = '\n') {
  return readFile(filename).split(separator);
}

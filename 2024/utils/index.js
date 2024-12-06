import { readFileSync } from 'fs';

export function readLines(filename = 'input', separator = '\n') {
  return readFileSync(filename, 'utf-8').trimEnd().split(separator);
}

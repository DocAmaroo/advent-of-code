type Rotation = { direction: 1 | -1; value: number };

async function readInput(filename: string): Promise<Rotation[]> {
  const content = await Bun.file(filename).text();
  return content
    .split("\n")
    .filter(Boolean)
    .map((v) => {
      let [, direction, value] = v.match(/(L|R)(\d+)/) as RegExpMatchArray;
      return { direction: direction === "L" ? -1 : 1, value: Number(value) };
    });
}

function part1(rotations: Rotation[]) {
  let password = 0;
  rotations.reduce((acc, curr) => {
    let next = (acc + curr.direction * curr.value) % 100;
    if (next < 0) next += 100;
    if (next === 0) password++;
    return next;
  }, 50);

  console.log("part1", password);
}

function part2(rotations: Rotation[]) {
  let password = 0;
  rotations.reduce((acc, curr) => {
    for (let i = 0; i < curr.value; i++) {
      acc += curr.direction;
      if (acc === -1) acc = 99;
      if (acc === 100) acc = 0;
      if (acc === 0) password++;
    }
    return acc;
  }, 50);

  console.log("part2", password);
}

async function main() {
  const input = await readInput("input");
  part1(input);
  part2(input);
}

main();

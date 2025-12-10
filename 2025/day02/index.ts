type Range = [number, number];

function sumInvalidId(ranges: Range[], regExp: RegExp) {
  return ranges.reduce((acc, range) => {
    for (let i = range[0]; i <= range[1]; i++) {
      if (`${i}`.match(regExp)) acc += i;
    }

    return acc;
  }, 0);
}

async function main() {
  const input = await Bun.file("input").text();
  const ranges = input
    .split(",")
    .map((v) => v.split("-").map(Number)) as Range[];

  console.table({
    part1: sumInvalidId(ranges, /^(\d+)\1$/),
    part2: sumInvalidId(ranges, /^(\d+)\1+$/),
  });
}

main();

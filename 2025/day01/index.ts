async function readInput(filename: string) {
  const file = Bun.file(filename);
  return file.text();
}

type ParsedInput = { direction: "L" | "R"; value: number };

function parseInput(input: string): ParsedInput[] {
  return input
    .split("\n")
    .filter(Boolean)
    .map((v) => {
      const match = v.match(/(L|R)(\d*)/);
      if (!match) {
        throw new Error(`Regex failed with ${v}`);
      }

      return { direction: match[1], value: Number(match[2]) } as ParsedInput;
    });
}

async function main() {
  const input = await readInput("input.test");
  const parsed = parseInput(input).splice(0, 20);
  console.log(parsed.length);

  const start = 50;
  let i = start;
  let dialCount = 0;
  parsed.forEach(({ direction, value }) => {
    if (direction === "L") {
      i -= value;
    } else {
      i += value;
    }

    if (i > 99) {
      i -= 100;
    }

    if (i < 0) {
      i += 100;
    }

    if (i === 0) {
      dialCount++;
    }

    console.log(`${direction}${value} => i=${i}`);
  });

  console.log(dialCount);
}

main();

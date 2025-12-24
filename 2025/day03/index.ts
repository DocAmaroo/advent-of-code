function solve(lines: number[][], batteryLength: number) {
  return lines
    .filter((line) => line.length >= batteryLength)
    .map((bank) => {
      let res = 0;
      for (let i = batteryLength - 1; i >= 0; i--) {
        let max = Math.max(...bank.slice(0, bank.length - i));
        let index = bank.findIndex((x) => x === max);
        bank = bank.slice(index + 1);
        res = res * 10 + max;
      }
      return res;
    })
    .reduce((a, b) => a + b);
}

async function main() {
  const input = await Bun.file("input").text();
  const banks = input.split("\n").map((line) => line.split("").map(Number));

  console.table({
    part1: solve(banks, 2),
    part2: solve(banks, 12),
  });
}

main();

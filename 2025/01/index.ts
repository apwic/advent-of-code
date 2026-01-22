import * as fs from "fs"

const input = fs.readFileSync("input.txt", "utf-8").trim()
const lines: string[] = input.split("\n")
let start = 50;
let ans = 0;

for (const line of lines) {
	const rotation = line[0];
	const rotationValueStr = line.slice(1);
	let rotationValue = parseInt(rotationValueStr, 10);
	// mod 100 for value that is more than 100
	rotationValue %= 100

	if (rotation === "L") {
		start = (start - rotationValue)
	}

	if (rotation === "R") {
		start = (start + rotationValue)
	}

	if (start < 0) {
		start += 100
	} else if (start >= 100) {
		start -= 100
	}

	if (start == 0) {
		ans += 1
	}
}

console.log(ans)

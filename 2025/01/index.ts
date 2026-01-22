import * as fs from "fs"

const input = fs.readFileSync("input.txt", "utf-8").trim()
const lines: string[] = input.split("\n")
let start: number = 50;
let ans: number = 0;

for (const line of lines) {
	const rotation: string | undefined = line[0];
	const rotationValueStr: string | undefined = line.slice(1);
	let rotationValue: number = parseInt(rotationValueStr, 10);
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

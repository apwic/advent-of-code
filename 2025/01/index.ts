import * as fs from "fs"

const input = fs.readFileSync("input.txt", "utf-8").trim()
const lines: string[] = input.split("\n")
let dial = 50;
let ans1 = 0;
let ans2 = 0;

for (const line of lines) {
	const rotation = line[0];
	const rotationValueStr = line.slice(1);
	let rotationValue = parseInt(rotationValueStr, 10);

	// add div value for ans2
	ans2 += Math.floor(rotationValue / 100)

	// mod 100 for value that is more than 100
	rotationValue %= 100
	const dialBefore = dial;

	if (rotation === "L") {
		dial = (dial - rotationValue)
	}

	if (rotation === "R") {
		dial = (dial + rotationValue)
	}


	if (dial < 0) {
		if (dialBefore != 0) {
			ans2 += 1
		}

		dial += 100
	} else if (dial >= 100) {
		if (dialBefore != 0) {
			ans2 += 1
		}

		dial -= 100
	} else if (dial == 0) {
		ans2 += 1
	}

	if (dial == 0) {
		ans1 += 1
	}

}

console.log("Puzzle 1: ", ans1)
console.log("Puzzle 2: ", ans2)

import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = [];

	const seeds = lines[0].split(": ")[1].split(" ").map(number => +number)
	let stepsToPlanted = []
	let curStep = []

	for (let r = 3; r < lines.length; r++) {
		// logic here
		const line = lines[r]

		if (line === "") {
			stepsToPlanted.push(curStep)
			curStep = []
			r++
			continue
		}

		const [destination, source, range] = line.split(" ").map(number => +number);

		curStep.push({
			destination: {
				from: destination,
				to: destination + range - 1
			},
			source: {
				from: source,
				to: source + range - 1
			},
		})
	}
	stepsToPlanted.push(curStep)

	for (const seed of seeds) {
		let curSeed = seed
		for (const step of stepsToPlanted) {
			for (const line of step) {
				if (line.source.from <= curSeed && curSeed <= line.source.to) {
					curSeed = line.destination.from + (curSeed - line.source.from)
					break
				}
			}
		}

		result.push(curSeed)
	}

	return result.reduce((min, value) => Math.min(min, value), Infinity)
}

async function startChallenge(test = false) {
	test && console.log('Running test...');
	fs.readFile(test ? 'input-test.txt' : 'input.txt', (err, inputD) => {
		if (err) throw err;
		const contentFile = inputD.toString();
		console.log("Result:", doChallenge(contentFile));
	})
}

startChallenge()
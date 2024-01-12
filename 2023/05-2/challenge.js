import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = [];

	const ranges = lines[0].split(": ")[1].split(" ").map(number => +number)
	let seeds = []
	for (let i = 0; i < ranges.length; i += 2) {
		const start = ranges[i];
		const length = ranges[i + 1];
		seeds = [...seeds, {
			from: start,
			to: start + length - 1
		}]
	}

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
		let nextSeeds = [seed]
		for (const step of stepsToPlanted) {
			const curSeeds = nextSeeds
			nextSeeds = []
			for (const curSeed of curSeeds) {
				let fromIsProcessed = false
				let toIsProcessed = false
				for (const line of step) {
					let from = -1
					let to = -1

					if (line.source.from <= curSeed.from && curSeed.from <= line.source.to) {
						from = line.destination.from + (curSeed.from - line.source.from)
						fromIsProcessed = true
					}
					if (line.source.from <= curSeed.to && curSeed.to <= line.source.to) {
						to = line.destination.from + (curSeed.to - line.source.from)
						toIsProcessed = true
					}
					if (curSeed.from <= line.source.from && line.source.to <= curSeed.to) {
						from = line.destination.from
						to = line.destination.to
					}

					if (from >= 0 || to >= 0) {
						nextSeeds.push({
							from: from >= 0 ? from : line.destination.from,
							to: to >= 0 ? to : line.destination.to
						})
					}
				}

				/**
				 * Check if the start and/or end of the seed is inside the step.
				 */
				if (!fromIsProcessed || !toIsProcessed) {
					let from = curSeed.from
					let to = curSeed.to
					let maxTo = step.reduce((max, value) => Math.max(max, value.destination.to), -Infinity)
					let minFrom = step.reduce((min, value) => Math.min(min, value.destination.from), Infinity)

					/**
					 * There are 2 cases:
					 * 1. The seed is fully outside the step
					 * 2. The seed is partially inside the step
					 * 	2.1. The start of the seed is outside the step
					 * 	2.2. The end of the seed is outside the step
					 */
					if (from > maxTo || to < minFrom) { // Case 1
						nextSeeds.push({
							from: from,
							to: to
						})
					} else { // Case 2
						if (from < minFrom) { // Case 2.1
							nextSeeds.push({
								from: from,
								to: minFrom
							})
						}
						if (to > maxTo) { // Case 2.2
							nextSeeds.push({
								from: maxTo,
								to: to
							})
						}
					}
				}
			}
		}

		result.push(nextSeeds.reduce((min, value) => Math.min(min, value.from), Infinity))
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
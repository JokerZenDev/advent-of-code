import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = [];

	const times = lines[0].replace(/\s{2,}/g, " ").split("Time: ")[1].split(" ").map(number => +number)
	const distances = lines[1].replace(/\s{2,}/g, " ").split("Distance: ")[1].split(" ").map(number => +number)

	for (let i = 0; i < times.length; i++) {
		const time = times[i]
		const distance = distances[i]
		let win = 0

		for (let j = 1; j < time; j++) {
			const curDistance = j * (time - j)

			if (curDistance > distance) {
				win++
			}
		}

		result.push(win)
	}

	return result.reduce((res, value) => res * value, 1)
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
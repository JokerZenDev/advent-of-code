import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = 0;

	for (let r = 0; r < lines.length; r++) {
		// logic here
		const line = lines[r].replace(/\s{2,}/g, " ")
		const [card, numbers] = line.split(": ")
		const [winningNumbers, drawnNumbers] = numbers.split(" | ")

		const winningNumbersArray = winningNumbers.split(" ").map(number => +number)
		const drawnNumbersArray = drawnNumbers.split(" ").map(number => +number)

		let won = 0
		for (const drawnNumber of drawnNumbersArray) {
			if (winningNumbersArray.includes(drawnNumber)) {
				won++
			}
		}

		if (won > 0) {
			result += Math.pow(2, won - 1)
		}
	}

	return result
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
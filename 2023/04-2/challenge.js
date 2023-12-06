import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = new Array(lines.length).fill(0);

	for (let r = 0; r < lines.length; r++) {
		// logic here
		const line = lines[r].replace(/\s{2,}/g, " ")
		const numbers = line.split(": ")[1]
		const [winningNumbers, drawnNumbers] = numbers.split(" | ")

		const winningNumbersArray = winningNumbers.split(" ").map(number => +number)
		const drawnNumbersArray = drawnNumbers.split(" ").map(number => +number)

		let won = 0
		for (const drawnNumber of drawnNumbersArray) {
			if (winningNumbersArray.includes(drawnNumber)) {
				won++
			}
		}

		result[r]++

		for (let i = 1; i + r < lines.length && i <= won; i++) {
			result[i + r] += result[r]
		}
	}

	return result.reduce((sum, value) => sum + value, 0)
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
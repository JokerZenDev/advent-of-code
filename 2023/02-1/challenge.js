import fs from 'fs'

const TOTAL_CUBES = {
	"red": 12,
	"green": 13,
	"blue": 14
}

function doChallenge(input) {
	const lines = input.split('\n');
	let result = 0;
	for (const line of lines) {
		// Logic here
		const [game, subsets] = line.split(': ');
		const gameId = +game.split(' ')[1];

		let isValid = true

		subsets.split("; ").forEach(subset => {
			subset.split(', ').forEach(cube => {
				const [number, color] = cube.split(' ');
				if (TOTAL_CUBES[color] < +number) {
					isValid = false
				}
			})
		})

		if (isValid) {
			result += gameId
		}
	}

	return result
}

async function startChallenge(test = false) {
	test && console.log('Running test...');
	fs.readFile(test ? 'input-test.txt' : 'input.txt', (err, inputD) => {
		if (err) throw err;
		const contentFile = inputD.toString();
		console.log("Result: ", doChallenge(contentFile));
	})
}

startChallenge();
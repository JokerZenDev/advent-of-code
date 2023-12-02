import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = 0;
	for (const line of lines) {
		// Logic here
		const [game, subsets] = line.split(': ');
		const gameId = +game.split(' ')[1];

		let isValid = true

		const max_cubes = {
			red: 0,
			green: 0,
			blue: 0
		}

		subsets.split("; ").forEach(subset => {
			subset.split(', ').forEach(cube => {
				const [number, color] = cube.split(' ');
				max_cubes[color] = Math.max(max_cubes[color], +number)
			})
		})

		result += max_cubes.blue * max_cubes.green * max_cubes.red
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
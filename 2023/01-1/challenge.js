import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = 0;
	for (const line of lines) {
		// Logic here
		let first = null
		let last = null
		for (let i = 0; i < line.length; i++) {
			const char = line[i];

			if (+char >= 0) {
				if (first === null) {
					first = char;
				}
				last = char;
			}
		}

		result += +(`${first}${last}`);
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
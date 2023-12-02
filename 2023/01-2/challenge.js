import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = 0;
	for (const line of lines) {
		// Logic here
		let first = null
		let last = null
		let lastNumberPosition = -1
		for (let i = 0; i < line.length; i++) {
			const char = line[i];

			let number = null
			if (+char >= 0) {
				lastNumberPosition = i

				number = +char
			} else {
				const minPosition = i - 2 // Because the minumum letters of numbers is 3
				for (let j = lastNumberPosition + 1; j <= minPosition && number === null; j++) {
					const substring = line.substring(j, i + 1)

					switch (substring) {
						case 'one':
							number = 1
							break;
						case 'two':
							number = 2
							break;
						case 'three':
							number = 3
							break;
						case 'four':
							number = 4
							break;
						case 'five':
							number = 5
							break;
						case 'six':
							number = 6
							break;
						case 'seven':
							number = 7
							break;
						case 'eight':
							number = 8
							break;
						case 'nine':
							number = 9
							break;
					}
				}
			}

			if (number !== null) {
				if (first === null) {
					first = number;
				}
				last = number;
			}
		}

	console.log("Number: ", `${first}${last}`);

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
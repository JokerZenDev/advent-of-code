import fs from 'fs'

function isValidNumber(start, end, row, array) {
	const minRow = row === 0 ? 0 : row - 1
	const maxRow = row === array.length - 1 ? row : row + 1
	const minCol = start === 0 ? 0 : start - 1
	const maxCol = end === array[row].length - 1 ? end : end + 1

	for (let r = minRow; r <= maxRow; r++) {
		for (let c = minCol; c <= maxCol; c++) {
			if (r === row && c >= start && c <= end) {
				continue
			}

			const char = array[r][c]
			if (isNaN(+char) && char !== '.') {
				return true
			}
		}
	}

	return false
}

function doChallenge(input) {
	const lines = input.split('\n');
	let result = 0;
	let posStartNumber = -1

	for (let r = 0; r < lines.length; r++) {
		const line = lines[r];

		if (posStartNumber >= 0) {
			if (isValidNumber(posStartNumber, line.length - 1, r - 1, lines)) {
				result += +lines[r - 1].substring(posStartNumber, line.length)
			}

			posStartNumber = -1
		}

		for (let c = 0; c < line.length; c++) {
			const char = line[c];

			if (isNaN(+char) && posStartNumber >= 0) {
				if (isValidNumber(posStartNumber, c - 1, r, lines)) {
					result += +line.substring(posStartNumber, c)
				}

				posStartNumber = -1
			} else if (!isNaN(+char) && posStartNumber < 0) {
				posStartNumber = c
			}
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

startChallenge()
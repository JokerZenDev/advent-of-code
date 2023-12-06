import fs from 'fs'

function isValidGear(col, row, array) {
	const minRow = row === 0 ? 0 : row - 1
	const maxRow = row === array.length - 1 ? row : row + 1
	const minCol = col === 0 ? 0 : col - 1
	const maxCol = col === array[row].length - 1 ? col : col + 1

	let posNumberFound = []

	for (let r = minRow; r <= maxRow; r++) {
		let justFound = false

		for (let c = minCol; c <= maxCol; c++) {
			if (r === row && c === col) {
				justFound = false
				continue
			}

			const char = array[r][c]
			if (!isNaN(+char) && !justFound) {
				posNumberFound = [...posNumberFound, { col: c, row: r }]
				justFound = true
			} else if (isNaN(+char)) {
				justFound = false
			}
		}
	}

	return posNumberFound
}

function fetchNumber(col, row, array) {
	let start = col
	let end = col

	while (start - 1 >= 0 && !isNaN(+array[row][start - 1])) {
		start--
	}

	while (end + 1 < array[row].length && !isNaN(+array[row][end + 1])) {
		end++
	}

	return +array[row].substring(start, end + 1)
}

function doChallenge(input) {
	const lines = input.split('\n');
	let result = 0;

	for (let r = 0; r < lines.length; r++) {
		const line = lines[r];

		for (let c = 0; c < line.length; c++) {
			const char = line[c];

			if (char === "*") {
				const posNumberFound = isValidGear(c, r, lines)

				if (posNumberFound.length === 2) {
					result += fetchNumber(posNumberFound[0].col, posNumberFound[0].row, lines) * fetchNumber(posNumberFound[1].col, posNumberFound[1].row, lines)
				}
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
import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = [];

	const time = +lines[0].replace(/\s{2,}/g, " ").split("Time: ")[1].replace(/\s/g, "")
	const distance = +lines[1].replace(/\s{2,}/g, " ").split("Distance: ")[1].replace(/\s/g, "")

	const range = {
		start: 1,
		end: time - 1
	}

	console.log(distance)

	while (range.end - range.start > 1) {
		const speed = Math.floor((range.end - range.start) / 2) + range.start
		const curDistance = speed * (time - speed)

		if (curDistance > distance) {
			range.end = speed
		} else {
			range.start = speed
		}

		console.log(range)
	}

	return time - range.start * 2 - 1
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
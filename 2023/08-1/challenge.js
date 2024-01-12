import fs from 'fs'

function doChallenge(input) {
	const lines = input.split('\n');
	let result = 0

	const directions = lines[0]
	const nodes = {}

	for (let i = 2; i < lines.length; i++) {
		const line = lines[i]

		const [nodeName, children] = line.split(' = ')
		const [childL, childR] = children ? children.replace(/[\(\)]/g, "").split(', ') : []

		nodes[nodeName] = {
			left: childL,
			right: childR
		}
	}

	let curNode = 'AAA'

	while (curNode !== 'ZZZ') {
		const node = nodes[curNode]
		if (directions[result % directions.length] === 'L') {
			curNode = node.left
		} else {
			curNode = node.right
		}
		result++
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
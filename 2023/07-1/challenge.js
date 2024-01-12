import fs from 'fs'

const cardValue = {
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12
}

function doChallenge(input) {
	const lines = input.split('\n');
	let hands = []

	for (let i = 0; i < lines.length; i++) {
		const line = lines[i]

		const [hand, bid] = line.split(' ')

		const cards = {}

		for (let j = 0; j < hand.length; j++) {
			const card = hand[j]
			cards[card] = (cards[card] || 0) + 1
		}

		const handValue = Object.values(cards).reduce((res, value) => Math.pow(value, 2) + res, 0)

		hands.push({
			hand,
			handValue,
			bid
		})
	}

	hands.sort((a, b) => {
		if (b.handValue == a.handValue) {
			for (let i = 0; i < a.hand.length; i++) {
				const cardA = a.hand[i]
				const cardB = b.hand[i]

				if (cardValue[cardA] != cardValue[cardB]) {
					return cardValue[cardA] - cardValue[cardB]
				}
			}
		}

		return a.handValue - b.handValue
	})

	return hands.reduce((res, value, idx) => res + +value.bid * (idx + 1), 0)
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
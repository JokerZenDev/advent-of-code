import fs from 'fs'

const cardValue = {
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
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

		if (cards['J'] > 0) {
			const sortedCards = [{
				card: 'K',
				num: 0
			}]
			for (const card in cards) {
				if (card === 'J') {
					continue
				}
				sortedCards.push({
					card,
					num: cards[card]
				})
			}
			sortedCards.sort((a, b) => {
				if (a.num === b.num) {
					return cardValue[b.card] - cardValue[a.card]
				}

				return b.num - a.num
			})

			cards[sortedCards[0].card] = (cards[sortedCards[0].card] || 0) + cards['J']

			delete cards['J']
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
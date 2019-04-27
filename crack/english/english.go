package english

import (
	"fmt"
	"unicode"

	"github.com/Xjs/cryptopals/statistics"
	"github.com/Xjs/cryptopals/xor"
)

// GetRuneFrequency returns a histogram of all runes in the text
func GetRuneFrequency(text string) map[rune]int {
	result := make(map[rune]int)
	for _, letter := range text {
		result[letter]++
	}
	return result
}

// GetLetterFrequency returns a histogram of all letters in the text (case-sensitive)
func GetLetterFrequency(text string) map[rune]int {
	result := make(map[rune]int)
	for _, letter := range text {
		if !unicode.IsLetter(letter) {
			continue
		}
		result[letter]++
	}
	return result
}

// GetLowerLetterFrequency returns a histogram of all letters in the text (case-insensitive)
func GetLowerLetterFrequency(text string) map[rune]int {
	result := make(map[rune]int)
	for _, letter := range text {
		if !unicode.IsLetter(letter) {
			continue
		}

		result[unicode.ToLower(letter)]++
	}
	return result
}

// GetScore returns how well the text scores against the given reference
// by dividing the cumulative frequencies to the reference frequencies
func GetScore(text string, reference map[rune]int) statistics.Score {
	var cumulativeFrequency int
	for _, letter := range text {
		cumulativeFrequency += reference[letter]
	}
	return statistics.RelativeScore(cumulativeFrequency, len(text))
}

// TryKeylen tries to decrypt the given input with the given keylen,
// assuming english text.
//
// The input will be split in keylen blocks, starting at 0..keylen-1 and
// always skipping keylen bytes; this assumes a rotating key of given length.
// Decryption of each block will be attempted with the FindSingleByteKey function.
// If at least keylen/2 blocks can be decoded with a high enough score,
// TryKeylen will return nil as third result. Otherwise, decryption will still
// be performed, and results will be returned, but the third parameter will
// contain an non-nil error.
func TryKeylen(input []byte, keylen int) ([]byte, int, error) {
	blocks := make([][]byte, keylen)

	for i, b := range input {
		blocks[i%keylen] = append(blocks[i%keylen], b)
	}

	english := make([][]byte, keylen)

	var errors int
	for i, block := range blocks {
		en, _, _, err := FindSingleByteKey(block)
		if err != nil {
			errors++
		}
		english[i] = []byte(en)
	}

	decrypted := make([]byte, len(input))
	for i, block := range english {
		for j, b := range block {
			decrypted[i+(j*keylen)] = b
		}
	}

	if errors >= keylen/2 {
		return decrypted, errors, fmt.Errorf("%d out of %d blocks look non-english", errors, keylen)
	}

	return decrypted, errors, nil
}

// ReferenceHistogram is the reference English histogram that is used for scoring texts
// for their english-ness
var ReferenceHistogram map[rune]int

// ReferenceScore is the score of the reference text against itself
var ReferenceScore statistics.Score

// ScoreDistance is an empirical value that denotes how much distance a text may have to the reference
// score to be still considered English.
// Currently chosen via gut feeling.
const ScoreDistance statistics.Score = 50

// IsEnglishScore returns true if a text with the given score is considered English
func IsEnglishScore(score statistics.Score) bool {
	return ReferenceScore-ScoreDistance < score && score < ReferenceScore+ScoreDistance
}

// IsEnglish returns whether a text is considered modern English based on the rune frequencies
func IsEnglish(text string) bool {
	return IsEnglishScore(GetScore(text, ReferenceHistogram))
}

func init() {
	ReferenceHistogram = GetRuneFrequency(EnglishText)
	ReferenceScore = GetScore(EnglishText, ReferenceHistogram)
}

// FindSingleByteKey attempts to decrypt a text assumed to be english, xor-encrypted
// with a single byte key. It will return the decrypted text, the key byte, the
// score of the text (referenceScore)
func FindSingleByteKey(input []byte) (string, byte, statistics.Score, error) {
	scores := make(map[byte]statistics.Score)

	frequencies := statistics.NewByteHistogramFromRunes(GetRuneFrequency(string(input)))
	referenceFrequencies := statistics.NewByteHistogramFromRunes(ReferenceHistogram)

	const maxDistance = 3

	for _, tuple := range [][2]int{{0, 0}, {1, 1}, {2, 2}, {0, 1}, {1, 0}, {1, 2}, {2, 1}} {
		encryptedB := frequencies.GetHigh(tuple[0]).Byte
		plaintextB := referenceFrequencies.GetHigh(tuple[1]).Byte

		b := encryptedB ^ plaintextB

		candidate := string(xor.Single(input, b))
		scores[b] = GetScore(candidate, ReferenceHistogram)
	}

	hist := statistics.NewByteScoreHistogram(scores)
	high := hist.GetHigh(0)

	result := string(xor.Single(input, high.Byte))

	if high.Score < 350.0 {
		return result, high.Byte, high.Score, fmt.Errorf("%q has score %.f, probably not english", result, high.Score)
	}

	return result, high.Byte, high.Score, nil
}

// EnglishText is a sample modern English text.
// Copied from https://www.public.asu.edu/~gelderen/hel/lchatterly.html
const EnglishText = `She saw a secret little clearing, and a secret little hot made of rustic poles. And she had never been here before! She realized it was the quiet place where the growing pheasants were reared; the keeper in his shirt‑sleeves was kneeling, hammering. The dog trotted forward with a short, sharp bark, and the keeper lifted his face suddenly and saw her. He had a startled look in his eyes.

He straightened himself and saluted, watching her in silence, as she came forward with weakening limbs. He resented the intrusion; he cherished his solitude as his only and last freedom in life.

"I wondered what the hammering was," she said, feeling weak and breathless, and a little afraid of him, as he looked so straight at her.

"Ah'm gettin' th' coops ready for th' young bods,'" he said, in broad vernacular.

She did not know what to say, and she felt weak.

"I should like to sit down a bit," she said.

"Come and sit 'ere i' th' 'ut," he said, going in front of her to the hut, pushing aside some timber and stuff, and drawing out a rustic chair, made of hazel sticks.

"Am Ah t' light yer a little fire?" he asked, with the curious naïveté of the dialect.

"Oh, don't bother," she replied.

But he looked at her hands; they were rather blue. So he quickly took some larch twigs to the little brick fire‑place in the corner, and in a moment the yellow flame was running up the chimney. He made a place by the brick hearth.

"Sit 'ere then a bit, and warm yer,'" he said.

A wet brown dog came running and did not bark, lifting a wet feather of a tail. The man followed in a wet black oilskin jacket, like a chauffeur, and face flushed a little. She felt him recoil in his quick walk, when he saw her. She stood up in the handbreadth of dryness under the rustic porch. He saluted without speaking, coming slowly near. She began to withdraw.

"I'm just going," she said.

"Was yer waitin' to get in?" he asked, looking at the hut, not at her.

"No, I only sat a few minutes in the shelter," she said, with quiet dignity.

He looked at her. She looked cold.

"Sir Clifford 'adn't got no other key then?" he asked.

"No, but it doesn't matter. I can sit perfectly dry under this porch. Good afternoon!" She hated the excess of vernacular in his speech.

He watched her closely, as she was moving away. Then he hitched up his jacket, and put his hand in his breeches pocket, taking out the key of the hut.

"'Appen yer'd better 'ave this key, an' Ah min fend for t' bods some other road."

She looked at him.

"What do you mean?" she asked.

"I mean as 'appen Ah can find anuther pleece as'll du for rearin' th' pheasants. If yer want ter be 'ere, yo'll non want me messin' abaht a' th' time."

She looked at him, getting his meaning through the fog of the dialect.

"Why don't you speak ordinary English?" she said coldly.

"Me! Ah thowt it wor ordinary."

She was silent for a few moments in anger.

"So if yer want t' key, yer'd better tacit. Or 'appen Ah'd better gi'e 't yer termorrer, an' clear all t' stuff aht fust. Would that du for yer?'"

She became more angry.

"I didn't want your key," she said. "I don't want you to clear anything out at all. I don't in the least want to turn you out of your hut, thank you! I only wanted to be able to sit here sometimes, like today. But I can sit perfectly well under the porch, so please say no more about it."

He looked at her again, with his wicked blue eyes.

"Why," he began, in the broad slow dialect. "Your Ladyship's as welcome as Christmas ter th' hut an' th' key an' iverythink as is. On'y this time O' th' year ther's bods ter set, an' Ah've got ter be potterin' abaht a good bit, seein' after 'em, an' a'. Winter time Ah ned 'ardly come nigh th' pleece. But what wi' spring, an' Sir Clifford wantin' ter start th' pheasants. . .An' your Ladyship'd non want me tinkerin' around an' about when she was 'ere, all the time."

She listened with a dim kind of amazement.

"Why should I mind your being here?" she asked.

He looked at her curiously.

"T'nuisance on me!" he said briefly, but significantly. She flushed. "Very well!" she said finally. "I won't trouble you. But I don't think I should have minded at all sitting and seeing you look after the birds. I should have liked it. But since you think it interferes with you, I won't disturb you, don't be afraid. You are Sir Clifford's keeper, not mine."

The phrase sounded queer, she didn't know why. But she let it pass.

"Nay, your Ladyship. It's your Ladyship's own 'ut. It's as your Ladyship likes an' pleases, every time. Yer can turn me off at a wik's notice. It wor only. . ."

"Only what?" she asked, baffled.

He pushed back his hat in an odd comic way.

"On'y as 'appen yo'd like the place ter yersen, when yer did come, an' not me messin' abaht.'"

"But why?" she said, angry. "Aren't you a civilized human being? Do you think I ought to be afraid of you? Why should I take any notice of you and your being here or not? Why is it important?"

He looked at her, all his face glimmering with wicked laughter.

"It's not, your Ladyship. Not in the very least," he said.

"Well, why then?" she asked.

"Shall I get your Ladyship another key then?"

"No thank you! I don't want it."

"Ah'll get it anyhow. We'd best 'ave two keys ter th' place."

"And I consider you are insolent," said Connie, with her colour up, panting a little.

"Nay, nay!" he said quickly. "Dunna yer say that! Nay, nay! I niver meant nuthink. Ah on'y thought as if yo' come 'ere, Ah s'd ave ter clear out, an' it'd mean a lot of work, settin' up somewheres else. But if your Ladyship isn't going ter take no notice O' me, then. . .it's Sir Clifford's 'ut, an' everythink is as your Ladyship likes, everythink is as your Ladyship likes an' pleases, barrin' yer take no notice O' me, doin' th' bits of jobs as Ah've got ter do."

Connie went away completely bewildered. She was not sure whether she had been insulted and mortally of fended, or not. Perhaps the man really only meant what he said; that he thought she would expect him to keep away. As if she would dream of it! And as if he could possibly be so important, he and his stupid presence.

`

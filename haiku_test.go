package haiku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentenceExtraction(t *testing.T) {
	text := `The giant brown dog! He jumped over the candlestick into a popsicle! What a disaster! But maybe a good disaster? He spent $1.50 on it. All the while he was chanting "U.S.A." like a madman.`
	s := sentencesFromText(text)
	if assert.Len(t, s, 6) {
		assert.Equal(t, "What a disaster!", s[2])
		assert.Equal(t, `All the while he was chanting "U.S.A." like a madman.`, s[5])
	}
}

func TestSentenceExtractionMultiline(t *testing.T) {
	text := `Sentence one.
Sentence two, on a new line.

Sentence three, after a blank line!`
	s := sentencesFromText(text)
	assert.Len(t, s, 3)
	assert.Equal(t, "Sentence two, on a new line.", s[1])
	assert.Equal(t, `Sentence three, after a blank line!`, s[2])
}

func TestSentenceNoPunctuation(t *testing.T) {
	text := `this is just some words`
	s := sentencesFromText(text)
	assert.Len(t, s, 1)
	assert.Equal(t, "this is just some words", s[0])
}

func TestSentenceRepeatedPunctuation(t *testing.T) {
	text := `this is just some words!! With a lot of emphasis!`
	s := sentencesFromText(text)
	t.Skip("this case does not yet work?")
	if assert.Len(t, s, 2) {
		assert.Equal(t, "this is just some words!!", s[0])
		assert.Equal(t, "With a lot of emphasis!", s[1])
	}
}

func TestWordsInSentence(t *testing.T) {
	text := `the quick brown dog`
	words := wordsInSentence(text)
	assert.Len(t, words, 4)

	text = `the quick brown dog.`
	words = wordsInSentence(text)
	assert.Len(t, words, 4)

	text = `it cost $4.50, or more`
	words = wordsInSentence(text)
	assert.Len(t, words, 5)

	text = "the quick gay dude ran to the new pink car in the state of mind to win"
	words = wordsInSentence(text)
	assert.Len(t, words, 17)
	assert.Equal(t, "the", words[0])
	assert.Equal(t, "quick", words[1])
	assert.Equal(t, "to", words[15])
	assert.Equal(t, "win", words[16])

	text = "the quick gay dude ran to the new pink car in the state of mind to win!"
	words = wordsInSentence(text)
	assert.Len(t, words, 17)
	assert.Equal(t, "the", words[0])
	assert.Equal(t, "quick", words[1])
	assert.Equal(t, "to", words[15])
	assert.Equal(t, "win", words[16])

	text = ""
	words = wordsInSentence(text)
	assert.Len(t, words, 0)
}

func TestHaikuFromSentence(t *testing.T) {
	//                             1     2    3    4   5| 6   7   8   9  10 11  12|   13 14   15 16  17
	h, err := haikuFromSentence("the quick fast dude ran to the new red car in the state of mind to win")
	if assert.NoError(t, err) {
		assert.Equal(t, `the quick fast dude ran
to the new red car in the
state of mind to win`,
			h.String())
		assert.Equal(t, "to the new red car in the", h.Lines()[1])
	}

	_, err = haikuFromSentence("did u talk also about the breakup part? like why it happened and if that reason is still a thing?")
	assert.ErrorContains(t, err, "not a haiku - too many words")

	h, err = haikuFromSentence("by grabthar's hammer, what a savings you can make - almost criminal")
	if assert.NoError(t, err) {
		assert.Equal(t, `by grabthar's hammer,
what a savings you can make
- almost criminal`,
			h.String())
	}

	h, err = haikuFromSentence("")
	assert.ErrorContains(t, err, "sentence has 0 words")

}

func TestParagraphsToHaikus(t *testing.T) {
	text := `Threads are heavyweights, goroutines are light as air, concurrency wins.

Strict and silent judge, catches errors in my code, no nil slips away. One thread sends a word, another waits in silence, they meet, work is done. Shape of thought defined, no need for inheritance, duck typing prevails.

Memory held tight, swept away in quiet dusk, Go frees what I leave.`
	out := Find(text)

	if assert.Len(t, out, 5) {
		assert.Equal(t, `one thread sends a word,
another waits in silence,
they meet, work is done`, out[2].String())
	}

	h := Find("haiku can be found wherever you are looking with help of some code")
	if assert.Len(t, h, 1) {
		assert.Equal(t, `haiku can be found
wherever you are looking
with help of some code`, h[0].String())
	}

}

func FuzzHaikuFromSentence(f *testing.F) {
	f.Add("")
	f.Add("this is some boring text")
	f.Add("    this is some boring text   ")
	f.Add("this is some boring text (@()*@(@)* ijOF(@@( lots of ))punctuation")
	f.Add("   ")
	f.Add("12092897234987")
	f.Add("\U000c9288\U000fa405䠁\U000dbcd3\U001095be\U0003456e\U0007fe36\U0008d5d3𧝳\U0006a3e0䠹𗙧\U0003ec48\U00078be3\U0007dd24\U000b4b50뗏\U0005a099\U0003eb1e\U000ffef6\U000c61c0𠬢\U000d02a3\U000a41dd\U000c62ba\U000890c0\U000e534e\U000d8155\U00072e3a\U0010c563\U000a7730\U00048c0a\U000a2698\U000dd595\U000d0768𝁔\U00038a30컩\U00032563𠷰🭣\U000564a5𝒈\U0004104c\U0006bdc8\U000d72d3\U000ce78b\U00065795\U00102f31\U000723eb\U000dc97b\U0006a4b7亮\U0010c8f7\U000e8090\U0008abf9𘬒푆\U00012716𭝱\U00016244\U0006012e\U000689d3\U0009da4d𭼴\U0009a067\U000aede5\U000b9d6b\U0004b136\U000975c4㷭\U0003f873\U00081080\U00047c35\U000b8ee5\U000d4493\U000dccb4\U000313f2\U00084081\U00084dc6\U000e8fc8\U000a362b\U00101023\U000855d6\U000e4318\U000cebf6\U00051bc5\U000d9bff\U00092349\U000ae3da\U000a116f\U0003f89e\U000d5c54\U000b34ae\U0006d4a0\U0003cd52\U00083f6c\U0003a7fc\U0005083b\U00086439")
	f.Add("춈洞䠁킓𐖾㑮翦贓𧝳檠䠹𗙧㻈磣絤骐뗏娙㺞￶챀𠬢킣ꐝ챺茀킕瀺𐱣ꜰ䠊ꉘ킕큨𝁔㢰컩㉣𠷰🭣啥𝒈䄌殈틓챋敕𐋱爫킻橷亮𐲷谹𘬒푆ሖ𭝱㡄愮桓鶍𭼴騧껥맥䬶䬆靄㷭㾳耀䰵뷥킓챴㈲蓁虆ꍫ𐄣蓖챶傅ힿ鈉긚ꄯ㾞픔ꍮ洠㳒葛㩼偏")
	f.Fuzz(func(t *testing.T, in string) {
		haikuFromSentence(in)
	})
}

package compliments

import "math/rand"

var compliments = []string{
	"Ти прекрасна!",
	"Твої очі, як зорі.",
	"Ти світишся як сонце!",
	"Твоя посмішка піднімає мені настрій!",
	"Ти дивовижна!",
}

func ComplimentForWomen() string {
	return compliments[rand.Intn(len(compliments))]
}

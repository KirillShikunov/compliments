package compliments

import "math/rand"

var complimentsForWomen = []string{
	"Ти прекрасна!",
	"Твої очі, як зорі.",
	"Ти світишся як сонце!",
	"Твоя посмішка піднімає мені настрій!",
	"Ти дивовижна!",
}

var complimentsForMen = []string{
	"Ти справжній чоловік!",
	"Твоя впевненість вражає!",
	"Ти завжди такий розумний!",
	"Твоя сміливість надихає!",
	"Ти неймовірно харизматичний!",
}

func ComplimentForWomen() string {
	return complimentsForWomen[rand.Intn(len(complimentsForWomen))]
}

func ComplimentForMen() string {
	return complimentsForMen[rand.Intn(len(complimentsForMen))]
}

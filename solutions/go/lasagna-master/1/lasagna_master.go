package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, mins int) int {
    if mins == 0 {
        return len(layers) * 2
    }
    return len(layers) * mins
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodleQnt int, sauceQnt float64) {
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodleQnt += 50
        }
        if layers[i] == "sauce" {
            sauceQnt += 0.2
        }
    }

    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    myList[len(myList) - 1] = friendList[len(friendList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64{
    scaled := make([]float64, len(quantities))
    factor := float64(portions) / 2.0
    for i := 0; i < len(quantities); i++ {
        scaled[i] = quantities[i] * factor
    }

    return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.

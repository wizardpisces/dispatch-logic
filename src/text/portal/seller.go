package portal
import "fmt"

// Hello returns a greeting for the named person.
func HelloSeller(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi Seller, %v!", name)
    return message
}
package portal
import "fmt"

// Hello returns a greeting for the named person.
func HelloNormal(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi,Normal %v!", name)
    return message
}
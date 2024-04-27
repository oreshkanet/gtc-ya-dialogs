package alice

import "fmt"

func assertInitialized(component interface{}, name string) {
	if component == nil {
		panic(fmt.Sprintf("%s wasn't initialized before usage", name))
	}
}

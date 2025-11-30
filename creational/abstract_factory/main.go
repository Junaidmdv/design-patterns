package main


// The Abstract Factory provides an interface for creating families of related objects without specifying their concrete types.

// Great when:
// You want to switch between product families at runtime
// You want to enforce consistency between related objects
// You want to decouple creation from usage



import "fmt"

//ui factory

type Button interface {
	Render() string
}

type CheckBox interface {
	Checker() string
}

type UiFactory interface {
	CreateButton() Button
	CreateCheckBox() CheckBox
}

//mac-os  ui implementation

type MacButton struct{}

func (m MacButton) Render() string {
	return "Mac os button Rendering"
}

type MacCheckButton struct{}

func (m MacCheckButton) Checker() string {
	return "Mac os checking the checkbox option"
}

type MacUIFactory struct{}

func (mi MacUIFactory) CreateButton() Button {
	return MacButton{}
}

func (mi MacUIFactory) CreateCheckBox() CheckBox {
	return MacCheckButton{}
}


// windows  implementation

type WindowsButton struct{}

func (wb WindowsButton) Render() string {
	return "rendering the windows os button"
}

type WindowsCheckBox struct{}

func (wc WindowsCheckBox) Checker() string {
	return "checking options in the windows os checkbox"
}

type WindowsUIFactory struct{}

func (wi WindowsUIFactory) CreateButton() Button {
	return WindowsButton{}
}

func (wi WindowsUIFactory) CreateCheckBox() CheckBox {
	return WindowsCheckBox{}
}

func main() {
	var factory UiFactory
	factory = MacUIFactory{}

	button := factory.CreateButton()
	checkbox := factory.CreateCheckBox()

	fmt.Println(button.Render())
	fmt.Println(checkbox.Checker())

}

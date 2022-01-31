package app

import (
	"log"
	"os/exec"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type Window struct {
	*gtk.ApplicationWindow
	builder *gtk.Builder

	backButton *gtk.Button
	nextButton *gtk.Button
	stack      *gtk.Stack
	pageIndex  int
	pages      []*gtk.Widget
}

const URL = "https://api.rlxos.dev/system/submit"

func New(builder *gtk.Builder) (*Window, error) {
	window := &Window{
		builder: builder,
	}
	window.ApplicationWindow = window.getWidget("mainWindow").(*gtk.ApplicationWindow)
	window.backButton = window.getWidget("backButton").(*gtk.Button)
	window.nextButton = window.getWidget("nextButton").(*gtk.Button)
	window.stack = window.getWidget("stack").(*gtk.Stack)

	builder.ConnectSignals(map[string]interface{}{
		"onNextButtonClicked": window.onNextButtonClicked,
		"onBackButtonClicked": window.onBackButtonClicked,
		"onBazaarClicked": func() {
			window.openUrl("https://rlxos.dev/apps")
		},
		"onAppImageHubClicked": func() {
			window.openUrl("https://www.appimagehub.com/")
		},
		"onFlathubClicked": func() {
			window.openUrl("https://flathub.org/")
		},
		"onDocsClicked": func() {
			window.openUrl("https://docs.rlxos.dev")
		},
		"onCommunityClicked": func() {
			window.openUrl("https://discord.com/invite/aEez9XhE2y")
		},
		"onChannelClicked": func() {
			window.openUrl("https://www.youtube.com/channel/UC2Dd2tHneeaxvhqz_IaO9bw")
		},

		// "onViewDataClicked": func() {
		// 	dialog, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	dialog.SetTitle("System Information")

		// 	textView, _ := gtk.TextViewNew()
		// 	textView.SetEditable(false)
		// 	textView.SetCursorVisible(false)
		// 	textView.SetMonospace(true)
		// 	dialog.Add(textView)

		// 	buffer, _ := textView.GetBuffer()
		// 	buffer.SetText(window.systemInformation())
		// 	dialog.ShowAll()
		// },
		// "onSubmitClicked": func() {
		// 	jsonString := []byte(`{"data": "` + window.systemInformation() + `"}`)
		// 	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonString))
		// 	if err != nil {
		// 		window.showError("failed to create new POST request " + err.Error())
		// 		return
		// 	}

		// 	req.Header.Set("Content-Type", "application/json")
		// 	client := &http.Client{}
		// 	resp, err := client.Do(req)
		// 	if err != nil {
		// 		window.showError("failed to submit POST request " + err.Error())
		// 		return
		// 	}
		// 	defer resp.Body.Close()
		// },
		"onEnjoyClicked": func() {
			err := exec.Command("gnome-terminal", "-e", "pkexec update-grub").Run()
			if err != nil {
				window.showError("failed to update grub " + err.Error())
			}

			app, _ := window.GetApplication()
			app.Quit()
		},
	})

	window.stack.GetChildren().Foreach(func(item interface{}) {
		window.pages = append(window.pages, item.(*gtk.Widget))
	})
	window.pageIndex = 0

	window.updateButtons()

	return window, nil
}

func (win Window) updateButtons() {
	if win.pageIndex <= 0 {
		win.backButton.SetSensitive(false)
	} else {
		win.backButton.SetSensitive(true)
	}

	if win.pageIndex >= len(win.pages)-1 {
		win.nextButton.SetSensitive(false)
	} else {
		win.nextButton.SetSensitive(true)
	}
}

func (win Window) getWidget(id string) glib.IObject {
	widget, err := win.builder.GetObject(id)
	if err != nil {
		log.Fatal(err)
	}

	return widget
}

func (win *Window) onNextButtonClicked() {
	win.pageIndex += 1
	win.stack.SetVisibleChild(win.pages[win.pageIndex])

	win.updateButtons()
}

func (win *Window) onBackButtonClicked() {
	win.pageIndex -= 1
	win.stack.SetVisibleChild(win.pages[win.pageIndex])

	win.updateButtons()
}

func (win *Window) openUrl(url string) {
	exec.Command("xdg-open", url).Start()
}

// func (win *Window) systemInformation() string {
// 	data, err := exec.Command("inxi", "-F", "-c", "0").Output()
// 	if err != nil {
// 		return err.Error()
// 	}

// 	return string(data)
// }

func (win *Window) showError(mesg string) {
	dialog := gtk.MessageDialogNew(&win.Window, gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_ERROR, gtk.BUTTONS_CLOSE, mesg)
	dialog.Run()

	dialog.Destroy()
}

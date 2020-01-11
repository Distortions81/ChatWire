package admin

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"../../support"
	"github.com/bwmarrin/discordgo"
)

func Preview(s *discordgo.Session, m *discordgo.MessageCreate) {

	var filename = ""
	t := time.Now()
	ourseed := fmt.Sprintf("%ld", t.Unix())

	path := fmt.Sprintf("/home/fact/map-prev/%s.png", ourseed)
	strseed := fmt.Sprintf("%d", ourseed)

	args := []string{"--generate-map-preview", path, "--preset", support.Config.MapPreset, "--map-gen-seed", strseed}

	cmd := exec.Command(support.Config.MapGenExec, args...)

	//Debug
	support.Log(fmt.Sprintf("Ran: %s %s", support.Config.MapGenExec, strings.Join(args, ", ")))

	out, aerr := cmd.CombinedOutput()

	if aerr != nil {
		support.ErrorLog(aerr)
	}

	lines := strings.Split(string(out), "\n")
	support.Log(lines[0])

	for _, l := range lines {
		if strings.Contains(l, "Wrote map preview image file:") {
			result := regexp.MustCompile(`(?m).*Wrote map preview image file: \/home\/fact\/(.*)`)
			filename = result.ReplaceAllString(l, "http://bhmm.net/${1}")
		}
	}

	buffer := "Preview failed."
	if filename != "" {
		buffer = fmt.Sprintf("Preview: %s", filename)
	}

	_, err := s.ChannelMessageSend(support.Config.FactorioChannelID, buffer)
	if err != nil {
		support.ErrorLog(err)
	}
	return
}

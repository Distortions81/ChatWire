package user

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"../../cfg"
	"../../constants"
	"../../disc"
	"../../fact"
	"../../glob"
	"../../sclean"
	"github.com/bwmarrin/discordgo"
)

//Last Seen
type ByLastSeen []glob.PList

func (a ByLastSeen) Len() int           { return len(a) }
func (a ByLastSeen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastSeen) Less(i, j int) bool { return a[i].LastSeen > a[j].LastSeen }

//Created time
type ByNew []glob.PList

func (a ByNew) Len() int           { return len(a) }
func (a ByNew) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNew) Less(i, j int) bool { return a[i].Creation > a[j].Creation }

func levelToString(level int) string {

	name := "Error"

	if level <= -254 {
		name = "Deleted"
	} else if level == -1 {
		name = "Banned"
	} else if level == 0 {
		name = "New"
	} else if level == 1 {
		name = "Member"
	} else if level == 2 {
		name = "Regular"
	} else if level >= 255 {
		name = "Admin"
	}

	return name
}

func CheckAdmin(ID string) bool {
	for _, admin := range cfg.Global.AdminData.IDs {
		if ID == admin {
			return true
		}
	}
	return false
}

func Whois(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	maxresults := constants.WhoisResults
	if CheckAdmin(m.Author.ID) {
		maxresults = constants.AdminWhoisResults
	}
	var slist []glob.PList
	argnum := len(args)

	//Reconstruct list, to remove empty entries and to reduce lock time
	glob.PlayerListLock.RLock()
	for i := 0; i < glob.PlayerListMax; i++ {
		slist = append(slist, glob.PlayerList[i])
	}
	glob.PlayerListLock.RUnlock()

	buf := ""

	if argnum < 1 {
		fact.CMS(m.ChannelID, "**Arguments:** <option>\n\n```options:\nrecent (recently online)\nnew (by time joined)\nregistered (recently registered)\n<factorio/discord name search>```")
		return
	} else if strings.ToLower(args[0]) == "recent" {
		buf = "Recently online:\n"

		sort.Sort(ByLastSeen(slist))

		buf = buf + fmt.Sprintf("`%20s : %20s : %12s : %12s : %7s`\n", "Factorio Name", "Discord Name", "Last Seen", "Joined", "Level")

		tnow := time.Now()
		tnow = tnow.Round(time.Second)

		count := 0
		for _, p := range slist {
			if p.LastSeen > 0 && count < maxresults {
				lseen := ""
				if p.LastSeen == 0 {
					lseen = constants.Unknown
				} else {
					ltime := time.Unix(p.LastSeen, 0)
					lseen = tnow.Sub(ltime.Round(time.Second)).String()
				}

				joined := ""
				if p.Creation == 0 {
					joined = constants.Unknown
				} else {
					jtime := time.Unix(p.Creation, 0)
					joined = tnow.Sub(jtime.Round(time.Second)).String()
				}
				buf = buf + fmt.Sprintf("`%20s : %20s : %12s : %12s : %7s`\n", sclean.TruncateString(p.Name, 20), sclean.TruncateString(disc.GetNameFromID(p.ID, false), 20), lseen, joined, levelToString(p.Level))
				count++
			}
		}

	} else if strings.ToLower(args[0]) == "new" {
		buf = "Recently joined:\n"

		sort.Sort(ByNew(slist))

		buf = buf + fmt.Sprintf("`%20s : %20s : %12s : %12s : %7s`\n", "Factorio Name", "Discord Name", "Last Seen", "Joined", "Level")

		tnow := time.Now()
		tnow = tnow.Round(time.Second)

		count := 0
		for _, p := range slist {
			if p.LastSeen > 0 && count < maxresults {
				lseen := ""
				if p.LastSeen == 0 {
					lseen = constants.Unknown
				} else {
					ltime := time.Unix(p.LastSeen, 0)
					lseen = tnow.Sub(ltime.Round(time.Second)).String()
				}

				joined := ""
				if p.Creation == 0 {
					joined = constants.Unknown
				} else {
					jtime := time.Unix(p.Creation, 0)
					joined = tnow.Sub(jtime.Round(time.Second)).String()
				}
				buf = buf + fmt.Sprintf("`%20s : %20s : %12s : %12s : %7s`\n", sclean.TruncateString(p.Name, 20), sclean.TruncateString(disc.GetNameFromID(p.ID, false), 20), lseen, joined, levelToString(p.Level))
				count++
			}
		}

	} else if strings.ToLower(args[0]) == "registered" {
		buf = "Recently joined and registered:\n"

		sort.Sort(ByNew(slist))

		buf = buf + fmt.Sprintf("`%20s : %20s : %12s : %12s : %7s`\n", "Factorio Name", "Discord Name", "Last Seen", "Joined", "Level")

		tnow := time.Now()
		tnow = tnow.Round(time.Second)

		count := 0
		for _, p := range slist {
			if p.ID != "" && p.Name != "" {
				if p.LastSeen > 0 && count < maxresults {
					lseen := ""
					if p.LastSeen == 0 {
						lseen = constants.Unknown
					} else {
						ltime := time.Unix(p.LastSeen, 0)
						lseen = tnow.Sub(ltime.Round(time.Second)).String()
					}

					joined := ""
					if p.Creation == 0 {
						joined = constants.Unknown
					} else {
						jtime := time.Unix(p.Creation, 0)
						joined = tnow.Sub(jtime.Round(time.Second)).String()
					}
					buf = buf + fmt.Sprintf("`%20s : %20s : %12s : %12s : %7s`\n", sclean.TruncateString(p.Name, 20), sclean.TruncateString(disc.GetNameFromID(p.ID, false), 20), lseen, joined, levelToString(p.Level))
					count++
				}
			}
		}

	} else {
		tnow := time.Now()
		tnow = tnow.Round(time.Second)

		count := 0
		for _, p := range slist {
			if count > maxresults {
				break
			}
			if strings.Contains(strings.ToLower(p.Name), strings.ToLower(args[0])) || strings.Contains(strings.ToLower(disc.GetNameFromID(p.ID, false)), strings.ToLower(args[0])) {

				lseen := ""
				if p.LastSeen == 0 {
					lseen = constants.Unknown
				} else {
					ltime := time.Unix(p.LastSeen, 0)
					lseen = tnow.Sub(ltime.Round(time.Second)).String()
				}

				joined := ""
				if p.Creation == 0 {
					joined = constants.Unknown
				} else {
					jtime := time.Unix(p.Creation, 0)
					joined = tnow.Sub(jtime.Round(time.Second)).String()
				}
				buf = buf + fmt.Sprintf("`%20s : %20s : %12s : %12s : %7s`\n", sclean.TruncateString(p.Name, 20), sclean.TruncateString(disc.GetNameFromID(p.ID, false), 20), lseen, joined, levelToString(p.Level))
			}
		}
		if buf == "" {
			buf = "No results."
		}
	}

	fact.CMS(m.ChannelID, buf)
}

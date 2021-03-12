package glob

import (
	"io"
	"os"
	"time"

	"../constants"
	"github.com/bwmarrin/discordgo"
	"github.com/sasha-s/go-deadlock"
)

var Uptime time.Time

var Guild *discordgo.Guild
var Guildname = constants.Unknown
var GuildLock deadlock.RWMutex

type CMSBuf struct {
	Added   time.Time
	Channel string
	Text    string
}

var CMSBuffer []CMSBuf
var CMSBufferLock deadlock.Mutex

var Pipe io.WriteCloser
var PipeLock deadlock.Mutex

var DS *discordgo.Session

var FactorioVersion = constants.Unknown
var NewVersion = constants.Unknown
var NewPatchName = constants.Unknown

var OldChanName = constants.Unknown
var NewChanName = constants.Unknown
var GameLogName = ""
var BotLogName = ""
var GameLogDesc *os.File
var BotLogDesc *os.File

var GameMapName = ""
var GameMapPath = ""
var GameMapLock deadlock.Mutex

var ModLoadMessage *discordgo.Message
var ModLoadString = constants.Unknown
var ModLoadLock deadlock.RWMutex

var LastGametime = ""
var PausedTicks = 0
var PausedTicksLock deadlock.RWMutex

var Gametime = constants.Unknown
var GametimeString = constants.Unknown
var GametimeLock deadlock.RWMutex

var SaveTimer time.Time
var SaveTimerLock deadlock.RWMutex

var FactIsRunning = false
var FactIsRunningLock deadlock.RWMutex

var FactAutoStart = false
var FactAutoStartLock deadlock.RWMutex

var QueueReload = false
var QueueReloadLock deadlock.RWMutex

var DoRebootBot = false
var DoRebootBotLock deadlock.RWMutex

var RecordPlayers = 0
var NumPlayers = 0
var NumPlayersLock deadlock.RWMutex

var RelaunchThrottle = 0
var RelaunchThrottleLock deadlock.RWMutex

type PList struct {
	Name     string
	Level    int
	ID       string
	Creation int64
	LastSeen int64
}

var PlayerListMax = 0
var PlayerList [constants.MaxPlayers + 1]PList
var PlayerListLock deadlock.RWMutex
var PlayerListWriteLock deadlock.Mutex

var PasswordList [constants.MaxPasswords + 1]string
var PasswordID [constants.MaxPasswords + 1]string
var PasswordTime [constants.MaxPasswords + 1]int64
var PasswordMax = 0

var NumLogins = 0
var NumLoginsLock deadlock.RWMutex
var UpdateChannelLock deadlock.Mutex

var FactorioBooted = false
var FactorioBootedLock deadlock.RWMutex

var PlayerListUpdated = false
var PlayerListUpdatedLock deadlock.Mutex

var PlayerListDirty = false
var PlayerListDirtyLock deadlock.Mutex

var PlayerListSeenDirty = false
var PlayerListSeenDirtyLock deadlock.Mutex

var RecordPlayersWriteLock deadlock.Mutex
var RecordPlayersLock deadlock.RWMutex

var PasswordListLock deadlock.RWMutex

var NoResponseCount = 0
var NoResponseCountLock deadlock.RWMutex

var MaxMapTypes = 0
var LastMapSeed uint64 = 0
var LastMapCode = ""
var FactorioLaunchLock deadlock.Mutex

var UpdateFactorioLock deadlock.Mutex
var DoUpdateFactorio = false
var DoUpdateFactorioLock deadlock.Mutex

var ManMinutes = 0
var ManMinutesLock deadlock.Mutex

var LastColor = 0

var ConnectPauseLock deadlock.Mutex
var ConnectPauseTimer int64 = 0
var ConnectPauseCount = 0

var LastTotalStat = ""
var LastMemberStat = ""
var LastRegularStat = ""

var FactQuitTimerLock deadlock.Mutex
var FactQuitTimer time.Time

var UpdateWarnCounterLock deadlock.Mutex
var UpdateWarnCounter = 0
var UpdateGraceMinutes = 5

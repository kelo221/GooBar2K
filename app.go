package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ebitengine/oto/v3"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"sync"
	"time"
)

// TODO  move to another file later

type AudioPlayer struct {
	player     *oto.Player
	context    *oto.Context
	done       chan bool
	pause      chan bool
	wg         sync.WaitGroup
	audioPath  string
	sampleRate int
	channels   int
}

func NewAudioPlayer(audioPath string) *AudioPlayer {
	op := &oto.NewContextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}

	// Create context once during initialization
	ctx, readyChan, err := oto.NewContext(op)
	if err != nil {
		fmt.Println("Error creating context:", err)
		return nil
	}
	<-readyChan

	return &AudioPlayer{
		audioPath:  audioPath,
		context:    ctx,
		done:       make(chan bool),
		pause:      make(chan bool),
		sampleRate: 44100,
		channels:   2,
	}
}

func (ap *AudioPlayer) Pause() {
	ap.pause <- true
}

func (ap *AudioPlayer) Stop() {
	ap.done <- true
	ap.wg.Wait()
	close(ap.done)
	close(ap.pause)
	ap.player.Close()
}

func (ap *AudioPlayer) Start() error {
	buf := bytes.NewBuffer(nil)

	err := ffmpeg.Input(ap.audioPath).
		Output("pipe:",
			ffmpeg.KwArgs{
				"f":      "s16le",
				"acodec": "pcm_s16le",
				"ar":     ap.sampleRate,
				"ac":     ap.channels,
			}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return err
	}

	// Use existing context instead of creating new one
	if ap.player != nil {
		ap.player.Close()
	}
	ap.player = ap.context.NewPlayer(bytes.NewReader(buf.Bytes()))

	ap.wg.Add(1)
	go func() {
		defer ap.wg.Done()
		ap.player.Play()
		for {
			select {
			case <-ap.pause:
				ap.player.Pause()
				time.Sleep(3 * time.Second)
				ap.player.Play()
			case <-ap.done:
				return
			default:
				if !ap.player.IsPlaying() {
					return
				}
				time.Sleep(time.Millisecond)
			}
		}
	}()

	return nil
}

//

// App struct
type App struct {
	ctx    context.Context
	player *AudioPlayer
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		player: NewAudioPlayer("tmp/test_sound.wav"),
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	fmt.Println("Go backend Greet function called with name: ", name)
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) LoadFileRequest(path string) {
	if a.player != nil {
		a.player.Stop()
	}
	a.player = NewAudioPlayer(path)
}

func (a *App) PlayRequest() {
	if err := a.player.Start(); err != nil {
		fmt.Println("Error playing audio:", err)
	}
}

func (a *App) PauseRequest() {
	a.player.Pause()
}

func (a *App) StopRequest() {
	a.player.Stop()
}

func (a *App) PlayNextReques() {
	fmt.Println("PlayNext function called from Go backend")
}

func (a *App) PlayPreviousReques() {
	fmt.Println("PlayPrevious function called from Go backend")
}

func (a *App) PlayRandomReques() {
	fmt.Println("PlayRandom function called from Go backend")
}

func (a *App) SeekToReques(time int) {
	fmt.Println("SeekTo function called from Go backend with time: ", time)
}

func (a *App) SetVolumeReques(volume int) {
	fmt.Println("SetVolume function called from Go backend with volume: ", volume)
}

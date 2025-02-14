package main

//
//import (
//	"bytes"
//	"github.com/ebitengine/oto/v3"
//	ffmpeg "github.com/u2takey/ffmpeg-go"
//	"os"
//	"sync"
//	"time"
//)
//
//type AudioPlayer struct {
//	player     *oto.Player
//	done       chan bool
//	pause      chan bool
//	wg         sync.WaitGroup
//	audioPath  string
//	sampleRate int
//	channels   int
//}
//
//func NewAudioPlayer(audioPath string) *AudioPlayer {
//	return &AudioPlayer{
//		audioPath:  audioPath,
//		done:       make(chan bool),
//		pause:      make(chan bool),
//		sampleRate: 44100,
//		channels:   2,
//	}
//}
//
//func (ap *AudioPlayer) Pause() {
//	ap.pause <- true
//}
//
//func (ap *AudioPlayer) Stop() {
//	ap.done <- true
//	ap.wg.Wait()
//	close(ap.done)
//	close(ap.pause)
//	ap.player.Close()
//}
//
//func (ap *AudioPlayer) Start() error {
//	buf := bytes.NewBuffer(nil)
//
//	err := ffmpeg.Input(ap.audioPath).
//		Output("pipe:",
//			ffmpeg.KwArgs{
//				"f":      "s16le",
//				"acodec": "pcm_s16le",
//				"ar":     ap.sampleRate,
//				"ac":     ap.channels,
//			}).
//		WithOutput(buf, os.Stdout).
//		Run()
//	if err != nil {
//		return err
//	}
//
//	op := &oto.NewContextOptions{
//		SampleRate:   ap.sampleRate,
//		ChannelCount: ap.channels,
//		Format:       oto.FormatSignedInt16LE,
//	}
//
//	otoCtx, readyChan, err := oto.NewContext(op)
//	if err != nil {
//		return err
//	}
//	<-readyChan
//
//	ap.player = otoCtx.NewPlayer(bytes.NewReader(buf.Bytes()))
//
//	ap.wg.Add(1)
//	go func() {
//		defer ap.wg.Done()
//		ap.player.Play()
//		for {
//			select {
//			case <-ap.pause:
//				ap.player.Pause()
//				time.Sleep(3 * time.Second)
//				ap.player.Play()
//			case <-ap.done:
//				return
//			default:
//				if !ap.player.IsPlaying() {
//					return
//				}
//				time.Sleep(time.Millisecond)
//			}
//		}
//	}()
//
//	return nil
//}

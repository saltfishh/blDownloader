package lib

import (
	"os/exec"
)

func MergeAudioVideo(videoPath, audioPath, outputPath string) error {
	//log.Printf("Merge audio video %s %s to %s", videoPath, audioPath, outputPath)
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-i", audioPath, "-c:v", "copy", "-c:a", "copy", outputPath)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

package hls

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateSegments(inputPath, outputDir string) error {
    // Absolute yo‘llarni olish
    absInputPath, err := filepath.Abs(inputPath)
    if err != nil {
        return fmt.Errorf("error getting absolute path for input: %w", err)
    }

    absOutputDir, err := filepath.Abs(outputDir)
    if err != nil {
        return fmt.Errorf("error getting absolute path for output: %w", err)
    }

    // Input katalogi mavjudligini tekshirish
    if _, err := os.Stat(absInputPath); os.IsNotExist(err) {
        return fmt.Errorf("input file does not exist: %s", absInputPath)
    }

    // Output katalogi mavjudligini tekshirish
    if _, err := os.Stat(absOutputDir); os.IsNotExist(err) {
        return fmt.Errorf("output directory does not exist: %s", absOutputDir)
    }

    // FFmpeg-ni chaqirish
    cmd := exec.Command("ffmpeg", "-i", absInputPath,
        "-codec", "copy",
        "-start_number", "0",
        "-hls_time", "10",
        "-hls_list_size", "0",
        "-f", "hls",
        filepath.Join(absOutputDir, "playlist.m3u8"))

    // Buyruq natijalarini olish
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Printf("FFmpeg error: %v\nOutput: %s", err, out)
        return fmt.Errorf("FFmpeg execution failed: %w", err)
    }

    log.Println("Video segmentation completed successfully.")
    return nil
}

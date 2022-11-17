package main

import (
	"image"
	"os"
	// "fmt"
	"testing"
)

func TestLoadImage(t *testing.T) {
	got, _ := loadImage("")

	if got != nil {
		t.Errorf("Failed to load empty image.")
	}

	got, err := loadImage("test.png")
	if err != nil {
		t.Errorf("Failed to load test image")
	}

}
func TestConrtastCalc(t *testing.T) {
	f, err := os.Open("test.png")
	if err != nil {
		t.Errorf("Failed to open test image")
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("Failed to decode image")
	}

	got := contrastCalc(img.At(1, 1))
	want := 155

	if got < want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func TestInterpolatePixels(t *testing.T) {
	f, err := os.Open("test.png")
	if err != nil {
		t.Errorf("Failed to open test image")
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("Failed to decode testimage")
	}

	got := interpolatePixels(img, 1, 1, 10, 5)
	want := 150

	if got <= want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func TestcalcScale(t *testing.T) {
	got, got2 := calcScale(1, 1)
	want, want2 := 5, 10

	if got != want && got2 != want2 {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

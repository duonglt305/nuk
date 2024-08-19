package services

import (
	"fmt"
	"math/rand"
	"time"

	"duonglt.net/pkg/cache"
)

type OtpService struct {
	cache cache.ICache
}

// NewOtpService creates a new OTP service
func NewOtpService(cache cache.ICache) OtpService {
	return OtpService{cache}
}

// Generate generates an OTP
func (s OtpService) Generate(key string, seconds int) string {
	// Create a new random source
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate unique OTP
	otp := fmt.Sprintf("%06d", src.Intn(999999))
	// Save OTP to cache
	s.cache.Set(key, otp, time.Second*time.Duration(seconds))

	return otp
}

// Verify verifies an OTP
func (s OtpService) Verify(key, otp string) error {
	// Get OTP from cache
	cachedOtp, err := s.cache.Get(key)
	if err != nil {
		return err
	}
	// Compare OTP
	if cachedOtp != otp {
		return fmt.Errorf("invalid OTP")
	}
	// Delete OTP from cache
	s.cache.Delete(key)

	return nil
}

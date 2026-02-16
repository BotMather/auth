package services

import (
	"context"
	"time"

	"github.com/JscorpTech/auth/internal/modules/auth"
	"go.uber.org/zap"
)

func OtpClean(ctx context.Context, logger *zap.Logger, repo auth.AuthRepository) {
	logger.Info("Otp cleaner started")
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			logger.Info("Otp cleaner to'xtatildi")
			return
		case <-ticker.C:
			otps, err := repo.GetOldOtps(ctx)
			if err != nil {
				logger.Error("Otp clean error", zap.Error(err))
			}
			for _, otp := range otps {
				repo.DeleteOtp(ctx, &otp)
				logger.Info("Otp expired", zap.String("code", otp.Code), zap.String("phone", otp.Phone))
			}
		}
	}
}

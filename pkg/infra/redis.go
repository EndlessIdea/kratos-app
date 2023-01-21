// Copyright (c) 2012-2022 Grabtaxi Holdings PTE LTD (GRAB), All Rights Reserved. NOTICE: All information contained herein
// is, and remains the property of GRAB. The intellectual and technical concepts contained herein are confidential, proprietary
// and controlled by GRAB and may be covered by patents, patents in process, and are protected by trade secret or copyright law.
//
// You are strictly forbidden to copy, download, store (in any medium), transmit, disseminate, adapt or change this material
// in any way unless prior written permission is obtained from GRAB. Access to the source code contained herein is hereby
// forbidden to anyone except current GRAB employees or contractors with binding Confidentiality and Non-disclosure agreements
// explicitly covering such access.
//
// The copyright notice above does not evidence any actual or intended publication or disclosure of this source code,
// which includes information that is confidential and/or proprietary, and is a trade secret, of GRAB.
//
// ANY REPRODUCTION, MODIFICATION, DISTRIBUTION, PUBLIC PERFORMANCE, OR PUBLIC DISPLAY OF OR THROUGH USE OF THIS SOURCE
// CODE WITHOUT THE EXPRESS WRITTEN CONSENT OF GRAB IS STRICTLY PROHIBITED, AND IN VIOLATION OF APPLICABLE LAWS AND
// INTERNATIONAL TREATIES. THE RECEIPT OR POSSESSION OF THIS SOURCE CODE AND/OR RELATED INFORMATION DOES NOT CONVEY
// OR IMPLY ANY RIGHTS TO REPRODUCE, DISCLOSE OR DISTRIBUTE ITS CONTENTS, OR TO MANUFACTURE, USE, OR SELL ANYTHING
// THAT IT MAY DESCRIBE, IN WHOLE OR IN PART.

package infra

import (
	"crypto/tls"

	"github.com/EndlessIdea/kratos-app/pkg/infra/config"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// NewRedisClient init redis client
func NewRedisClient(cfg *config.Redis) (*redis.Client, error) {
	if cfg == nil {
		return nil, errors.Errorf("redis config is empty")
	}

	redisOpts := redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DialTimeout:  cfg.DialTimeout.AsDuration(),
		WriteTimeout: cfg.WriteTimeout.AsDuration(),
		ReadTimeout:  cfg.ReadTimeout.AsDuration(),
		PoolTimeout:  cfg.PoolTimeout.AsDuration(),
		PoolSize:     int(cfg.PoolSize),
	}
	if cfg.Password != "" {
		redisOpts.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return redis.NewClient(&redisOpts), nil
}

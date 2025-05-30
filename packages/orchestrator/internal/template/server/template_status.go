package server

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	template_manager "github.com/e2b-dev/infra/packages/shared/pkg/grpc/template-manager"
)

func (s *ServerStore) TemplateBuildStatus(ctx context.Context, in *template_manager.TemplateStatusRequest) (*template_manager.TemplateBuildStatusResponse, error) {
	ctx, ctxSpan := s.tracer.Start(ctx, "template-build-status-request")
	defer ctxSpan.End()

	logger := s.logger.With(zap.String("buildID", in.BuildID), zap.String("envID", in.TemplateID))
	logger.Info("Template build status request")

	buildInfo, err := s.buildCache.Get(in.BuildID)
	if err != nil {
		return nil, errors.Wrap(err, "error while getting build info, maybe already expired")
	}

	return &template_manager.TemplateBuildStatusResponse{
		Status:   buildInfo.GetStatus(),
		Metadata: buildInfo.GetMetadata(),
	}, nil
}

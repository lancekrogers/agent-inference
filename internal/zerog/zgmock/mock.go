// Package zgmock provides mock implementations of all 0G subsystem interfaces
// for demo mode. No real 0G chain connections are made.
package zgmock

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/lancekrogers/agent-inference/internal/zerog/compute"
	"github.com/lancekrogers/agent-inference/internal/zerog/da"
	"github.com/lancekrogers/agent-inference/internal/zerog/inft"
	"github.com/lancekrogers/agent-inference/internal/zerog/storage"
)

// ComputeBroker returns simulated inference results.
type ComputeBroker struct {
	jobCounter int
}

func NewComputeBroker() compute.ComputeBroker { return &ComputeBroker{} }

func (m *ComputeBroker) SubmitJob(_ context.Context, _ compute.JobRequest) (string, error) {
	m.jobCounter++
	return fmt.Sprintf("mock-job-%d", m.jobCounter), nil
}

func (m *ComputeBroker) GetResult(_ context.Context, jobID string) (*compute.JobResult, error) {
	return &compute.JobResult{
		JobID:      jobID,
		Status:     "completed",
		Output:     `{"result": "mock inference output"}`,
		ModelID:    "llama-3-8b",
		TokensUsed: 80 + rand.Intn(400),
		Duration:   time.Duration(40+rand.Intn(180)) * time.Millisecond,
	}, nil
}

func (m *ComputeBroker) ListModels(_ context.Context) ([]compute.Model, error) {
	return []compute.Model{
		{ID: "model-1", Name: "llama-3-8b", Provider: "0g-compute"},
		{ID: "model-2", Name: "mistral-7b", Provider: "0g-compute"},
	}, nil
}

// StorageClient returns simulated storage operations.
type StorageClient struct {
	uploadCounter int
}

func NewStorageClient() storage.StorageClient { return &StorageClient{} }

func (m *StorageClient) Upload(_ context.Context, _ []byte, _ storage.Metadata) (string, error) {
	m.uploadCounter++
	return fmt.Sprintf("mock-content-%d", m.uploadCounter), nil
}

func (m *StorageClient) Download(_ context.Context, _ string) ([]byte, error) {
	return []byte(`{"mock": true}`), nil
}

func (m *StorageClient) List(_ context.Context, _ string) ([]storage.Metadata, error) {
	return nil, nil
}

// INFTMinter returns simulated iNFT operations.
type INFTMinter struct{}

func NewINFTMinter() inft.INFTMinter { return &INFTMinter{} }

func (m *INFTMinter) Mint(_ context.Context, _ inft.MintRequest) (string, error) {
	return "mock-inft-001", nil
}

func (m *INFTMinter) UpdateMetadata(_ context.Context, _ string, _ inft.EncryptedMeta) error {
	return nil
}

func (m *INFTMinter) GetStatus(_ context.Context, tokenID string) (*inft.INFTStatus, error) {
	return &inft.INFTStatus{
		TokenID:      tokenID,
		Owner:        "0x0000000000000000000000000000000000000000",
		MintedAt:     time.Now().Add(-24 * time.Hour),
		MetadataHash: "0xmockhash",
		ChainID:      16602,
	}, nil
}

// AuditPublisher returns simulated DA operations.
type AuditPublisher struct {
	pubCounter int
}

func NewAuditPublisher() da.AuditPublisher { return &AuditPublisher{} }

func (m *AuditPublisher) Publish(_ context.Context, _ da.AuditEvent) (string, error) {
	m.pubCounter++
	return fmt.Sprintf("mock-audit-%d", m.pubCounter), nil
}

func (m *AuditPublisher) Verify(_ context.Context, _ string) (bool, error) {
	return true, nil
}

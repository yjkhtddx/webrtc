// +build !js

package webrtc

import (
	"errors"
	"io"

	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3/pkg/media"
)

var (
	ErrUnsupportedCodec = errors.New("unsupported codec")
	ErrUnbindFailed     = errors.New("failed to unbind Track from PeerConnection")
)

// trackBinding is a single bind for a Track
// Bind can be called multiple times, this stores the
// result for a single bind call so that it can be used when writing
type trackBinding struct {
	ssrc        SSRC
	payloadType PayloadType
	writeStream TrackLocalWriter
}

// TrackLocalStaticRTP  is a track that has a pre-set codec
type TrackLocalStaticRTP struct {
	bindings     []trackBinding
	codec        RTPCodecCapability
	id, streamID string
}

// NewStaticTrackLocal returns a TrackLocalRTP with a pre-set codec.
func NewTrackLocalStaticRTP(c RTPCodecCapability, id, streamID string) (*TrackLocalStaticRTP, error) {
	return &TrackLocalStaticRTP{
		codec:    c,
		bindings: []trackBinding{},
		id:       id,
		streamID: streamID,
	}, nil
}

// Bind is called by the PeerConnection after negotiation is complete
// This asserts that the code requested is supported by the remote peer.
// If so it setups all the state (SSRC and PayloadType) to have a call
func (s *TrackLocalStaticRTP) Bind(t TrackLocalContext) error {
	for _, codec := range t.CodecParameters() {
		if codec.MimeType == s.codec.MimeType {
			s.bindings = append(s.bindings, trackBinding{
				ssrc:        t.SSRC(),
				payloadType: codec.PayloadType,
				writeStream: t.WriteStream(),
			})
			return nil
		}
	}
	return ErrUnsupportedCodec
}

func (s *TrackLocalStaticRTP) Unbind(t TrackLocalContext) error {
	for i := range s.bindings {
		if s.bindings[i].writeStream == t.WriteStream() {
			s.bindings[i] = s.bindings[len(s.bindings)-1]
			s.bindings = s.bindings[:len(s.bindings)-1]
			return nil
		}
	}

	return ErrUnbindFailed
}

func (s *TrackLocalStaticRTP) ID() string         { return s.id }
func (s *TrackLocalStaticRTP) StreamID() string   { return s.streamID }
func (s *TrackLocalStaticRTP) Kind() RTPCodecType { return RTPCodecTypeVideo }

// Loop each binding and set the proper SSRC/PayloadType before writing
func (s *TrackLocalStaticRTP) WriteRTP(p *rtp.Packet) error      { return io.EOF }
func (s *TrackLocalStaticRTP) Write(b []byte) (n int, err error) { return }

type TrackLocalStaticSample struct {
	packetizer interface{}
	rtpTrack   *TrackLocalStaticRTP
}

func NewTrackLocalStaticSample(c RTPCodecCapability, id, streamID string) (*TrackLocalStaticSample, error) {
	rtpTrack, err := NewTrackLocalStaticRTP(c, id, streamID)
	if err != nil {
		return nil, err
	}

	return &TrackLocalStaticSample{
		packetizer: nil,
		rtpTrack:   rtpTrack,
	}, nil
}

func (s *TrackLocalStaticSample) ID() string         { return s.rtpTrack.ID() }
func (s *TrackLocalStaticSample) StreamID() string   { return s.rtpTrack.StreamID() }
func (s *TrackLocalStaticSample) Kind() RTPCodecType { return s.rtpTrack.Kind() }

// Call rtpTrack.Bind + setup packetizer
func (s *TrackLocalStaticSample) Bind(t TrackLocalContext) error   { return io.EOF }
func (s *TrackLocalStaticSample) Unbind(t TrackLocalContext) error { return io.EOF }

func (s *TrackLocalStaticSample) WriteSample(samp media.Sample) error { return io.EOF }

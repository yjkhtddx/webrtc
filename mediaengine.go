// +build !js

package webrtc

import (
	"fmt"
	"time"
)

// A MediaEngine defines the codecs supported by a PeerConnection.
// MediaEngines populated using RegisterCodec (and RegisterDefaultCodecs)
// may be set up once and reused, including concurrently,
// as long as no other codecs are added subsequently.
type MediaEngine struct {
	videoCodecs           []RTPCodecParameters
	videoHeaderExtensions []RTPHeaderExtensionCapability

	audioCodecs           []RTPCodecParameters
	audioHeaderExtensions []RTPHeaderExtensionCapability
}

// RegisterDefaultCodecs registers the default codecs supported by Pion WebRTC.
// RegisterDefaultCodecs is not safe for concurrent use.
func (m *MediaEngine) RegisterDefaultCodecs() error {
	// Default Pion Audio Codecs
	for _, codec := range []RTPCodecParameters{
		{
			RTPCodecCapability: RTPCodecCapability{"audio/opus", 48000, 2, "minptime=10;useinbandfec=1", []RTCPFeedback{{"transport-cc", ""}}},
			PayloadType:        111,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"audio/G722", 8000, 1, "", nil},
			PayloadType:        9,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"audio/PCMU", 8000, 1, "", nil},
			PayloadType:        0,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"audio/PCMA", 8000, 1, "", nil},
			PayloadType:        8,
		},
	} {
		if err := m.RegisterCodec(codec, RTPCodecTypeAudio); err != nil {
			return err
		}
	}

	// Default Pion Audio Header Extensions
	for _, extension := range []string{
		"urn:ietf:params:rtp-hdrext:ssrc-audio-level",
		"http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time",
		"http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01",
		"urn:ietf:params:rtp-hdrext:sdes:mid",
		"urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id",
		"urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id",
	} {
		if err := m.RegisterHeaderExtension(RTPHeaderExtensionCapability{extension}, RTPCodecTypeAudio); err != nil {
			return err
		}
	}

	videoRTCPFeedback := []RTCPFeedback{{"goog-remb", ""}, {"transport-cc", ""}, {"ccm", "fir"}, {"nack", ""}, {"nack", "pli"}}

	for _, codec := range []RTPCodecParameters{
		{
			RTPCodecCapability: RTPCodecCapability{"video/VP8", 90000, 1, "", videoRTCPFeedback},
			PayloadType:        96,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=96", nil},
			PayloadType:        97,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/VP9", 90000, 1, "profile-id=0", videoRTCPFeedback},
			PayloadType:        98,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=98", nil},
			PayloadType:        99,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/VP9", 90000, 1, "profile-id=1", videoRTCPFeedback},
			PayloadType:        100,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=100", nil},
			PayloadType:        101,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/H264", 90000, 1, "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42001f", videoRTCPFeedback},
			PayloadType:        102,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=102", nil},
			PayloadType:        121,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/H264", 90000, 1, "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42001f", videoRTCPFeedback},
			PayloadType:        127,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=127", nil},
			PayloadType:        120,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/H264", 90000, 1, "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f", videoRTCPFeedback},
			PayloadType:        125,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=125", nil},
			PayloadType:        107,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/H264", 90000, 1, "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42e01f", videoRTCPFeedback},
			PayloadType:        108,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=108", nil},
			PayloadType:        109,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/H264", 90000, 1, "level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42001f", videoRTCPFeedback},
			PayloadType:        127,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=127", nil},
			PayloadType:        120,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/H264", 90000, 1, "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=640032", videoRTCPFeedback},
			PayloadType:        123,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=123", nil},
			PayloadType:        118,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/red", 90000, 1, "", videoRTCPFeedback},
			PayloadType:        114,
		},
		{
			RTPCodecCapability: RTPCodecCapability{"video/rtx", 90000, 1, "apt=114", nil},
			PayloadType:        115,
		},

		{
			RTPCodecCapability: RTPCodecCapability{"video/ulpfec", 90000, 1, "", nil},
			PayloadType:        116,
		},
	} {
		if err := m.RegisterCodec(codec, RTPCodecTypeAudio); err != nil {
			return err
		}
	}

	// Default Pion Video Header Extensions
	for _, extension := range []string{
		"urn:ietf:params:rtp-hdrext:toffset",
		"http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time",
		"urn:3gpp:video-orientation",
		"http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01",
		"http://www.webrtc.org/experiments/rtp-hdrext/playout-delay",
		"http://www.webrtc.org/experiments/rtp-hdrext/video-content-type",
		"http://www.webrtc.org/experiments/rtp-hdrext/video-timing",
		"http://www.webrtc.org/experiments/rtp-hdrext/color-space",
		"urn:ietf:params:rtp-hdrext:sdes:mid",
		"urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id",
		"urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id",
	} {
		if err := m.RegisterHeaderExtension(RTPHeaderExtensionCapability{extension}, RTPCodecTypeVideo); err != nil {
			return err
		}
	}

	return nil
}

// RegisterCodec adds codec to the MediaEngine
// These are the list of codecs supported by this PeerConnection.
// RegisterCodec is not safe for concurrent use.
func (m *MediaEngine) RegisterCodec(codec RTPCodecParameters, typ RTPCodecType) error {
	codec.statsID = fmt.Sprintf("RTPCodec-%d", time.Now().UnixNano())
	switch typ {
	case RTPCodecTypeAudio:
		m.audioCodecs = append(m.audioCodecs, codec)
	case RTPCodecTypeVideo:
		m.videoCodecs = append(m.videoCodecs, codec)
	default:
		return ErrUnknownType
	}
	return nil
}

func (m *MediaEngine) RegisterHeaderExtension(extension RTPHeaderExtensionCapability, typ RTPCodecType) error {
	switch typ {
	case RTPCodecTypeAudio:
		m.audioHeaderExtensions = append(m.audioHeaderExtensions, extension)
	case RTPCodecTypeVideo:
		m.videoHeaderExtensions = append(m.videoHeaderExtensions, extension)
	default:
		return ErrUnknownType
	}
	return nil
}

func (m *MediaEngine) getCodec(payloadType PayloadType, typ RTPCodecType) (RTPCodecParameters, error) {
	for _, codec := range m.GetCodecsByKind(typ) {
		if codec.PayloadType == payloadType {
			return codec, nil
		}
	}
	return RTPCodecParameters{}, ErrCodecNotFound
}

func (m *MediaEngine) GetCodecsByKind(kind RTPCodecType) []RTPCodecParameters {
	switch kind {
	case RTPCodecTypeAudio:
		return m.audioCodecs
	case RTPCodecTypeVideo:
		return m.videoCodecs
	default:
		return []RTPCodecParameters{}
	}
}

func (m *MediaEngine) collectStats(collector *statsReportCollector) {
	for _, codec := range m.videoCodecs {
		collector.Collecting()
		stats := CodecStats{
			Timestamp:   statsTimestampFrom(time.Now()),
			Type:        StatsTypeCodec,
			ID:          codec.statsID,
			PayloadType: codec.PayloadType,
			MimeType:    codec.MimeType,
			ClockRate:   codec.ClockRate,
			Channels:    uint8(codec.Channels),
			SDPFmtpLine: codec.SDPFmtpLine,
		}

		collector.Collect(stats.ID, stats)
	}
}

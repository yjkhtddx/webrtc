// +build !js

package webrtc

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// pion/webrtc#1078
func TestOpusCase(t *testing.T) {
	pc, err := NewPeerConnection(Configuration{})
	assert.NoError(t, err)

	_, err = pc.AddTransceiverFromKind(RTPCodecTypeAudio)
	assert.NoError(t, err)

	offer, err := pc.CreateOffer(nil)
	assert.NoError(t, err)

	assert.True(t, regexp.MustCompile(`(?m)^a=rtpmap:\d+ opus/48000/2`).MatchString(offer.SDP))
	assert.NoError(t, pc.Close())
}

// // pion/webrtc#1442
// func TestCaseInsensitive(t *testing.T) {
// 	m := MediaEngine{}
// 	m.RegisterDefaultCodecs()
//
// 	testCases := []struct {
// 		nameUpperCase string
// 		nameLowerCase string
// 		clockrate     uint32
// 		fmtp          string
// 	}{
// 		{strings.ToUpper(Opus), strings.ToLower(Opus), 48000, "minptime=10;useinbandfec=1"},
// 		{strings.ToUpper(PCMU), strings.ToLower(PCMU), 8000, ""},
// 		{strings.ToUpper(PCMA), strings.ToLower(PCMA), 8000, ""},
// 		{strings.ToUpper(G722), strings.ToLower(G722), 8000, ""},
// 		{strings.ToUpper(VP8), strings.ToLower(VP8), 90000, ""},
// 		{strings.ToUpper(VP9), strings.ToLower(VP9), 90000, ""},
// 		{strings.ToUpper(H264), strings.ToLower(H264), 90000, "level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42001f"},
// 	}
//
// 	for _, f := range testCases {
// 		upperCase, err := m.getCodecSDP(sdp.Codec{
// 			Name:      f.nameUpperCase,
// 			ClockRate: f.clockrate,
// 			Fmtp:      f.fmtp,
// 		})
// 		assert.NoError(t, err)
//
// 		lowerCase, err := m.getCodecSDP(sdp.Codec{
// 			Name:      f.nameLowerCase,
// 			ClockRate: f.clockrate,
// 			Fmtp:      f.fmtp,
// 		})
// 		assert.NoError(t, err)
//
// 		assert.Equal(t, upperCase, lowerCase)
// 	}
// }

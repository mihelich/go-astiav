package astiav_test

import (
	"testing"

	"github.com/asticode/go-astiav"
	"github.com/stretchr/testify/require"
)

func TestCodecParameters(t *testing.T) {
	_, s1, s2, err := videoInputStreams()
	require.NoError(t, err)

	cp1 := s1.CodecParameters()
	require.Equal(t, int64(441324), cp1.BitRate())
	require.Equal(t, astiav.ChromaLocationLeft, cp1.ChromaLocation())
	require.Equal(t, astiav.CodecIDH264, cp1.CodecID())
	require.Equal(t, astiav.CodecTag(0x31637661), cp1.CodecTag())
	require.Equal(t, astiav.ColorPrimariesUnspecified, cp1.ColorPrimaries())
	require.Equal(t, astiav.ColorRangeUnspecified, cp1.ColorRange())
	require.Equal(t, astiav.ColorSpaceUnspecified, cp1.ColorSpace())
	require.Equal(t, astiav.ColorTransferCharacteristicUnspecified, cp1.ColorTransferCharacteristic())
	require.Equal(t, 180, cp1.Height())
	require.Equal(t, astiav.Level(13), cp1.Level())
	require.Equal(t, astiav.MediaTypeVideo, cp1.MediaType())
	require.Equal(t, astiav.PixelFormatYuv420P, cp1.PixelFormat())
	require.Equal(t, astiav.ProfileH264ConstrainedBaseline, cp1.Profile())
	require.Equal(t, astiav.NewRational(1, 1), cp1.SampleAspectRatio())
	require.Equal(t, 320, cp1.Width())

	cp2 := s2.CodecParameters()
	require.Equal(t, int64(161052), cp2.BitRate())
	require.Equal(t, 2, cp2.Channels())
	require.Equal(t, astiav.ChannelLayoutStereo, cp2.ChannelLayout())
	require.Equal(t, astiav.CodecIDAac, cp2.CodecID())
	require.Equal(t, astiav.CodecTag(0x6134706d), cp2.CodecTag())
	require.Equal(t, 1024, cp2.FrameSize())
	require.Equal(t, astiav.MediaTypeAudio, cp2.MediaType())
	require.Equal(t, astiav.SampleFormatFltp, cp2.SampleFormat())
	require.Equal(t, 48000, cp2.SampleRate())

	cp3 := astiav.AllocCodecParameters()
	require.NotNil(t, cp3)
	defer cp3.Free()
	err = cp2.Copy(cp3)
	require.NoError(t, err)
	require.Equal(t, 2, cp3.Channels())

	cc4 := astiav.AllocCodecContext(nil)
	require.NotNil(t, cc4)
	defer cc4.Free()
	err = cp2.ToCodecContext(cc4)
	require.NoError(t, err)
	require.Equal(t, 2, cc4.Channels())

	cp5 := astiav.AllocCodecParameters()
	require.NotNil(t, cp5)
	defer cp5.Free()
	err = cp5.FromCodecContext(cc4)
	require.NoError(t, err)
	require.Equal(t, 2, cp5.Channels())

	cp6 := astiav.AllocCodecParameters()
	require.NotNil(t, cp6)
	defer cp6.Free()
	cp6.SetCodecTag(astiav.CodecTag(2))
	require.Equal(t, astiav.CodecTag(2), cp6.CodecTag())
}

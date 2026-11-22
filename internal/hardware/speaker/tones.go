package speaker

import (
	"github.com/franzudev/libcountertop/pkg/audio/tone"
	"github.com/franzudev/libcountertop/pkg/audio/upsample"

	"github.com/franzudev/libcountertop/pkg/settings/schema"

	"github.com/franzudev/echolocal/internal/config"
)

func WakeTones() []schema.Chime {
	return []schema.Chime{schema.ChimeNone, schema.ChimeChirp, schema.ChimeDing, schema.ChimeRise}
}

func MuteTones() []schema.Chime {
	return []schema.Chime{
		schema.ChimeNone,
		schema.ChimeMuteSwitch,
		schema.ChimeChirp,
		schema.ChimeDing,
		schema.ChimeRise,
	}
}

func MuteTone(t schema.Chime, muted bool) []tone.Note {
	if t == schema.ChimeMuteSwitch {
		if muted {
			return tone.ToneMute
		}
		return tone.ToneUnmute
	}
	return tone.Wake(t)
}

func Resamplings() []config.Resampling {
	var out []config.Resampling
	for _, k := range upsample.Kinds() {
		out = append(out, config.Resampling(k))
	}
	return out
}

package denormalizer

import "github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/semantic"

type Base struct {
	// TODO refactor into denormalizer somehow.
	Name             *semantic.String // denormalized always disabled param
	RecycleCandidate *semantic.String // denormalized always disabled param
}

// base_to_add.Name = (&semantic.String{}).Map(base, KEY_NAME, semantic.TypeComment, inireader.OPTIONAL_p)
// base_to_add.RecycleCandidate = (&semantic.String{}).Map(base, KEY_RECYCLE, semantic.TypeComment, inireader.OPTIONAL_p)

type BaseGood struct {
	// TODO refactor into denormalizer somehow.
	Name             *semantic.String // denormalized always disabled param
	RecycleCandidate *semantic.String // denormalized always disabled param
}

// base_to_add.Name = (&semantic.String{}).Map(section, KEY_NAME, semantic.TypeComment, inireader.OPTIONAL_p)
// base_to_add.RecycleCandidate = (&semantic.String{}).Map(section, KEY_RECYCLE, semantic.TypeComment, inireader.OPTIONAL_p)

package golang

import zerologger "github.com/rs/zerolog/log"

var log = zerologger.With().Str("module", "golang").Logger()

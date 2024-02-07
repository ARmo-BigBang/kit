package faker

import (
	"time"

	"github.com/ARmo-BigBang/kit/dtp"
)

func SQLNow() dtp.NullTime {
	return dtp.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}

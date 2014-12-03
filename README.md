Google Speak
============

Go implementation for transform text to speech using Google Translate.

Usage:
```go
package main

import (
  "github.com/yanzay/googlespeak"
)

func main() {
  googlespeak.Say("hi") // English by default
  googlespeak.Say("привет", "ru")
}
```

Supported languages:
 * af (Afrikaans)
 * ar (Arabic)
 * ca (Catalan)
 * cs (Czech)
 * cy (Welsh)
 * da (Danish)
 * de (German)
 * el (Greek)
 * en (English)
 * es (Spanish)
 * fi (Finnish)
 * fr (French)
 * hi (Hindi)
 * hr (Croatian)
 * ht (Haitian)
 * hu (Hungarian)
 * hy (Armenian)
 * id (Indonesian)
 * is (Icelandic)
 * it (Italian)
 * ja (Japanese)
 * ko (Korean)
 * la (Latin)
 * lv (Latvian)
 * mk (Macedonian)
 * nl (Dutch)
 * no (Norwegian)
 * pl (Polish)
 * pt (Portuguese)
 * ro (Romanian)
 * ru (Russian)
 * sk (Slovak)
 * sq (Albanian)
 * sr (Serbian)
 * sv (Swedish)
 * sw (Swahili)
 * ta (Tamil)
 * tr (Turkish)
 * vi (Vietnamese)
 * zh (Chinese)

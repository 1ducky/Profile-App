package article

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ===================== Helpers =====================

// // loadOwned mengambil artikel dan memastikan authorID adalah pemiliknya.
// func (uc *usecase) loadOwned(ctx context.Context, id, authorID string) (*ArticleEntity, error) {
// 	a, err := uc.repo.FindByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if a.AuthorID != authorID {
// 		// Alternatif: kembalikan ErrNotFound agar keberadaan artikel orang lain tidak bocor.
// 		return nil, ErrForbidden
// 	}
// 	return a, nil
// }

func (uc *usecase) isPublished(a *ArticleEntity) bool {
	return !a.PublishAt.IsZero() && !a.PublishAt.After(time.Now())
}

func invalid(msg string) error {
	return fmt.Errorf("%w: %s", ErrInvalidInput, msg)
}

func articleURL(slug string) string {
	return "/articles/" + slug
}

var nonSlugChars = regexp.MustCompile(`[^a-z0-9]+`)

// slugify: "Belajar Go 101!" -> "belajar-go-101"
func slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = nonSlugChars.ReplaceAllString(s, "-")
	return strings.Trim(s, "-")
}

func nonNil(s []string) []string {
	if s == nil {
		return []string{}
	}
	return s
}

func timePtr(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func IntParam(raw, name string) (int, error) {
	if raw == "" {
		return 0, nil // 0 = pakai default di Normalize()
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n < 0 {
		return 0, ErrBadRequest
	}
	return n, nil
}

var filekeyRe = regexp.MustCompile(`assets:([^\s)]+)`)

func ExtractFileKeys(content string) []string {
	// Cocokkan [/filekey/<nama>.png] dan ambil bagian di dalam kurung siku

	matches := filekeyRe.FindAllStringSubmatch(content, -1)

	result := make([]string, 0, len(matches))
	for _, m := range matches {
		result = append(result, m[1]) // m[1] = isi capture group
		log.Print("ExtractFileKeys: ", m)
	}
	log.Print("ExtractFileKeys: ", result)
	return result
}

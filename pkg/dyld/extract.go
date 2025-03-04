package dyld

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/apex/log"
	"github.com/blacktop/ipsw/internal/utils"
	"github.com/blacktop/ipsw/pkg/info"
	"github.com/blacktop/ipsw/pkg/ota/ridiff"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func GetDscPathsInMount(mountPoint string) ([]string, error) {
	var matches []string

	re := regexp.MustCompile(fmt.Sprintf("%s%c%s", mountPoint, filepath.Separator, CacheRegex))

	if err := filepath.Walk(mountPoint, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			utils.Indent(log.Warn, 3)(fmt.Sprintf("failed to walk %s: %v", path, err))
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if re.MatchString(path) {
			matches = append(matches, path)
		}
		return nil
	}); err != nil {
		return nil, err // FIXME: this will never error
	}

	return matches, nil
}

func ExtractFromDMG(i *info.Info, dmgPath, destPath string, arches []string) ([]string, error) {

	utils.Indent(log.Info, 2)(fmt.Sprintf("Mounting DMG %s", dmgPath))
	var alreadyMounted bool
	mountPoint, alreadyMounted, err := utils.MountFS(dmgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to IPSW FS dmg: %v", err)
	}
	if alreadyMounted {
		utils.Indent(log.Debug, 3)(fmt.Sprintf("%s already mounted", dmgPath))
	} else {
		defer func() {
			utils.Indent(log.Debug, 2)(fmt.Sprintf("Unmounting %s", dmgPath))
			if err := utils.Retry(3, 2*time.Second, func() error {
				return utils.Unmount(mountPoint, false)
			}); err != nil {
				log.Errorf("failed to unmount DMG %s at %s: %v", dmgPath, mountPoint, err)
			}
		}()
	}

	if runtime.GOOS == "darwin" {
		if err := os.MkdirAll(destPath, 0750); err != nil {
			return nil, fmt.Errorf("failed to create destination directory %s: %v", destPath, err)
		}
	} else {
		if _, ok := os.LookupEnv("IPSW_IN_DOCKER"); ok {
			if err := os.MkdirAll(filepath.Join("/data", destPath), 0750); err != nil {
				return nil, fmt.Errorf("failed to create destination directory %s: %v", destPath, err)
			}
		} else {
			if err := os.MkdirAll(destPath, 0750); err != nil {
				return nil, fmt.Errorf("failed to create destination directory %s: %v", destPath, err)
			}
		}
	}

	matches, err := GetDscPathsInMount(mountPoint)
	if err != nil {
		return nil, err
	}

	if utils.StrSliceContains(i.Plists.BuildManifest.SupportedProductTypes, "mac") { // Is macOS IPSW
		if len(arches) == 0 {
			selMatches := []string{}
			prompt := &survey.MultiSelect{
				Message:  "Which files would you like to extract:",
				Options:  matches,
				PageSize: 15,
			}
			if err := survey.AskOne(prompt, &selMatches); err != nil {
				if err == terminal.InterruptErr {
					log.Warn("Exiting...")
					return nil, nil
				}
				return nil, err
			}
			matches = selMatches
		} else {
			var filtered []string
			r := regexp.MustCompile(fmt.Sprintf("%s(%s)%s", CacheRegex, strings.Join(arches, "|"), CacheRegexEnding))
			for _, match := range matches {
				if r.MatchString(match) {
					filtered = append(filtered, match)
				}
			}

			if len(filtered) == 0 {
				return nil, fmt.Errorf("no dyld_shared_cache files found matching the specified archs: %v", arches)
			}
			matches = filtered
		}
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("failed to find dyld_shared_cache(s) in DMG: %s", dmgPath)
	}

	var artifacts []string
	for _, match := range matches {
		dyldDest := filepath.Join(destPath, filepath.Base(match))
		utils.Indent(log.Info, 3)(fmt.Sprintf("Extracting %s to %s", filepath.Base(match), dyldDest))
		if err := utils.Copy(match, dyldDest); err != nil {
			return nil, fmt.Errorf("failed to copy %s to %s: %v", match, dyldDest, err)
		}
		artifacts = append(artifacts, dyldDest)
	}

	return artifacts, nil
}

// Extract extracts dyld_shared_cache from IPSW
func Extract(ipsw, destPath string, arches []string) ([]string, error) {

	if runtime.GOOS == "windows" {
		return nil, fmt.Errorf("dyld extraction is not supported on Windows (see github.com/blacktop/go-apfs)")
	}

	i, err := info.Parse(ipsw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse IPSW: %v", err)
	}

	dmgPath, err := i.GetSystemOsDmg()
	if err != nil {
		dmgPath, err = i.GetFileSystemOsDmg()
		if err != nil {
			return nil, fmt.Errorf("failed to get File System DMG: %v", err)
		}
	}

	// check if filesystem DMG already exists (due to previous mount command)
	if _, err := os.Stat(dmgPath); os.IsNotExist(err) {
		dmgs, err := utils.Unzip(ipsw, "", func(f *zip.File) bool {
			return strings.EqualFold(filepath.Base(f.Name), dmgPath)
		})
		if err != nil {
			return nil, fmt.Errorf("failed to extract %s from IPSW: %v", dmgPath, err)
		}
		if len(dmgs) == 0 {
			return nil, fmt.Errorf("File System %s NOT found in IPSW", dmgPath)
		}
		defer os.Remove(dmgs[0])
	}

	return ExtractFromDMG(i, dmgPath, destPath, arches)
}

func ExtractFromRemoteCryptex(zr *zip.Reader, destPath string, arches []string) error {
	found := false

	for _, zf := range zr.File {
		if regexp.MustCompile(`cryptex-system-arm64?e$`).MatchString(zf.Name) {
			found = true
			rc, err := zf.Open()
			if err != nil {
				return fmt.Errorf("failed to open cryptex-system-arm64e: %v", err)
			}
			defer rc.Close()
			// setup progress bar
			var total int64 = int64(zf.UncompressedSize64)
			p := mpb.New(
				mpb.WithWidth(60),
				mpb.WithRefreshRate(180*time.Millisecond),
			)
			bar := p.New(total,
				mpb.BarStyle().Lbound("[").Filler("=").Tip(">").Padding("-").Rbound("|"),
				mpb.PrependDecorators(
					decor.CountersKibiByte("\t% .2f / % .2f"),
				),
				mpb.AppendDecorators(
					decor.OnComplete(decor.AverageETA(decor.ET_STYLE_GO), "✅ "),
					decor.Name(" ] "),
					decor.AverageSpeed(decor.UnitKiB, "% .2f"),
				),
			)
			// create proxy reader
			proxyReader := bar.ProxyReader(io.LimitReader(rc, total))
			defer proxyReader.Close()

			in, err := os.CreateTemp("", "cryptex-system-arm64e")
			if err != nil {
				return fmt.Errorf("failed to create temp file for cryptex-system-arm64e: %v", err)
			}
			defer os.Remove(in.Name())

			log.Info("Extracting cryptex-system-arm64e from remote OTA")
			io.Copy(in, proxyReader)
			// wait for our bar to complete and flush and close remote zip and temp file
			p.Wait()
			in.Close()

			out, err := os.CreateTemp("", "cryptex-system-arm64e.decrypted.*.dmg")
			if err != nil {
				return fmt.Errorf("failed to create temp file for cryptex-system-arm64e.decrypted: %v", err)
			}
			defer os.Remove(out.Name())
			out.Close()

			log.Info("Patching cryptex-system-arm64e")
			if err := ridiff.RawImagePatch("", in.Name(), out.Name(), 0); err != nil {
				return fmt.Errorf("failed to patch cryptex-system-arm64e: %v", err)

			}

			i, err := info.ParseZipFiles(zr.File)
			if err != nil {
				return fmt.Errorf("failed to parse info from cryptex-system-arm64e: %v", err)
			}

			if _, err := ExtractFromDMG(i, out.Name(), destPath, arches); err != nil {
				return fmt.Errorf("failed to extract dyld_shared_cache from cryptex-system-arm64e: %v", err)
			}

			return nil
		}
	}

	if !found {
		return fmt.Errorf("cryptex-system-arm64e NOT found in remote zip")
	}

	return nil
}

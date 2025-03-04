/*
Copyright © 2018-2023 blacktop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package dyld

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/alecthomas/chroma/quick"
	"github.com/apex/log"
	"github.com/blacktop/go-macho"
	"github.com/blacktop/go-macho/pkg/fixupchains"
	swift "github.com/blacktop/ipsw/internal/swift"
	"github.com/blacktop/ipsw/internal/utils"
	"github.com/blacktop/ipsw/pkg/dyld"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var symAddrColor = color.New(color.Faint).SprintfFunc()
var symImageColor = color.New(color.Faint, color.FgBlue).SprintfFunc()
var symTypeColor = color.New(color.Faint, color.FgCyan).SprintfFunc()
var symLibColor = color.New(color.Faint, color.FgMagenta).SprintfFunc()
var symNameColor = color.New(color.Bold).SprintFunc()

func init() {
	DyldCmd.AddCommand(MachoCmd)
	MachoCmd.Flags().BoolP("all", "a", false, "Parse ALL dylibs")
	MachoCmd.Flags().BoolP("loads", "l", false, "Print the load commands")
	MachoCmd.Flags().BoolP("json", "j", false, "Print the TOC as JSON")
	MachoCmd.Flags().BoolP("objc", "o", false, "Print ObjC info")
	MachoCmd.Flags().BoolP("objc-refs", "r", false, "Print ObjC references")
	// MachoCmd.Flags().BoolP("swift", "w", false, "🚧 Print Swift info")
	MachoCmd.Flags().BoolP("symbols", "n", false, "Print symbols")
	MachoCmd.Flags().BoolP("starts", "f", false, "Print function starts")
	MachoCmd.Flags().BoolP("strings", "s", false, "Print cstrings")
	MachoCmd.Flags().BoolP("stubs", "b", false, "Print stubs")
	MachoCmd.Flags().String("search", "", "Search for byte pattern")

	MachoCmd.Flags().BoolP("extract", "x", false, "🚧 Extract the dylib")
	MachoCmd.Flags().String("output", "", "Directory to extract the dylib(s)")
	MachoCmd.Flags().Bool("force", false, "Overwrite existing extracted dylib(s)")

	MachoCmd.MarkZshCompPositionalArgumentFile(1)
}

// MachoCmd represents the macho command
var MachoCmd = &cobra.Command{
	Use:           "macho <dyld_shared_cache> <dylib>",
	Aliases:       []string{"m"},
	Short:         "Parse a dylib file",
	Args:          cobra.MinimumNArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		if viper.GetBool("verbose") {
			log.SetLevel(log.DebugLevel)
		}

		showLoadCommands, _ := cmd.Flags().GetBool("loads")
		showLoadCommandsAsJSON, _ := cmd.Flags().GetBool("json")
		showObjC, _ := cmd.Flags().GetBool("objc")
		showObjcRefs, _ := cmd.Flags().GetBool("objc-refs")
		showSwift, _ := cmd.Flags().GetBool("swift")
		dumpSymbols, _ := cmd.Flags().GetBool("symbols")
		showFuncStarts, _ := cmd.Flags().GetBool("starts")
		dumpStrings, _ := cmd.Flags().GetBool("strings")
		dumpStubs, _ := cmd.Flags().GetBool("stubs")
		searchPattern, _ := cmd.Flags().GetString("search")
		dumpALL, _ := cmd.Flags().GetBool("all")
		extractDylib, _ := cmd.Flags().GetBool("extract")
		extractPath, _ := cmd.Flags().GetString("output")
		forceExtract, _ := cmd.Flags().GetBool("force")

		onlyFuncStarts := !showLoadCommands && !showObjC && !showSwift && !dumpStubs && showFuncStarts
		onlyStubs := !showLoadCommands && !showObjC && !showSwift && !showFuncStarts && dumpStubs
		onlySearch := !showLoadCommands && !showObjC && !showSwift && !showFuncStarts && !dumpStubs && searchPattern != ""

		dscPath := filepath.Clean(args[0])

		fileInfo, err := os.Lstat(dscPath)
		if err != nil {
			return fmt.Errorf("file %s does not exist", dscPath)
		}

		// Check if file is a symlink
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			symlinkPath, err := os.Readlink(dscPath)
			if err != nil {
				return errors.Wrapf(err, "failed to read symlink %s", dscPath)
			}
			// TODO: this seems like it would break
			linkParent := filepath.Dir(dscPath)
			linkRoot := filepath.Dir(linkParent)

			dscPath = filepath.Join(linkRoot, symlinkPath)
		}

		f, err := dyld.Open(dscPath)
		if err != nil {
			return err
		}
		defer f.Close()

		if len(args) > 1 || dumpALL {

			var images []*dyld.CacheImage

			foundPattern := false

			if dumpALL {
				images = f.Images
			} else {
				image, err := f.Image(args[1])
				if err != nil {
					return fmt.Errorf("image not in %s: %v", dscPath, err)
				}
				images = append(images, image)
			}

			for _, image := range images {

				if dumpALL {
					log.WithField("name", image.Name).Debug("Image")
				}

				m, err := image.GetMacho()
				if err != nil {
					log.Warnf("failed to parse full MachO for %s: %v", image.Name, err)
					if extractDylib {
						continue
					}
					m, err = image.GetPartialMacho()
					if err != nil {
						return err
					}
				}
				defer m.Close()

				if extractDylib {

					folder := filepath.Dir(dscPath) // default to folder of shared cache
					if len(extractPath) > 0 {
						folder = extractPath
					}
					fname := filepath.Join(folder, filepath.Base(image.Name)) // default to NOT full dylib path
					if dumpALL {
						fname = filepath.Join(folder, image.Name)
					}

					if _, err := os.Stat(fname); os.IsNotExist(err) || forceExtract {
						var dcf *fixupchains.DyldChainedFixups
						if m.HasFixups() {
							dcf, err = m.DyldChainedFixups()
							if err != nil {
								return fmt.Errorf("failed to parse fixups from in memory MachO: %v", err)
							}
						}

						image.ParseLocalSymbols(false)

						// cc, err := m.GetObjCClasses()
						// if err != nil {
						// 	return err
						// }
						// for _, c := range cc {
						// 	fmt.Println(c)
						// }

						err = m.Export(fname, dcf, m.GetBaseAddress(), image.GetLocalSymbolsAsMachoSymbols())
						if err != nil {
							return fmt.Errorf("failed to export entry MachO %s; %v", image.Name, err)
						}

						if err := rebaseMachO(f, fname); err != nil {
							return fmt.Errorf("failed to rebase macho via cache slide info: %v", err)
						}
						if !dumpALL {
							log.Infof("Created %s", fname)
						}
					} else {
						if !dumpALL {
							log.Warnf("dylib already exists: %s", fname)
						}
					}
					continue
				}

				if showLoadCommands || !showObjC && !dumpSymbols && !dumpStrings && !showFuncStarts && !dumpStubs && searchPattern == "" {
					if showLoadCommandsAsJSON {
						dat, err := m.FileTOC.MarshalJSON()
						if err != nil {
							return fmt.Errorf("failed to marshal MachO table of contents as JSON: %v", err)
						}
						fmt.Println(string(dat))
					} else {
						fmt.Println(m.FileTOC.String())
					}
				}

				if showObjC {
					fmt.Println("Objective-C")
					fmt.Println("===========")
					if m.HasObjC() {
						if info, err := m.GetObjCImageInfo(); err == nil {
							fmt.Println(info.Flags)
						} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
							log.Error(err.Error())
						}
						if viper.GetBool("verbose") {
							fmt.Println(m.GetObjCToc())
						}
						if protos, err := m.GetObjCProtocols(); err == nil {
							for _, proto := range protos {
								if viper.GetBool("verbose") {
									if viper.GetBool("color") {
										quick.Highlight(os.Stdout, swift.DemangleBlob(proto.Verbose()), "objc", "terminal256", "nord")
										quick.Highlight(os.Stdout, "\n/****************************************/\n\n", "objc", "terminal256", "nord")
									} else {
										fmt.Println(swift.DemangleBlob(proto.Verbose()))
									}
								} else {
									if viper.GetBool("color") {
										quick.Highlight(os.Stdout, proto.String()+"\n", "objc", "terminal256", "nord")
									} else {
										fmt.Println(proto.String())
									}
								}
							}
						} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
							log.Error(err.Error())
						}
						if classes, err := m.GetObjCClasses(); err == nil {
							for _, class := range classes {
								if viper.GetBool("verbose") {
									if viper.GetBool("color") {
										quick.Highlight(os.Stdout, swift.DemangleBlob(class.Verbose()), "objc", "terminal256", "nord")
										quick.Highlight(os.Stdout, "\n/****************************************/\n\n", "objc", "terminal256", "nord")
									} else {
										fmt.Println(swift.DemangleBlob(class.Verbose()))
									}
								} else {
									if viper.GetBool("color") {
										quick.Highlight(os.Stdout, class.String()+"\n", "objc", "terminal256", "nord")
									} else {
										fmt.Println(class.String())
									}
								}
							}
						} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
							log.Error(err.Error())
						}
						if cats, err := m.GetObjCCategories(); err == nil {
							for _, cat := range cats {
								if viper.GetBool("verbose") {
									if viper.GetBool("color") {
										quick.Highlight(os.Stdout, swift.DemangleBlob(cat.Verbose()), "objc", "terminal256", "nord")
										quick.Highlight(os.Stdout, "\n/****************************************/\n\n", "objc", "terminal256", "nord")
									} else {
										fmt.Println(swift.DemangleBlob(cat.Verbose()))
									}
								} else {
									if viper.GetBool("color") {
										quick.Highlight(os.Stdout, cat.String()+"\n", "objc", "terminal256", "nord")
									} else {
										fmt.Println(cat.String())
									}
								}
							}
						} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
							log.Error(err.Error())
						}
						if showObjcRefs {
							if protRefs, err := m.GetObjCProtoReferences(); err == nil {
								fmt.Printf("\n@protocol refs\n")
								for off, prot := range protRefs {
									fmt.Printf("0x%011x => 0x%011x: %s\n", off, prot.Ptr, prot.Name)
								}
							} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
								log.Error(err.Error())
							}
							if clsRefs, err := m.GetObjCClassReferences(); err == nil {
								fmt.Printf("\n@class refs\n")
								for off, cls := range clsRefs {
									fmt.Printf("0x%011x => 0x%011x: %s\n", off, cls.ClassPtr, cls.Name)
									// if viper.GetBool("verbose") {
									// 	fmt.Println(cls.Verbose())
									// } else {
									// 	fmt.Println(cls.String())
									// }
								}
							} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
								log.Error(err.Error())
							}
							if supRefs, err := m.GetObjCSuperReferences(); err == nil {
								fmt.Printf("\n@super refs\n")
								for off, sup := range supRefs {
									fmt.Printf("0x%011x => 0x%011x: %s\n", off, sup.ClassPtr, sup.Name)
								}
							} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
								log.Error(err.Error())
							}
							if selRefs, err := m.GetObjCSelectorReferences(); err == nil {
								fmt.Printf("\n@selectors refs\n")
								for off, sel := range selRefs {
									fmt.Printf("0x%011x => 0x%011x: %s\n", off, sel.VMAddr, sel.Name)
								}
							} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
								log.Error(err.Error())
							}
							if viper.GetBool("verbose") {
								if classes, err := m.GetObjCClassNames(); err == nil {
									fmt.Printf("\n@objc_classname\n")
									for vmaddr, className := range classes {
										fmt.Printf("0x%011x: %s\n", vmaddr, className)
									}
								} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
									log.Error(err.Error())
								}
								if methods, err := m.GetObjCMethodNames(); err == nil {
									fmt.Printf("\n@objc_methname\n")
									for vmaddr, method := range methods {
										fmt.Printf("0x%011x: %s\n", vmaddr, method)
									}
								} else if !errors.Is(err, macho.ErrObjcSectionNotFound) {
									log.Error(err.Error())
								}
							}
						}

					} else {
						fmt.Println("  - no objc")
					}
					fmt.Println()
				}

				if showSwift {
					fmt.Println("Swift")
					fmt.Println("=====")
					info, err := m.GetObjCImageInfo()
					if err != nil {
						if !errors.Is(err, macho.ErrObjcSectionNotFound) {
							return err
						}
					}
					if info != nil && info.HasSwift() {
						if fields, err := m.GetSwiftFields(); err == nil {
							for _, field := range fields {
								if viper.GetBool("verbose") {
									if viper.GetBool("color") {
										quick.Highlight(os.Stdout, swift.DemangleBlob(field.String()), "swift", "terminal256", "nord")
										quick.Highlight(os.Stdout, "\n/****************************************/\n\n", "swift", "terminal256", "nord")
									} else {
										fmt.Println(swift.DemangleBlob(field.String()))
									}
								} else {
									if viper.GetBool("color") {
										quick.Highlight(os.Stdout, field.String(), "swift", "terminal256", "nord")
										quick.Highlight(os.Stdout, "\n/****************************************/\n\n", "swift", "terminal256", "nord")
									} else {
										fmt.Println(field.String())
									}
								}
							}
						} else if !errors.Is(err, macho.ErrSwiftSectionError) {
							log.Error(err.Error())
						}
					} else {
						fmt.Println("  - no swift")
					}
					fmt.Println()
				}

				if showFuncStarts {
					if !onlyFuncStarts {
						fmt.Println("FUNCTION STARTS")
						fmt.Println("===============")
					}
					if m.FunctionStarts() != nil {
						for _, fn := range m.GetFunctions() {
							if viper.GetBool("verbose") {
								fmt.Printf("%#016x-%#016x\n", fn.StartAddr, fn.EndAddr)
							} else {
								fmt.Printf("0x%016X\n", fn.StartAddr)
							}
						}
					}
				}

				if dumpSymbols {
					image.ParseLocalSymbols(false)
					w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
					if m.Symtab != nil {
						fmt.Println("SYMBOLS")
						fmt.Println("=======")
						undeflush := false
						for _, sym := range m.Symtab.Syms {
							if sym.Type.IsUndefinedSym() && !undeflush {
								w.Flush()
								undeflush = true
							}
							if sym.Value == 0 {
								fmt.Fprintf(w, "              %s\n", strings.Join([]string{symTypeColor(sym.GetType(m)), symNameColor(sym.Name), symLibColor(sym.GetLib(m))}, "\t"))
							} else {
								if sym.Name == "<redacted>" {
									if name, ok := f.AddressToSymbol[sym.Value]; ok {
										sym.Name = name
									}
								}
								fmt.Fprintf(w, "%s:  %s\n", symAddrColor("%#09x", sym.Value), strings.Join([]string{symTypeColor(sym.GetType(m)), symNameColor(sym.Name), symLibColor(sym.GetLib(m))}, "\t"))
							}
						}
						w.Flush()
					}
					if binds, err := m.GetBindInfo(); err == nil {
						fmt.Printf("\nDyld Binds\n")
						fmt.Println("----------")
						for _, bind := range binds {
							fmt.Fprintf(w, "%#09x:\t(%s.%s|from %s)\t%s\n", bind.Start+bind.Offset, bind.Segment, bind.Section, bind.Dylib, bind.Name)
						}
						w.Flush()
					}
					// Dedup these symbols (has repeats but also additional symbols??)
					if m.DyldExportsTrie() != nil && m.DyldExportsTrie().Size > 0 && viper.GetBool("verbose") {
						fmt.Printf("\nDyld Exports\n")
						fmt.Println("------------")
						exports, err := m.DyldExports()
						if err != nil {
							return err
						}
						for _, export := range exports {
							if export.Flags.ReExport() {
								export.FoundInDylib = m.ImportedLibraries()[export.Other-1]
								reimg, err := f.Image(export.FoundInDylib)
								if err != nil {
									return err
								}
								if rexpSym, err := reimg.GetExport(export.ReExport); err == nil {
									export.Address = rexpSym.Address
								}
							}
							fmt.Println(export)
						}
					}
				}

				if dumpStrings {
					fmt.Printf("\nCStrings\n")
					fmt.Println("--------")
					for _, sec := range m.Sections {
						if sec.Flags.IsCstringLiterals() || sec.Seg == "__TEXT" && sec.Name == "__const" {
							uuid, off, err := f.GetOffset(sec.Addr)
							if err != nil {
								return fmt.Errorf("failed to get offset for %s.%s: %v", sec.Seg, sec.Name, err)
							}
							dat, err := f.ReadBytesForUUID(uuid, int64(off), sec.Size)
							if err != nil {
								return fmt.Errorf("failed to read cstrings in %s.%s: %v", sec.Seg, sec.Name, err)
							}

							csr := bytes.NewBuffer(dat)

							for {
								pos := sec.Addr + uint64(csr.Cap()-csr.Len())

								s, err := csr.ReadString('\x00')

								if err == io.EOF {
									break
								}

								if err != nil {
									return fmt.Errorf("failed to read string: %v", err)
								}

								s = strings.Trim(s, "\x00")

								if len(s) > 0 {
									if (sec.Seg == "__TEXT" && sec.Name == "__const") && !utils.IsASCII(s) {
										continue // skip non-ascii strings when dumping __TEXT.__const
									}
									fmt.Printf("%#x: %#v\n", pos, s)
								}
							}
						}
					}

					if cfstrs, err := m.GetCFStrings(); err == nil {
						if len(cfstrs) > 0 {
							fmt.Printf("\nCFStrings\n")
							fmt.Println("---------")
							for _, cfstr := range cfstrs {
								fmt.Printf("%#09x:  %#v\n", cfstr.Address, cfstr.Name)
							}
						}
					}
				}

				if dumpStubs {
					if !onlyStubs {
						fmt.Printf("\nStubs\n")
						fmt.Println("=====")
					}
					if err := image.Analyze(); err != nil {
						return err
					}
					for stubAddr, addr := range image.Analysis.SymbolStubs {
						if symName, ok := f.AddressToSymbol[addr]; ok {
							fmt.Printf("%#x => %#x: %s\n", stubAddr, addr, symName)
						}
					}
				}

				if len(searchPattern) > 0 {
					var run string
					var gadget [][]byte

					patternBytes := strings.Fields(searchPattern) // split on whitespace

					for idx, abyte := range patternBytes {
						if abyte == "*" {
							if len(run) > 0 { // got a wildcard, but you were building a run
								pattern, err := hex.DecodeString(run)
								if err != nil {
									return fmt.Errorf("failed to decode pattern '%s': %v", run, err)
								}
								// add the run to the gadget list
								gadget = append(gadget, pattern)
								// zero out the run
								run = ""
							}
							gadget = append(gadget, []byte{}) // add a wildcard to the gadget as empty array
						} else {
							run += abyte
							if idx == len(patternBytes)-1 { // last byte
								pattern, err := hex.DecodeString(run)
								if err != nil {
									return fmt.Errorf("failed to decode pattern '%s': %v", run, err)
								}
								// add the run to the gadget list
								gadget = append(gadget, pattern)
							}
						}
					}

					if !dumpALL || !onlySearch {
						fmt.Printf("\nSearch Results\n")
						fmt.Println("--------------")
					}

					if textSeg := m.Segment("__TEXT"); textSeg != nil {
						uuid, off, err := f.GetOffset(textSeg.Addr)
						if err != nil {
							return fmt.Errorf("failed to get offset for %s: %v", textSeg.Name, err)
						}
						data, err := f.ReadBytesForUUID(uuid, int64(off), textSeg.Filesz)
						if err != nil {
							return fmt.Errorf("failed to read cstrings in %s: %v", textSeg.Name, err)
						}

						i := 0
						found := 0
						foundOffset := uint64(0)

						done := func() {
							foundPattern = true
							if dumpALL {
								fmt.Printf("%#x\t%s\n", textSeg.Addr+foundOffset, image.Name)
							} else {
								fmt.Printf("%#x\n", textSeg.Addr+foundOffset)
							}
						}

						for found >= 0 && i < len(data) { // stop if not found
							foundFirstPart := false
							for idx, gad := range gadget {
								if len(gad) == 0 { //  wildcards
									if foundFirstPart && idx == len(gadget)-1 { // last wildcard after found; DONE
										done()
									}
									i += 1
								} else if found = bytes.Index(data[i:], gad); foundFirstPart && found == 0 { // found next part of pattern
									if idx == len(gadget)-1 { // last part of pattern; DONE
										done()
									}
									i += len(gad)
								} else if !foundFirstPart && found >= 0 {
									if idx == 0 { // found first part of pattern
										foundFirstPart = true
										foundOffset = uint64(i + found)
									}
									if len(gadget) == 1 { // only one part of pattern; DONE
										done()
									}
									i += found + len(gad)
								} else if found < 0 { // pattern broken or not found
									i += len(gad)
									break
								}
							}
						}
					}
				}
			}

			if len(searchPattern) > 0 && !foundPattern {
				return fmt.Errorf("pattern '%s' not found", searchPattern)
			}

		} else {
			log.Error("you must supply a dylib MachO to parse")
		}

		return nil
	},
}

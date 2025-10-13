/*
Copyright © 2024 Ulises Ruz Puga <ulises.ruz@gmail.com>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

func copyFile(src, dst string, overwrite bool) error {
	fmt.Println(src, "->", dst)
	sourceFile, err := os.Open(src)
	var bytescopied int64
	bytescopied = 0
	if err != nil {
		return err
	}
	if !overwrite {
		err2 := CopyDataFile(sourceFile, dst, bytescopied)
		if err2 != nil {
			return err2
		}
	} else {
		f1, err := os.Stat(src)
		if err != nil {
			log.Println(err)
		}
		f2, err := os.Stat(dst)
		if err != nil {
			return err
		}
		if !os.SameFile(f1, f2) {
			err2 := CopyDataFile(sourceFile, dst, bytescopied)
			if err2 != nil {
				return err2
			}
		} else {
			fmt.Println("Files are the same, skipping")
		}
	}
	return nil
}

func CopyDataFile(sourceFile *os.File, dst string, bytescopied int64) error {
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer destFile.Close()

	bytescopied, err = io.Copy(destFile, sourceFile)

	if err = destFile.Sync(); err != nil {
		return err
	}
	if err != nil {
		return err
	}

	fmt.Println(dst, "Copied - ", bytescopied, "bytes")
	return nil
}

func copyDir(src, dst string, overwrite bool, wg *sync.WaitGroup, sem chan struct{}) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	os.MkdirAll(dst, os.ModePerm)

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			wg.Add(1)
			go func(src, dst string) {
				defer wg.Done()
				sem <- struct{}{} // acquire a token
				copyDir(src, dst, overwrite, wg, sem)
				<-sem // release the token
			}(srcPath, dstPath)
		} else {
			wg.Add(1)
			go func(src, dst string) {
				defer wg.Done()
				sem <- struct{}{} // acquire a token
				copyFile(src, dst, overwrite)
				<-sem // release the token
			}(srcPath, dstPath)
		}
	}

	return nil
}
func startCopy(srcDir, dstDir string, concurrent int, overwrite bool) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrent)

	wg.Add(1)
	go func() {
		defer wg.Done()
		sem <- struct{}{} // acquire a tokengit
		copyDir(srcDir, dstDir, overwrite, &wg, sem)
		<-sem
	}()

	wg.Wait()
}

// asyncCopyCmd represents the asyncCopy command
var asyncCopyCmd = &cobra.Command{
	Use:   "asyncCopy",
	Short: "Copy directories recursively and asynchronously from one location to another",
	Long:  `Copy directories recursively asynchronously from one location to another`,
	Run: func(cmd *cobra.Command, args []string) {
		srcDir, _ := cmd.Flags().GetString("srcDir")
		dstDir, _ := cmd.Flags().GetString("dstDir")
		concurrent, _ := cmd.Flags().GetInt("concurrency")
		overwrite, _ := cmd.Flags().GetBool("OverwriteIfDifferent")
		if srcDir == "" || dstDir == "" {
			cmd.Help()
			os.Exit(1)
		}
		startCopy(srcDir, dstDir, concurrent, overwrite)

	},
}

func init() {
	asyncCopyCmd.Flags().StringP("srcDir", "s", "", "Source directory")
	asyncCopyCmd.Flags().StringP("dstDir", "d", "", "Destination directory")
	asyncCopyCmd.Flags().IntP("concurrency", "c", 10, "Number of concurrent files to copy")
	asyncCopyCmd.Flags().BoolP("OverwriteIfDifferent", "o", true, "Overwrite destination file if different")
	rootCmd.AddCommand(asyncCopyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// asyncCopyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// asyncCopyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

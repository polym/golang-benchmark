package readline

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadLineByBufioScanner(fd *os.File) {
	fd.Seek(0, 0)
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		scanner.Text()
		//fmt.Println(len(line))
	}
}

func ReadLineByBufioReadLine(fd *os.File) {
	fd.Seek(0, 0)
	reader := bufio.NewReader(fd)
	for {
		_, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("read %v\n", err)
			return
		}
		//	fmt.Println(len(b))
	}
}

func ReadLineByBufioReadSlice(fd *os.File) {
	fd.Seek(0, 0)
	reader := bufio.NewReader(fd)
	for {
		_, err := reader.ReadSlice('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("read %v\n", err)
			return
		}
		//	fmt.Println(len(b))
	}
}

func ReadLineByDIY(fd *os.File) {
	fd.Seek(0, 0)
	buf := make([]byte, 1024)
	line := make([]byte, 1024)
	rest := 0
	for {
		n, err := fd.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("read %v\n", err)
			return
		}
		startPos := 0
		for k := range buf[:n] {
			if buf[k] == '\n' {
				copy(line[rest:], buf[startPos:k])
				//rest + k - startPos
				//fmt.Println(lineSize)
				// do someting
				startPos, rest = k+1, 0
				continue
			}
		}
		if startPos != 0 {
			copy(line, buf[startPos:])
			rest = n - startPos
		}
	}
}

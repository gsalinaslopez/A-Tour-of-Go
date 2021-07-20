package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rot13R rot13Reader) Read(p []byte) (n int, err error) {
    b := make([]byte, 1)
    out_count := 0
    for {
        _, err := rot13R.r.Read(b)
        if err != io.EOF {
            if out_count < len(p) {
                p[out_count] = b[0]
                if (b[0] >= 65 && b[0] <= 122) {
                    if (b[0] + 13 > 122) {
                        p[out_count] = 96 + (13 - (122 - b[0]))
                    } else {
                        p[out_count] += 13
                    }
                }
                out_count += 1
            } else {
                return out_count, nil
            }
        } else {
            return out_count, io.EOF
        }
    }
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

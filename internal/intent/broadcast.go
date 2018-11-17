package intent

import (
	"bytes"
	"context"
	"os/exec"
	"strconv"
	"strings"
)

type Broadcast struct {
	User         int
	Component    string
	Data         string
	Action       string
	extraString  map[string]string
	extraBool    map[string]bool
	extraInt     map[string]int
	extraLongs   map[string][]int32
	extraStrings map[string][]string
}

func (c *Broadcast) AddBool(key string, val bool) {
	if c.extraBool == nil {
		c.extraBool = make(map[string]bool)
	}
	c.extraBool[key] = val
}

func (c *Broadcast) AddString(key string, val string) {
	if c.extraString == nil {
		c.extraString = make(map[string]string)
	}
	c.extraString[key] = val
}

func (c *Broadcast) AddInt(key string, val int) {
	if c.extraInt == nil {
		c.extraInt = make(map[string]int)
	}
	c.extraInt[key] = val
}

func (c *Broadcast) AddLongs(key string, val []int32) {
	if c.extraLongs == nil {
		c.extraLongs = make(map[string][]int32)
	}
	c.extraLongs[key] = val
}

func (c *Broadcast) AddStrings(key string, val []string) {
	if c.extraStrings == nil {
		c.extraStrings = make(map[string][]string)
	}
	c.extraStrings[key] = val
}

func (c *Broadcast) Send(ctx context.Context) {
	args := []string{"broadcast", "--user", strconv.Itoa(c.User), "-n", c.Component}
	if c.Data != "" {
		args = append(args, "-d", c.Data)
	}
	if c.Action != "" {
		args = append(args, "-a", c.Action)
	}
	if c.extraString != nil {
		for key, val := range c.extraString {
			args = append(args, "--es", key, val)
		}
	}
	if c.extraBool != nil {
		for key, val := range c.extraBool {
			args = append(args, "--ez", key, strconv.FormatBool(val))
		}
	}
	if c.extraInt != nil {
		for key, val := range c.extraInt {
			args = append(args, "--ei", key, strconv.Itoa(val))
		}
	}
	if c.extraLongs != nil {
		for key, val := range c.extraLongs {
			t := bytes.NewBuffer([]byte{})
			for i, v := range val {
				t.WriteString(string(v))
				if i+1 != len(val) {
					t.WriteRune(',')
				}
			}
			args = append(args, "--ela", key, t.String())
		}
	}
	if c.extraStrings != nil {
		for key, val := range c.extraStrings {
			t := bytes.NewBuffer([]byte{})
			for i, v := range val {
				t.WriteString(strings.Replace(v, ",", "\\,", -1))
				if i+1 != len(val) {
					t.WriteRune(',')
				}
			}
			args = append(args, "--esa", key, t.String())
		}
	}
	cmd := exec.CommandContext(ctx, "/data/data/com.termux/files/usr/bin/am", args...)
	go cmd.Run()
}

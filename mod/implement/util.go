package implement

import (
	"fmt"
	"github.com/sunist-c/genius-invokation-simulator-backend/util"
	"gopkg.in/yaml.v2"
	"os"
	"os/user"
	"path"
	"time"
)

const (
	maxID uint16 = 1<<16 - 1
)

var (
	debugFlag = false
	metadata  = ModInfo{
		ModID:     0,
		LastID:    0,
		UsedID:    map[uint16]bool{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Author:    "",
	}
	metadataPath = ""
)

type ModInfo struct {
	ModID     uint64          `yaml:"mod_id"`
	LastID    uint16          `yaml:"last_id"`
	UsedID    map[uint16]bool `yaml:"id_pool"`
	CreatedAt time.Time       `yaml:"created_at"`
	UpdatedAt time.Time       `yaml:"updated_at"`
	Author    string          `yaml:"author"`
}

func InitMetaData() {
	// 非debug模式，获取真实ModID
	if !debugFlag {
		// 尝试读取配置
		if pwd, err := os.Getwd(); err != nil {
			// 无法获取pwd，意味着无法写入也无法读取，报panic
			panic(fmt.Sprintf("cannot get current working directory: %v", err))
		} else if bytes, err := os.ReadFile(path.Join(pwd, "metadata.yml")); err == nil {
			metadataPath = path.Join(pwd, "metadata.yml")

			// 存在metadata.yml，尝试加载
			if yaml.Unmarshal(bytes, &metadata) != nil {
				// 反序列化metadata.yml失败，报panic
				panic(fmt.Sprintf("cannot unmarshal metadata.yml: %v", err))
			} else {
				// 反序列化metadata.yml成功，加载
				return
			}
		} else {
			// 不存在metadata.yml，生成
			metadataPath = path.Join(pwd, "metadata.yml")

			// 生成ModID
			if macAddresses, err := util.GetMacAddresses(); err != nil {
				// 无法获取mac地址，实际环境不存在这种情况，报panic
				panic(fmt.Sprintf("couldn't get mac addresses: %v", err))
			} else if len(macAddresses) == 0 {
				// mac地址不存在，实际环境不存在这种情况，报panic
				panic(fmt.Sprintf("couldn't get mac addresses: %v", err))
			} else if uintMacAddress, err := util.GetUintMacAddress(macAddresses[0]); err != nil {
				// mac地址不合法，实际环境不存在这种情况，报panic
				panic(fmt.Sprintf("couldn't get mac address: %v", err))
			} else {
				// 成功获取mac地址，作为设备码使用自定义雪花算法生成UID
				metadata.ModID = util.GenerateUID(uintMacAddress, time.Now())
			}

			// 生成Author
			if currentUser, err := user.Current(); err != nil {
				// 获取CurrentUser失败，设置为Unknown
				metadata.Author = "Unknown"
			} else {
				// 成功获取CurrentUser
				metadata.Author = currentUser.Name
			}

			// 生成metadata上下文
			metadata.LastID = 0
			metadata.UsedID = map[uint16]bool{}

			// 生成TimeStamp
			metadata.CreatedAt = time.Now()

			// 将生成的metadata写入metadata.yml
			if err := flushMetadata(); err != nil {
				// 写入metadata.yml失败，报panic
				panic(err)
			}
		}
	}
}

func flushMetadata() error {
	if !debugFlag {
		metadata.UpdatedAt = time.Now()
		if oStream, err := yaml.Marshal(metadata); err != nil {
			// 序列化metadata.yml失败，实际环境不存在这种情况
			return fmt.Errorf("marshal metadata.yml failed: %v", err)
		} else if file, err := os.OpenFile(metadataPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755); err != nil {
			// 无法打开metadata.yml，实际环境不存在这种情况
			return fmt.Errorf("cannot open metadata.yml: %v", err)
		} else if _, err := file.Write(oStream); err != nil {
			// 无法写入metadata.yml，实际环境不存在这种情况
			return fmt.Errorf("cannot write metadata.yml: %v", err)
		} else {
			// 成功写入metadata.yml
			return nil
		}
	} else {
		return nil
	}
}

func SetDebugFlag(flag bool) {
	debugFlag = flag
}

func ModID() uint64 {
	if debugFlag {
		return 109951162778600
	} else if metadata.ModID != 0 {
		return metadata.ModID
	} else {
		InitMetaData()
		return metadata.ModID
	}
}

// NextID 使用内置的ID分配工具获取一个可用的不与其余被托管ID冲突的自增ID
func NextID() uint16 {
	defer func() {
		err := flushMetadata()
		if err != nil {
			panic(err)
		}
	}()

	if !metadata.UsedID[metadata.LastID+1] {
		metadata.LastID += 1
		metadata.UsedID[metadata.LastID] = true
		return metadata.LastID
	} else {
		for i := uint32(1); i <= uint32(maxID); i++ {
			if !metadata.UsedID[uint16(i)] {
				metadata.LastID = uint16(i)
				metadata.UsedID[uint16(i)] = true
				return uint16(i)
			}
		}
	}

	panic("entity id overflow")
}

// UseID 使用内置的ID分配工具分配一个指定的ID，若不可用，则分配一个自增的ID
func UseID(want uint16) (success bool, result uint16) {
	defer func() {
		err := flushMetadata()
		if err != nil {
			panic(err)
		}
	}()

	if !metadata.UsedID[want] {
		// 成功分配want作为ID
		metadata.UsedID[want] = true
		return true, want
	} else {
		// 分配失败，从want开始查找下一个可用ID并分配
		for i := uint32(want); i <= uint32(maxID); i++ {
			if !metadata.UsedID[uint16(i)] {
				metadata.LastID = uint16(i)
				metadata.UsedID[uint16(i)] = true
				return false, uint16(i)
			}
		}

		// 分配失败，ID溢出，报panic
		panic("entity id overflow")
	}
}

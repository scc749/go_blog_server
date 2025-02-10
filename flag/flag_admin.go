package flag

import (
	"errors"
	"fmt"
	"os"
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/utils"
	"syscall"

	"github.com/gofrs/uuid"
	"golang.org/x/term"
)

// Admin 用于创建一个管理员用户
func Admin() error {
	var user database.User

	// 提示用户输入邮箱
	fmt.Print("Enter email: ")
	// 读取用户输入的邮箱
	var email string
	_, err := fmt.Scanln(&email)
	if err != nil {
		return fmt.Errorf("failed to read email: %w", err)
	}
	user.Email = email

	// 获取标准输入的文件描述符
	fd := int(syscall.Stdin)
	// 关闭回显，确保密码不会在终端显示
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer term.Restore(fd, oldState) // 恢复终端状态

	// 提示用户输入密码
	fmt.Print("Enter password: ")
	// 读取第一次输入的密码
	password, err := readPassword()
	fmt.Println()
	if err != nil {
		return err
	}

	// 提示用户再次输入密码进行确认
	fmt.Print("Confirm password: ")
	// 读取第二次输入的密码
	rePassword, err := readPassword()
	fmt.Println()
	if err != nil {
		return err
	}

	// 检查两次密码输入是否一致
	if password != rePassword {
		return errors.New("passwords do not match")
	}

	// 检查密码长度是否符合要求
	if len(password) < 8 || len(password) > 20 {
		return errors.New("password length should be between 8 and 20 characters")
	}

	// 填充用户数据
	user.UUID = uuid.Must(uuid.NewV4())
	user.Username = global.Config.Website.Name
	user.Password = utils.BcryptHash(password)
	user.RoleID = appTypes.Admin
	user.Avatar = "/image/avatar.jpg"
	user.Address = global.Config.Website.Address

	// 在数据库中创建管理员用户
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// readPassword 用于读取密码并且避免回显
func readPassword() (string, error) {
	var password string
	var buf [1]byte

	// 持续读取字符直到遇到换行符为止
	for {
		_, err := os.Stdin.Read(buf[:])
		if err != nil {
			return "", err
		}
		char := buf[0]

		// 检查是否为回车键，若是则终止输入
		if char == '\n' || char == '\r' {
			break
		}

		// 将输入的字符附加到密码字符串中
		password += string(char)
	}

	return password, nil
}

package main

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
	"log"
	"io"
)
func fileCreate(name string)  {   //这里为创建文件的方法
	gang := "\\"                  //因为是在windows上，所以我这里定义了反斜杠，最笨的方法
	_, err1 := os.Stat(name)      //判断文件是否存在，存在为真
	if err1 != nil {
		path, err3 := os.Getwd()   //获取当前目录路径，
		if err3 != nil {
			fmt.Println("error 3")
		}
		path2 := path + gang      //最笨的拼接方法
		path2 += name
		fmt.Println(path2)
		new2File, err2 := os.Create(path2)  //创建文件
		if err2 != nil {
			fmt.Println("error 2: %v", err2)
		}
		log.Println(new2File)
		new2File.Close()    //最后记得关闭此文件

	}else {
		fmt.Println("file is existd")
	}

}

func filedelete(name string)  {   //这里是文件的删除方法
	err := os.Remove(name)        //这里比较简单，直接调用了内置的os.Remove方法进行删除
	if err != nil {
		fmt.Println(err)
	}
	log.Println(name)
}

func filecopy(srourname, mubiaoname string)  {   //这里是文件的复制拷贝方法
	src, err1 := os.Open(srourname)             //打开此文件返回一个*file类型的结构体和报错，不存在则报错
	defer src.Close()                          //整个方法最后执行关闭
	if err1 != nil {
		fmt.Println(err1)
	}

	_, err6 := os.Stat(mubiaoname)          //判断文件是否存在
	if err6 != nil {
		fmt.Println(err6)
		mubiao2, err4 := os.Create(mubiaoname)   //创建文件
		if err4 != nil {
			fmt.Println(err4)
		}
		_, err2 := io.Copy(mubiao2, src)       //拷贝文件，这里的函数是将两个存在文件互相覆盖，所以上面先创建了文件
		if err2 != nil {
			fmt.Println(err2)
		}

	}else {
		fmt.Println("mubiao wenjian yijing cunzai")
	}

}

func main()  {
	app := cli.NewApp()          //这里利用了cli命令执行程序的工具
	app.Name = "jiange"
	app.Version = "1.1.9"
	app.Flags = []cli.Flag{            //实现一些帮助参数  标签
		cli.StringFlag{
			Name: "copy",
			Value: "english",
			Usage: "User copy file",
		},
		cli.StringFlag{
			Name: "move",
			Value: "english",
			Usage: "move file path",
		},
		cli.StringFlag{
			Name: "create",
			Value: "english",
			Usage: "User create files",
		},
		cli.StringFlag{
			Name: "delete",
			Value: "english",
			Usage: "User delete files",
		},


	}
	app.Commands = []cli.Command{     //这里定义了所有的命令，也就是功能
		{
			Name: "copy",
			Aliases: []string{"c"},
			Usage: "User copy file",
			Action: func(c *cli.Context) error{
				firsname := c.Args().Get(0)
				secondname := c.Args().Get(1)
				filecopy(firsname, secondname)
				fmt.Println("file copy")
				return nil
			},
		},
		{
			Name: "move",
			Aliases: []string{"m"},
			Usage: "move file path",
			Action: func(c *cli.Context) error {

				fmt.Println("file move")
				return nil
			},
		},
		{
			Name: "create",
			Aliases: []string{"create"},
			Usage: "create files",
			Action: func(c *cli.Context) error {
				name := c.Args().Get(0)
				fileCreate(name)
				fmt.Println("file create")
				return nil
			},
		},
		{
			Name: "delete",
			Aliases: []string{"d"},
			Usage: "delete files",
			Action: func(c *cli.Context) error {
				name := c.Args().Get(0)
				filedelete(name)
				fmt.Println("file delete")
				return nil
			},
		},
	}
	app.Run(os.Args)   //最后运行此程序
}
